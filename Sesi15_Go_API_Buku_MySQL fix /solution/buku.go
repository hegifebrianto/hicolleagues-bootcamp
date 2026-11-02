package main

import (
	"encoding/json"
	"fmt"
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

// Respons -- bentuk seragam untuk semua balasan API. `omitempty`
// membuat field yang kosong tidak ikut ditulis ke JSON (jadi respons
// sukses tidak punya key "error" yang kosong, dan sebaliknya).
type Respons struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// tulisJSON -- SELALU dipakai untuk membalas request, supaya urutan
// Header -> WriteHeader -> Encode konsisten di semua handler (lihat
// slide "Response: JSON & Status Code" -- urutan ini WAJIB, kalau
// kebalik Content-Type tidak akan terpasang dengan benar).
func tulisJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func tulisError(w http.ResponseWriter, status int, pesan string) {
	tulisJSON(w, status, Respons{Error: pesan})
}

// GET /api/buku
// GET /api/buku?tahun_min=2020  (bonus: query parameter, sesuai
// contoh "WHERE tahun > 2020" di slide Latihan 1)
func daftarBuku(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, judul, penulis, tahun FROM buku "
	args := []interface{}{}

	tahunMin := r.URL.Query().Get("tahun_min")
	fmt.Println(tahunMin)
	if tahunMin != "" {
		query += " WHERE tahun > ?"
		args = append(args, tahunMin)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		tulisError(w, http.StatusInternalServerError, "gagal membaca data")
		return
	}
	defer rows.Close()

	hasil := []Buku{} // slice kosong (bukan nil) supaya JSON-nya "[]", bukan "null"
	for rows.Next() {
		var b Buku
		if err := rows.Scan(&b.ID, &b.Judul, &b.Penulis, &b.Tahun); err != nil {
			tulisError(w, http.StatusInternalServerError, "gagal membaca baris")
			return
		}
		hasil = append(hasil, b)
	}

	tulisJSON(w, http.StatusOK, Respons{Data: hasil})
}

// POST /api/buku
func tambahBuku(w http.ResponseWriter, r *http.Request) {
	var b Buku
	fmt.Println(r, "request")
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		tulisError(w, http.StatusBadRequest, "body tidak valid")
		return
	}
	fmt.Println(b, "request body")
	if b.Judul == "" {
		tulisError(w, http.StatusBadRequest, "judul wajib diisi")
		return
	}

	hasil, err := db.Exec(
		"INSERT INTO buku (judul, penulis, tahun) VALUES (?, ?, ?)",
		b.Judul, b.Penulis, b.Tahun,
	)
	if err != nil {
		tulisError(w, http.StatusInternalServerError, "gagal menyimpan")
		return
	}

	id, _ := hasil.LastInsertId() // MySQL tidak punya RETURNING -- ini penggantinya
	b.ID = id

	tulisJSON(w, http.StatusCreated, Respons{Data: b})
}

// PUT /api/buku/{id}  (dikerjakan sebagai Tugas Mandiri di slide penutup)
func ubahBuku(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Println(id, "id")
	var b Buku
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		tulisError(w, http.StatusBadRequest, "body tidak valid")
		return
	}
	if b.Judul == "" {
		tulisError(w, http.StatusBadRequest, "judul wajib diisi")
		return
	}

	hasil, err := db.Exec(
		"UPDATE buku SET judul = ?, penulis = ?, tahun = ? WHERE id = ?",
		b.Judul, b.Penulis, b.Tahun, id,
	)
	if err != nil {
		tulisError(w, http.StatusInternalServerError, "gagal mengubah")
		return
	}

	n, _ := hasil.RowsAffected()
	if n == 0 {
		tulisError(w, http.StatusNotFound, "buku tidak ditemukan")
		return
	}

	tulisJSON(w, http.StatusOK, Respons{Data: "berhasil diubah"})
}

// DELETE /api/buku/{id}
func hapusBuku(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	hasil, err := db.Exec("DELETE FROM buku WHERE id = ?", id)
	if err != nil {
		tulisError(w, http.StatusInternalServerError, "gagal menghapus")
		return
	}

	n, err := hasil.RowsAffected()
	if err != nil {
		tulisError(w, http.StatusInternalServerError, "gagal memeriksa hasil")
		return
	}
	if n == 0 {
		tulisError(w, http.StatusNotFound, "buku tidak ditemukan")
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 -- sukses, tidak ada body
}
