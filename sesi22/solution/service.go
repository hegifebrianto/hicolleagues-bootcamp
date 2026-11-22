package main

import (
	"strings"
	"time"
)

// BukuService berisi ATURAN BISNIS. Tidak ada SQL, tidak ada net/http di sini.
type BukuService struct {
	Repo BukuRepo // dependency, disuntik lewat konstruktor
}

// NewBukuService -- pola factory + dependency injection lewat konstruktor.
func NewBukuService(repo BukuRepo) BukuService { return BukuService{Repo: repo} }

func (s BukuService) Semua() ([]Buku, error)       { return s.Repo.Semua() }
func (s BukuService) Ambil(id int64) (Buku, error) { return s.Repo.Ambil(id) }
func (s BukuService) Hapus(id int64) error         { return s.Repo.Hapus(id) }

func (s BukuService) Tambah(b Buku) (Buku, error) {
	b.Judul = strings.TrimSpace(b.Judul)
	b.Penulis = strings.TrimSpace(b.Penulis)

	if b.Judul == "" {
		return Buku{}, ErrValidasi{"judul", "wajib diisi"}
	}
	if b.Tahun <= 1900 || b.Tahun > time.Now().Year()+1 {
		return Buku{}, ErrValidasi{"tahun", "tidak valid"}
	}

	// Bonus Latihan 2: tolak judul yang sudah ada (tanpa membedakan huruf besar/kecil)
	semua, err := s.Repo.Semua()
	if err != nil {
		return Buku{}, err
	}
	for _, ada := range semua {
		if strings.EqualFold(ada.Judul, b.Judul) {
			return Buku{}, ErrValidasi{"judul", "sudah ada"}
		}
	}
	return s.Repo.Simpan(b)
}
