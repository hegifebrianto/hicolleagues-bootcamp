// KONDISI "SEBELUM": semua menumpuk di handler (SQL + validasi + HTTP).
// Ini kode yang akan Anda refaktor di Latihan 1-3. Perilaku API harus tetap sama.
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Buku struct {
	ID      int64  `json:"id"`
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun   int    `json:"tahun"`
}

var db *sql.DB // global -- sulit diganti saat tes

func tulis(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func semua(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, judul, penulis, tahun FROM buku ORDER BY id") // SQL di handler
	if err != nil {
		tulis(w, 500, map[string]string{"error": err.Error()}) // bocorkan detail error ke klien!
		return
	}
	defer rows.Close()
	hasil := []Buku{}
	for rows.Next() {
		var b Buku
		rows.Scan(&b.ID, &b.Judul, &b.Penulis, &b.Tahun)
		hasil = append(hasil, b)
	}
	tulis(w, 200, map[string]any{"data": hasil})
}

func tambah(w http.ResponseWriter, r *http.Request) {
	var b Buku
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		tulis(w, 400, map[string]string{"error": "body tidak valid"})
		return
	}
	if b.Judul == "" { // validasi (aturan bisnis) di handler
		tulis(w, 400, map[string]string{"error": "judul wajib diisi"})
		return
	}
	if b.Tahun <= 1900 {
		tulis(w, 400, map[string]string{"error": "tahun tidak valid"})
		return
	}
	res, err := db.Exec("INSERT INTO buku (judul, penulis, tahun) VALUES (?, ?, ?)", b.Judul, b.Penulis, b.Tahun)
	if err != nil {
		tulis(w, 500, map[string]string{"error": err.Error()})
		return
	}
	b.ID, _ = res.LastInsertId()
	tulis(w, 201, map[string]any{"data": b})
}

func hapus(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		tulis(w, 400, map[string]string{"error": "id tidak valid"})
		return
	}
	res, err := db.Exec("DELETE FROM buku WHERE id = ?", id)
	if err != nil {
		tulis(w, 500, map[string]string{"error": err.Error()})
		return
	}
	if n, _ := res.RowsAffected(); n == 0 {
		tulis(w, 404, map[string]string{"error": "buku tidak ditemukan"})
		return
	}
	w.WriteHeader(204)
}

func main() {
	var err error
	// Kredensial di-hardcode di kode -- jelek! (diperbaiki di Latihan 3 / tugas mandiri)
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/bukudb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("tidak bisa konek ke MySQL: ", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/buku", semua)
	mux.HandleFunc("POST /api/buku", tambah)
	mux.HandleFunc("DELETE /api/buku/{id}", hapus)
	log.Println("Server (monolit) di http://localhost:8089")
	log.Fatal(http.ListenAndServe(":8089", mux))
}
