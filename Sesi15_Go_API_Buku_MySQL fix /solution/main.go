package main

import (
	"log"
	"net/http"
)

func main() {
	initDB()
	defer db.Close()

	// http.NewServeMux() -- router bawaan Go (lihat slide "Routing:
	// ServeMux & Framework"). "GET /api/buku" memakai sintaks method+path
	// yang tersedia sejak Go 1.22 -- tidak perlu library tambahan
	// seperti Gorilla Mux/Gin Gonic untuk routing sesederhana ini.
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/buku", daftarBuku)
	mux.HandleFunc("POST /api/buku", tambahBuku)
	mux.HandleFunc("PUT /api/buku/{id}", ubahBuku)
	mux.HandleFunc("DELETE /api/buku/{id}", hapusBuku)

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
