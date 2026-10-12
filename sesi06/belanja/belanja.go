package belanja

import (
	"fmt"
)

// DoPrint diawali huruf besar agar ter-export (bisa dipanggil dari package lain).
func DoPrint() string {
	return "running in module belanja"
}

type Item struct {
	id       int
	nama     string
	harga    float64
	quantity int
}

var items []Item
var nextId = 1

func Tambah(nama string, harga float64, quantity int) (Item, error) {
	fmt.Println(nama, harga, quantity)
	item := Item{
		id:       nextId,
		nama:     nama,
		harga:    harga,
		quantity: quantity,
	}
	nextId++

	// return item, nil

	items = append(items, item)
	return item, nil
}

func Update(id, quantity int) error {
	fmt.Println(id, quantity)
	for i := range items {
		fmt.Println(items[i].id == id)

		if items[i].id == id {
			items[i].quantity = quantity
			return nil
		}
	}
	return fmt.Errorf("item %d tidak ditemukan", id)

}

func delete(id int) {
	fmt.Println(id)
}

func show() []Item {
	fmt.Println(items)
	return nil
}

func total() {
	total := 0
	fmt.Println(total)
}
