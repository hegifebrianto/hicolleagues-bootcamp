package repository

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"

	"project3/model"
)

// SEMUA query SQL hanya boleh ada di file ini.

type BukuSQL struct{ DB *sql.DB }

func NewBukuSQL(db *sql.DB) *BukuSQL { return &BukuSQL{DB: db} }

func (r *BukuSQL) Semua() ([]model.Buku, error) {
	rows, err := r.DB.Query("SELECT id, judul, penulis, tahun FROM buku ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	hasil := []model.Buku{}
	for rows.Next() {
		var b model.Buku
		if err := rows.Scan(&b.ID, &b.Judul, &b.Penulis, &b.Tahun); err != nil {
			return nil, err
		}
		hasil = append(hasil, b)
	}
	return hasil, rows.Err()
}

func (r *BukuSQL) Ambil(id int64) (model.Buku, error) {
	var b model.Buku
	err := r.DB.QueryRow("SELECT id, judul, penulis, tahun FROM buku WHERE id = ?", id).
		Scan(&b.ID, &b.Judul, &b.Penulis, &b.Tahun)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Buku{}, model.ErrTidakDitemukan
	}
	return b, err
}

func (r *BukuSQL) Simpan(b model.Buku) (model.Buku, error) {
	hasil, err := r.DB.Exec("INSERT INTO buku (judul, penulis, tahun) VALUES (?, ?, ?)",
		b.Judul, b.Penulis, b.Tahun)
	if err != nil {
		return model.Buku{}, err
	}
	b.ID, _ = hasil.LastInsertId()
	return b, nil
}

func (r *BukuSQL) Hapus(id int64) error {
	hasil, err := r.DB.Exec("DELETE FROM buku WHERE id = ?", id)
	if err != nil {
		return err
	}
	if n, _ := hasil.RowsAffected(); n == 0 {
		return model.ErrTidakDitemukan
	}
	return nil
}

type UserSQL struct{ DB *sql.DB }

func NewUserSQL(db *sql.DB) *UserSQL { return &UserSQL{DB: db} }

func (r *UserSQL) CariEmail(email string) (model.User, error) {
	var u model.User
	err := r.DB.QueryRow("SELECT id, nama, email, password_hash, role FROM users WHERE email = ?", email).
		Scan(&u.ID, &u.Nama, &u.Email, &u.PasswordHash, &u.Role)
	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, model.ErrTidakDitemukan
	}
	return u, err
}

func (r *UserSQL) Simpan(u model.User) (model.User, error) {
	hasil, err := r.DB.Exec("INSERT INTO users (nama, email, password_hash, role) VALUES (?, ?, ?, ?)",
		u.Nama, u.Email, u.PasswordHash, u.Role)
	if err != nil {
		return model.User{}, err
	}
	u.ID, _ = hasil.LastInsertId()
	return u, nil
}
