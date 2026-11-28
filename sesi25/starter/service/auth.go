package service

import (
	"errors"
	"strings"

	"project3/auth"
	"project3/model"
	"project3/repository"
)

type AuthService struct {
	Users  repository.UserRepo
	Tokens *auth.Penerbit
	Mailer Mailer
}

func NewAuthService(u repository.UserRepo, t *auth.Penerbit, m Mailer) *AuthService {
	return &AuthService{Users: u, Tokens: t, Mailer: m}
}

// Daftar membuat user baru (selalu role "user"; admin hanya lewat seed/DB).
func (s *AuthService) Daftar(nama, email, password string) (model.User, error) {
	nama = strings.TrimSpace(nama)
	email = strings.ToLower(strings.TrimSpace(email))
	switch {
	case nama == "":
		return model.User{}, model.ErrValidasi{Field: "nama", Pesan: "wajib diisi"}
	case !strings.Contains(email, "@"):
		return model.User{}, model.ErrValidasi{Field: "email", Pesan: "tidak valid"}
	case len(password) < 6:
		return model.User{}, model.ErrValidasi{Field: "password", Pesan: "minimal 6 karakter"}
	}
	if _, err := s.Users.CariEmail(email); err == nil {
		return model.User{}, model.ErrEmailTerpakai
	} else if !errors.Is(err, model.ErrTidakDitemukan) {
		return model.User{}, err
	}
	hash, err := auth.HashPassword(password)
	if err != nil {
		return model.User{}, err
	}
	u, err := s.Users.Simpan(model.User{Nama: nama, Email: email, PasswordHash: hash, Role: model.RoleUser})
	if err != nil {
		return model.User{}, err
	}
	// TODO 4: email lambat (2 detik) membuat Daftar ikut lambat dan test async gagal.
	// Jalankan di goroutine agar handler langsung membalas.
	s.Mailer.KirimSelamatDatang(u.Email)
	return u, nil
}

// TODO 3: Login.
//  1. email -> huruf kecil + TrimSpace
//  2. u, err := s.Users.CariEmail(email)
//     errors.Is(err, model.ErrTidakDitemukan) -> return "", model.ErrKredensial
//     err lain -> return "", err
//  3. auth.CekPassword(u.PasswordHash, password) false -> return "", model.ErrKredensial
//     (pesan error SAMA untuk email salah & password salah: jangan bocorkan email mana yang terdaftar)
//  4. return s.Tokens.Buat(u.ID, u.Role)
func (s *AuthService) Login(email, password string) (string, error) {
	return "", errors.New("TODO: belum diimplementasi")
}
