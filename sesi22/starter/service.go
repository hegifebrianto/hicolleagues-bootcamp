package main

// LATIHAN 2 (9 menit): service + validasi.
// Aturan: judul tidak kosong, tahun > 1900. Gagal -> kembalikan ErrValidasi (lihat model.go).
type BukuService struct {
	Repo BukuRepo
}

func NewBukuService(repo BukuRepo) BukuService { return BukuService{Repo: repo} }

func (s BukuService) Semua() ([]Buku, error) {
	// TODO: panggil s.Repo.Semua()
	return nil, errBelum
}

func (s BukuService) Ambil(id int64) (Buku, error) {
	// TODO: panggil s.Repo.Ambil(id)
	return Buku{}, errBelum
}

func (s BukuService) Tambah(b Buku) (Buku, error) {
	// TODO: 1) judul kosong -> ErrValidasi{"judul", "wajib diisi"}
	//       2) tahun <= 1900 -> ErrValidasi{"tahun", "tidak valid"}
	//       3) s.Repo.Simpan(b)
	// Bonus: tolak judul duplikat (cek lewat s.Repo.Semua()).
	return Buku{}, errBelum
}

func (s BukuService) Hapus(id int64) error {
	// TODO: panggil s.Repo.Hapus(id)
	return errBelum
}
