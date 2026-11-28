package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"project3/auth"
)

func siap() (*Middleware, *auth.Penerbit) {
	p := auth.NewPenerbit("rahasia-test-1234567", time.Minute)
	return &Middleware{Tokens: p}, p
}

func ok(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }

func panggil(h http.HandlerFunc, header string) int {
	r := httptest.NewRequest("GET", "/x", nil)
	if header != "" {
		r.Header.Set("Authorization", header)
	}
	w := httptest.NewRecorder()
	h(w, r)
	return w.Code
}

func TestAuth(t *testing.T) {
	m, p := siap()
	valid, _ := p.Buat(1, "user")
	h := m.Auth(ok)

	kedaluwarsa := auth.NewPenerbit("rahasia-test-1234567", time.Minute)
	kedaluwarsa.Sekarang = func() time.Time { return time.Now().Add(-time.Hour) }
	lama, _ := kedaluwarsa.Buat(1, "user")

	tests := []struct {
		nama, header string
		want         int
	}{
		{"tanpa header", "", 401},
		{"bukan Bearer", "Basic abc", 401},
		{"Bearer kosong", "Bearer ", 401},
		{"token ngawur", "Bearer ngawur", 401},
		{"token kedaluwarsa", "Bearer " + lama, 401},
		{"token valid", "Bearer " + valid, 200},
	}
	for _, tt := range tests {
		t.Run(tt.nama, func(t *testing.T) {
			if got := panggil(h, tt.header); got != tt.want {
				t.Errorf("status %d, mau %d", got, tt.want)
			}
		})
	}
}

func TestRequireRole(t *testing.T) {
	m, p := siap()
	h := m.Auth(RequireRole("admin", ok))
	user, _ := p.Buat(1, "user")
	admin, _ := p.Buat(2, "admin")

	if got := panggil(h, ""); got != 401 {
		t.Errorf("tanpa token: %d, mau 401", got)
	}
	if got := panggil(h, "Bearer "+user); got != 403 {
		t.Errorf("role user: %d, mau 403", got)
	}
	if got := panggil(h, "Bearer "+admin); got != 200 {
		t.Errorf("role admin: %d, mau 200", got)
	}
}

func TestClaimsMasukContext(t *testing.T) {
	m, p := siap()
	token, _ := p.Buat(42, "admin")
	var dapat *auth.Claims
	h := m.Auth(func(w http.ResponseWriter, r *http.Request) { dapat = UserDariContext(r.Context()) })
	panggil(h, "Bearer "+token)
	if dapat == nil || dapat.Sub != "42" || dapat.Role != "admin" {
		t.Errorf("claims di context salah: %+v", dapat)
	}
}
