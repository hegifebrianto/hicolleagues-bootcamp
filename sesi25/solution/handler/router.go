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
	mux.HandleFunc("POST /api/register", a.Daftar) //untuk meregister user
	mux.HandleFunc("POST /api/login", a.Login)     //untuk login
	mux.HandleFunc("GET /api/buku", b.Semua)       //untuk mendapatkan semua buku
	mux.HandleFunc("GET /api/buku/{id}", b.Ambil)  //untuk mendapatkan buku {id}

	// Perlu login (token valid)
	mux.HandleFunc("GET /api/me", m.Auth(a.Saya))      //untuk mendapatkan data profile user
	mux.HandleFunc("POST /api/buku", m.Auth(b.Tambah)) //untuk menambahkan buku

	// Perlu login DAN role admin
	mux.HandleFunc("DELETE /api/buku/{id}", m.Auth(middleware.RequireRole(model.RoleAdmin, b.Hapus))) //untuk menghapus buku dari user

	return mux
}
