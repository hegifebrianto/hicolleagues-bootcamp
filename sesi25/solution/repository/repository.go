// Package repository: akses data. Service hanya mengenal INTERFACE di sini.
package repository

import "project3/model"

type BukuRepo interface {
	Semua() ([]model.Buku, error)
	Ambil(id int64) (model.Buku, error)
	Simpan(b model.Buku) (model.Buku, error)
	Hapus(id int64) error
}

type UserRepo interface {
	CariEmail(email string) (model.User, error) // tidak ada -> model.ErrTidakDitemukan
	Simpan(u model.User) (model.User, error)
}
