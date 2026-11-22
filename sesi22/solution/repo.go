package main

import "errors"

// ErrTidakDitemukan dikembalikan repository bila baris tidak ada.
// Handler mengubahnya menjadi HTTP 404.
var ErrTidakDitemukan = errors.New("buku tidak ditemukan")

// BukuRepo adalah KONTRAK akses data. Service hanya mengenal interface ini,
// bukan MySQL -- jadi implementasinya bisa ditukar (SQL, memori, mock tes).
type BukuRepo interface {
	Semua() ([]Buku, error)
	Ambil(id int64) (Buku, error)
	Simpan(b Buku) (Buku, error)
	Hapus(id int64) error // bonus Latihan 1
}
