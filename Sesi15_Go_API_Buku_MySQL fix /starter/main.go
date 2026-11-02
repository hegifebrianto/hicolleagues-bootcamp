package main

import (
	"log"
	"net/http"
)

func main() {
	initDB()
	defer db.Close()

	mux := http.NewServeMux()

	// TODO 1: daftarkan 4 route ke mux, pola: mux.HandleFunc("METODE /path", namaFungsi)
	//   GET    /api/buku       -> daftarBuku
	//   POST   /api/buku       -> tambahBuku
	//   PUT    /api/buku/{id}  -> ubahBuku
	//   DELETE /api/buku/{id}  -> hapusBuku
	// (Sebelum TODO ini dikerjakan, server tetap bisa dijalankan --
	// cuma semua endpoint akan membalas 404 karena belum ada route
	// yang terdaftar.)

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
