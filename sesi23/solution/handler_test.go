package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// BONUS: menguji lapisan HTTP juga tanpa server & tanpa database (httptest + repoPalsu).
func rakitTes(repo BukuRepo) http.Handler {
	h := NewBukuHandler(NewBukuService(repo))
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/buku", h.Semua)
	mux.HandleFunc("GET /api/buku/{id}", h.Ambil)
	mux.HandleFunc("POST /api/buku", h.Tambah)
	mux.HandleFunc("DELETE /api/buku/{id}", h.Hapus)
	return mux
}

func TestHandler_StatusCode(t *testing.T) {
	kasus := []struct {
		nama, method, url, body string
		want                    int
	}{
		{"daftar", "GET", "/api/buku", "", 200},
		{"ambil ada", "GET", "/api/buku/1", "", 200},
		{"ambil tidak ada", "GET", "/api/buku/99", "", 404},
		{"id bukan angka", "GET", "/api/buku/abc", "", 400},
		{"tambah valid", "POST", "/api/buku", `{"judul":"Baru","penulis":"A","tahun":2020}`, 201},
		{"tambah judul kosong", "POST", "/api/buku", `{"judul":"","tahun":2020}`, 400},
		{"tambah body rusak", "POST", "/api/buku", `bukan json`, 400},
		{"hapus ada", "DELETE", "/api/buku/1", "", 204},
		{"hapus tidak ada", "DELETE", "/api/buku/99", "", 404},
	}
	for _, k := range kasus {
		t.Run(k.nama, func(t *testing.T) {
			app := rakitTes(&repoPalsu{data: []Buku{{ID: 1, Judul: "Ada", Tahun: 2000}}})
			rec := httptest.NewRecorder()
			app.ServeHTTP(rec, httptest.NewRequest(k.method, k.url, strings.NewReader(k.body)))
			if rec.Code != k.want {
				t.Errorf("status %d, ingin %d (body: %s)", rec.Code, k.want, rec.Body.String())
			}
		})
	}
}
