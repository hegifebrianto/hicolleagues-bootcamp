package repository

import (
	"sync"

	"project3/model"
)

// ---- Buku di memori (untuk demo tanpa MySQL & untuk test) ----

type BukuMemori struct {
	mu     sync.Mutex
	data   []model.Buku
	nextID int64
}

func NewBukuMemori(awal ...model.Buku) *BukuMemori {
	r := &BukuMemori{nextID: 1}
	for _, b := range awal {
		r.Simpan(b)
	}
	return r
}

func (r *BukuMemori) Semua() ([]model.Buku, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]model.Buku{}, r.data...), nil
}

func (r *BukuMemori) Ambil(id int64) (model.Buku, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, b := range r.data {
		if b.ID == id {
			return b, nil
		}
	}
	return model.Buku{}, model.ErrTidakDitemukan
}

func (r *BukuMemori) Simpan(b model.Buku) (model.Buku, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	b.ID = r.nextID
	r.nextID++
	r.data = append(r.data, b)
	return b, nil
}

func (r *BukuMemori) Hapus(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, b := range r.data {
		if b.ID == id {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return model.ErrTidakDitemukan
}

// ---- User di memori ----

type UserMemori struct {
	mu     sync.Mutex
	data   []model.User
	nextID int64
}

func NewUserMemori() *UserMemori { return &UserMemori{nextID: 1} }

func (r *UserMemori) CariEmail(email string) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, u := range r.data {
		if u.Email == email {
			return u, nil
		}
	}
	return model.User{}, model.ErrTidakDitemukan
}

func (r *UserMemori) Simpan(u model.User) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u.ID = r.nextID
	r.nextID++
	r.data = append(r.data, u)
	return u, nil
}
