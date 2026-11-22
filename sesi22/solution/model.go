package main

import "fmt"

// Buku adalah model data yang dipakai SEMUA lapisan.
type Buku struct {
	ID      int64  `json:"id"`
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun   int    `json:"tahun"`
}

// Respons: bentuk balasan API yang seragam.
type Respons struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

// ErrValidasi dibuat oleh SERVICE ketika aturan bisnis dilanggar.
// Handler mengubahnya menjadi HTTP 400.
type ErrValidasi struct {
	Field string
	Pesan string
}

func (e ErrValidasi) Error() string { return fmt.Sprintf("%s %s", e.Field, e.Pesan) }
