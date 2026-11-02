package main

import (
	"encoding/json"
	"net/http"
)

// Buku -- field json:"..." menentukan nama key saat struct ini
// diubah jadi/dari JSON.
type Buku struct {
	ID      int64  `json:"id"`
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun   int    `json:"tahun"`
}

// Respons -- bentuk seragam untuk semua balasan API.
type Respons struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// tulisJSON dan tulisError SUDAH LENGKAP -- pakai ini di semua
// handler yang kamu kerjakan di bawah, jangan menulis
// w.Header()/w.WriteHeader()/Encode() manual lagi.
func tulisJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func tulisError(w http.ResponseWriter, status int, pesan string) {
	tulisJSON(w, status, Respons{Error: pesan})
}

// GET /api/buku
func daftarBuku(w http.ResponseWriter, r *http.Request) {
	// TODO 1: jalankan db.Query("SELECT id, judul, penulis, tahun FROM buku").
	// Kalau error, tulisError(w, http.StatusInternalServerError, "gagal membaca data") lalu return.
	// JANGAN LUPA: defer rows.Close() segera setelah berhasil Query.

	// TODO 2: buat hasil := []Buku{} (slice KOSONG, bukan nil, supaya
	// JSON-nya "[]" bukan "null" kalau tabelnya kosong).

	// TODO 3: loop "for rows.Next() { ... }" -- di dalamnya: deklarasikan
	// var b Buku, lalu rows.Scan(&b.ID, &b.Judul, &b.Penulis, &b.Tahun).
	// Kalau Scan error, tulisError(500, "gagal membaca baris") lalu return.
	// Kalau sukses, hasil = append(hasil, b).

	// TODO 4: tulisJSON(w, http.StatusOK, Respons{Data: hasil})
}

// POST /api/buku
func tambahBuku(w http.ResponseWriter, r *http.Request) {
	// TODO 5: deklarasikan var b Buku, lalu
	// json.NewDecoder(r.Body).Decode(&b). Kalau error,
	// tulisError(w, http.StatusBadRequest, "body tidak valid") lalu return.

	// TODO 6: validasi -- kalau b.Judul == "", tulisError(400, "judul wajib diisi") lalu return.

	// TODO 7: db.Exec("INSERT INTO buku (judul, penulis, tahun) VALUES (?, ?, ?)",
	// b.Judul, b.Penulis, b.Tahun). Simpan hasilnya (sql.Result) dan err.
	// Kalau error, tulisError(500, "gagal menyimpan") lalu return.

	// TODO 8: ambil id baris baru dengan hasil.LastInsertId() (MySQL
	// tidak punya RETURNING seperti PostgreSQL -- ini penggantinya),
	// simpan ke b.ID, lalu tulisJSON(w, http.StatusCreated, Respons{Data: b}).
}

// PUT /api/buku/{id}
func ubahBuku(w http.ResponseWriter, r *http.Request) {
	// TODO 9: ambil id dari URL dengan r.PathValue("id").

	// TODO 10: decode body JSON ke Buku (pola sama seperti tambahBuku),
	// termasuk validasi judul tidak boleh kosong.

	// TODO 11: db.Exec("UPDATE buku SET judul = ?, penulis = ?, tahun = ? WHERE id = ?",
	// b.Judul, b.Penulis, b.Tahun, id). Kalau error, tulisError(500, "gagal mengubah").

	// TODO 12: cek hasil.RowsAffected() -- kalau 0, berarti id tidak
	// ada, tulisError(w, http.StatusNotFound, "buku tidak ditemukan").
	// Kalau berhasil, tulisJSON(w, http.StatusOK, Respons{Data: "berhasil diubah"}).
}

// DELETE /api/buku/{id}
func hapusBuku(w http.ResponseWriter, r *http.Request) {
	// TODO 13: ambil id dari URL (r.PathValue("id")), lalu
	// db.Exec("DELETE FROM buku WHERE id = ?", id).

	// TODO 14: sama seperti ubahBuku -- cek RowsAffected(), kalau 0
	// berarti tidak ditemukan (404). Kalau berhasil, panggil
	// w.WriteHeader(http.StatusNoContent) -- TANPA body, karena 204
	// artinya "sukses, tidak ada yang perlu ditampilkan".
}
