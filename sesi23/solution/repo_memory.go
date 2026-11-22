package main

import "sync"

// bukuRepoMemori: implementasi BukuRepo di MEMORI (tanpa database).
// Dipakai untuk (1) unit test service, (2) demo tanpa MySQL: STORE=memory go run .
// Inilah manfaat nyata dari interface: service tidak sadar bedanya.
type bukuRepoMemori struct {
	mu     sync.Mutex
	data   []Buku
	nextID int64
}

func NewBukuRepoMemori() BukuRepo {
	return &bukuRepoMemori{
		nextID: 4,
		data: []Buku{
			{1, "Laskar Pelangi", "Andrea Hirata", 2005},
			{2, "Bumi Manusia", "Pramoedya Ananta Toer", 1980},
			{3, "Filosofi Teras", "Henry Manampiring", 2019},
		},
	}
}

func (r *bukuRepoMemori) Semua() ([]Buku, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]Buku{}, r.data...), nil
}

func (r *bukuRepoMemori) Ambil(id int64) (Buku, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, b := range r.data {
		if b.ID == id {
			return b, nil
		}
	}
	return Buku{}, ErrTidakDitemukan
}

func (r *bukuRepoMemori) Simpan(b Buku) (Buku, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	b.ID = r.nextID
	r.nextID++
	r.data = append(r.data, b)
	return b, nil
}

func (r *bukuRepoMemori) Hapus(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, b := range r.data {
		if b.ID == id {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return ErrTidakDitemukan
}
