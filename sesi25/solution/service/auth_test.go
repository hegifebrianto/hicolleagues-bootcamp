package service

import (
	"errors"
	"sync"
	"testing"
	"time"

	"project3/auth"
	"project3/model"
	"project3/repository"
)

// mailerPalsu mencatat email yang "dikirim". Aman dipakai lintas goroutine.
type mailerPalsu struct {
	mu       sync.Mutex
	terkirim []string
	tahan    chan struct{} // bila tidak nil, KirimSelamatDatang menunggu sampai channel ditutup
	selesai  chan string
}

func (m *mailerPalsu) KirimSelamatDatang(email string) {
	if m.tahan != nil {
		<-m.tahan
	}
	m.mu.Lock()
	m.terkirim = append(m.terkirim, email)
	m.mu.Unlock()
	if m.selesai != nil {
		m.selesai <- email
	}
}

func buatService(m Mailer) *AuthService {
	return NewAuthService(repository.NewUserMemori(), auth.NewPenerbit("rahasia-test-1234567", time.Minute), m)
}

func TestDaftarValidasi(t *testing.T) {
	tests := []struct {
		nama, email, pass string
		field             string
	}{
		{"", "a@b.c", "123456", "nama"},
		{"Ani", "tanpa-at", "123456", "email"},
		{"Ani", "a@b.c", "123", "password"},
	}
	for _, tt := range tests {
		svc := buatService(&mailerPalsu{})
		_, err := svc.Daftar(tt.nama, tt.email, tt.pass)
		var ev model.ErrValidasi
		if !errors.As(err, &ev) || ev.Field != tt.field {
			t.Errorf("%+v: dapat %v, mau validasi field %s", tt, err, tt.field)
		}
	}
}

func TestDaftarDuplikatDanHash(t *testing.T) {
	svc := buatService(&mailerPalsu{})
	u, err := svc.Daftar("Ani", "Ani@Mail.com", "rahasia1")
	if err != nil {
		t.Fatal(err)
	}
	if u.Role != model.RoleUser {
		t.Errorf("user baru harus role user, dapat %q", u.Role)
	}
	if u.Email != "ani@mail.com" {
		t.Errorf("email harus huruf kecil, dapat %q", u.Email)
	}
	if u.PasswordHash == "" || u.PasswordHash == "rahasia1" {
		t.Error("password harus tersimpan sebagai hash")
	}
	if _, err := svc.Daftar("Ani2", "ani@mail.com", "rahasia1"); !errors.Is(err, model.ErrEmailTerpakai) {
		t.Errorf("dapat %v, mau ErrEmailTerpakai", err)
	}
}

// Membuktikan email dikirim ASYNC: Daftar harus selesai walau mailer masih "macet".
func TestDaftarEmailAsync(t *testing.T) {
	m := &mailerPalsu{tahan: make(chan struct{}), selesai: make(chan string, 1)}
	svc := buatService(m)

	selesaiDaftar := make(chan error, 1)
	go func() {
		_, err := svc.Daftar("Ani", "ani@mail.com", "rahasia1")
		selesaiDaftar <- err
	}()

	select {
	case err := <-selesaiDaftar:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Daftar menunggu email: kirim email harus di goroutine (go ...)")
	}

	close(m.tahan) // lepaskan mailer
	select {
	case email := <-m.selesai:
		if email != "ani@mail.com" {
			t.Errorf("email salah: %s", email)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("email tidak pernah terkirim")
	}
}

func TestLogin(t *testing.T) {
	svc := buatService(&mailerPalsu{})
	svc.Daftar("Ani", "ani@mail.com", "rahasia1")

	token, err := svc.Login("ANI@mail.com", "rahasia1")
	if err != nil {
		t.Fatalf("login benar gagal: %v", err)
	}
	if c, err := svc.Tokens.Cek(token); err != nil || c.Role != model.RoleUser {
		t.Errorf("token tidak benar: %+v %v", c, err)
	}

	for _, tt := range []struct{ email, pass string }{
		{"ani@mail.com", "salah"},
		{"tidakada@mail.com", "rahasia1"},
	} {
		if _, err := svc.Login(tt.email, tt.pass); !errors.Is(err, model.ErrKredensial) {
			t.Errorf("%+v: dapat %v, mau ErrKredensial (pesan harus sama agar email tidak bocor)", tt, err)
		}
	}
}
