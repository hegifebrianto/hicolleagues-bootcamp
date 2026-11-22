package main

import (
	"database/sql"
	"errors"
)

// bukuRepoSQL: implementasi BukuRepo dengan database/sql + MySQL.
// SEMUA query SQL hanya boleh ada di file ini.
type bukuRepoSQL struct{ DB *sql.DB }

// NewBukuRepoSQL -- pola factory.
func NewBukuRepoSQL(db *sql.DB) BukuRepo { return bukuRepoSQL{DB: db} }

func (r bukuRepoSQL) Semua() ([]Buku, error) {
	rows, err := r.DB.Query("SELECT id, judul, penulis, tahun FROM buku ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hasil := []Buku{} // bukan nil, supaya JSON "[]"
	for rows.Next() {
		var b Buku
		if err := rows.Scan(&b.ID, &b.Judul, &b.Penulis, &b.Tahun); err != nil {
			return nil, err
		}
		hasil = append(hasil, b)
	}
	return hasil, rows.Err()
}

func (r bukuRepoSQL) Ambil(id int64) (Buku, error) {
	var b Buku
	err := r.DB.QueryRow("SELECT id, judul, penulis, tahun FROM buku WHERE id = ?", id).
		Scan(&b.ID, &b.Judul, &b.Penulis, &b.Tahun)
	if errors.Is(err, sql.ErrNoRows) {
		return Buku{}, ErrTidakDitemukan
	}
	return b, err
}

func (r bukuRepoSQL) Simpan(b Buku) (Buku, error) {
	hasil, err := r.DB.Exec("INSERT INTO buku (judul, penulis, tahun) VALUES (?, ?, ?)",
		b.Judul, b.Penulis, b.Tahun)
	if err != nil {
		return Buku{}, err
	}
	b.ID, _ = hasil.LastInsertId() // MySQL tidak punya RETURNING
	return b, nil
}

func (r bukuRepoSQL) Hapus(id int64) error {
	hasil, err := r.DB.Exec("DELETE FROM buku WHERE id = ?", id)
	if err != nil {
		return err
	}
	if n, _ := hasil.RowsAffected(); n == 0 {
		return ErrTidakDitemukan
	}
	return nil
}
