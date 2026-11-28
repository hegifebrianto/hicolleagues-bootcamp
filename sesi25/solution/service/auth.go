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
	// Email lambat dikirim di goroutine: handler langsung membalas, tidak menunggu.
	go s.Mailer.KirimSelamatDatang(u.Email)
	return u, nil
}

// Login memeriksa email+password dan mengembalikan token.
// Email tak dikenal dan password salah SENGAJA menghasilkan error yang sama.
func (s *AuthService) Login(email, password string) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	u, err := s.Users.CariEmail(email)
	if errors.Is(err, model.ErrTidakDitemukan) {
		return "", model.ErrKredensial
	}
	if err != nil {
		return "", err
	}
	if !auth.CekPassword(u.PasswordHash, password) {
		return "", model.ErrKredensial
	}
	return s.Tokens.Buat(u.ID, u.Role)
}
