package handler

import (
	"net/http"

	"project3/middleware"
	"project3/model"
	"project3/respon"
)

// NewRouter menyusun semua rute. Lihat bagaimana middleware "membungkus" handler.
func NewRouter(b *BukuHandler, a *AuthHandler, m *middleware.Middleware) http.Handler {
	mux := http.NewServeMux()

	// Publik
	mux.HandleFunc("GET /api/ping", func(w http.ResponseWriter, r *http.Request) {
		respon.Data(w, http.StatusOK, "pong")
	})
	mux.HandleFunc("POST /api/register", a.Daftar)
	mux.HandleFunc("POST /api/login", a.Login)
	mux.HandleFunc("GET /api/buku", b.Semua)
	mux.HandleFunc("GET /api/buku/{id}", b.Ambil)

	// Perlu login (token valid)
	mux.HandleFunc("GET /api/me", m.Auth(a.Saya))
	mux.HandleFunc("POST /api/buku", m.Auth(b.Tambah))

	// Perlu login DAN role admin
	mux.HandleFunc("DELETE /api/buku/{id}", m.Auth(middleware.RequireRole(model.RoleAdmin, b.Hapus)))

	return mux
}
