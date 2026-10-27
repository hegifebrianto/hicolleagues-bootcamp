package tugas

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("tugas tidak ditemukan")

type Tugas struct {
	ID      int    `json:"id"`
	Judul   string `json:"judul"`
	Selesai bool   `json:"selesai"`
}

type MemoryRepository struct {
	mu     sync.RWMutex
	nextID int
	items  []Tugas
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{nextID: 1, items: []Tugas{}}
}

func (r *MemoryRepository) List() []Tugas {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]Tugas, len(r.items))
	copy(result, r.items)
	return result
}

func (r *MemoryRepository) Create(judul string) Tugas {
	r.mu.Lock()
	defer r.mu.Unlock()

	item := Tugas{ID: r.nextID, Judul: judul, Selesai: false}
	r.nextID++
	r.items = append(r.items, item)
	return item
}

func (r *MemoryRepository) Get(id int) (Tugas, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, item := range r.items {
		if item.ID == id {
			return item, nil
		}
	}
	return Tugas{}, ErrNotFound
}

func (r *MemoryRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, item := range r.items {
		if item.ID == id {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
