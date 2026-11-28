// Package service: aturan bisnis. Tidak tahu HTTP, tidak tahu SQL.
package service

import (
	"strings"

	"project3/model"
	"project3/repository"
)

type BukuService struct{ Repo repository.BukuRepo }

func NewBukuService(r repository.BukuRepo) *BukuService { return &BukuService{Repo: r} }

func (s *BukuService) Semua() ([]model.Buku, error)       { return s.Repo.Semua() }
func (s *BukuService) Ambil(id int64) (model.Buku, error) { return s.Repo.Ambil(id) }
func (s *BukuService) Hapus(id int64) error               { return s.Repo.Hapus(id) }

func (s *BukuService) Tambah(b model.Buku) (model.Buku, error) {
	b.Judul = strings.TrimSpace(b.Judul)
	b.Penulis = strings.TrimSpace(b.Penulis)
	if b.Judul == "" {
		return model.Buku{}, model.ErrValidasi{Field: "judul", Pesan: "wajib diisi"}
	}
	if b.Tahun <= 1900 {
		return model.Buku{}, model.ErrValidasi{Field: "tahun", Pesan: "harus lebih dari 1900"}
	}
	semua, err := s.Repo.Semua()
	if err != nil {
		return model.Buku{}, err
	}
	for _, x := range semua {
		if strings.EqualFold(x.Judul, b.Judul) {
			return model.Buku{}, model.ErrValidasi{Field: "judul", Pesan: "sudah ada"}
		}
	}
	return s.Repo.Simpan(b)
}
