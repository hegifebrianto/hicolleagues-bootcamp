// Project 2 versi BERLAPIS: handler -> service -> repository -> database.
//
//	go run .                  (butuh MySQL; lihat README)
//	STORE=memory go run .     (tanpa database -- untuk demo/tes cepat)
package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// initApp merakit semua lapisan (bonus Latihan 3).
// Urutan: DB -> Repo -> Service -> Handler. Hanya di sinilah dependency "dikenalkan".
func initApp(cfg Config) (http.Handler, func()) {
	var repo BukuRepo
	tutup := func() {}

	if cfg.Store == "memory" {
		log.Println("Mode memory: data hilang saat server berhenti")
		repo = NewBukuRepoMemori()
	} else {
		db, err := sql.Open("mysql", cfg.DSN())
		if err != nil {
			log.Fatal(err)
		}
		if err := db.Ping(); err != nil {
			log.Fatalf("Tidak bisa konek ke MySQL: %v\nCek: MySQL menyala? schema.sql sudah dijalankan? isi .env benar?", err)
		}
		log.Println("Terhubung ke MySQL")
		repo = NewBukuRepoSQL(db)
		tutup = func() { db.Close() }
	}

	service := NewBukuService(repo)
	handler := NewBukuHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/ping", func(w http.ResponseWriter, r *http.Request) {
		tulisJSON(w, http.StatusOK, Respons{Data: "pong"})
	})
	mux.HandleFunc("GET /api/buku", handler.Semua)
	mux.HandleFunc("GET /api/buku/{id}", handler.Ambil)
	mux.HandleFunc("POST /api/buku", handler.Tambah)
	mux.HandleFunc("DELETE /api/buku/{id}", handler.Hapus)
	return mux, tutup
}

func main() {
	cfg := muatKonfigurasi()
	app, tutup := initApp(cfg)
	defer tutup()

	log.Println("Server berjalan di http://localhost:" + cfg.Port + "  (coba /api/ping, /api/buku)")
	log.Fatal(http.ListenAndServe(":"+cfg.Port, app))
}
