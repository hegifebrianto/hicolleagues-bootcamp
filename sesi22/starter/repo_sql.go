package main

import (
	"database/sql"
	"errors"
)

// LATIHAN 1 (8 menit): implementasikan BukuRepo dengan database/sql.
// Pindahkan query dari ../sebelum/main.go ke sini.
type bukuRepoSQL struct{ DB *sql.DB }

func NewBukuRepoSQL(db *sql.DB) BukuRepo { return bukuRepoSQL{DB: db} }

var errBelum = errors.New("TODO: belum diimplementasikan")

func (r bukuRepoSQL) Semua() ([]Buku, error) {
	// TODO: SELECT id, judul, penulis, tahun FROM buku ORDER BY id
	//       (slice hasil dibuat dengan []Buku{}, bukan nil)
	return nil, errBelum
}

func (r bukuRepoSQL) Ambil(id int64) (Buku, error) {
	// TODO: SELECT satu baris; sql.ErrNoRows -> ErrTidakDitemukan
	return Buku{}, errBelum
}

func (r bukuRepoSQL) Simpan(b Buku) (Buku, error) {
	// TODO: INSERT, isi b.ID dari LastInsertId(), kembalikan b
	return Buku{}, errBelum
}

func (r bukuRepoSQL) Hapus(id int64) error {
	// TODO (bonus): DELETE; RowsAffected()==0 -> ErrTidakDitemukan
	return errBelum
}
