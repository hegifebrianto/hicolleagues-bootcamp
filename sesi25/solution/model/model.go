// Package model berisi tipe data yang dipakai SEMUA lapisan.
package model

import (
	"errors"
	"fmt"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type Buku struct {
	ID      int64  `json:"id"`
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun   int    `json:"tahun"`
}

type User struct {
	ID           int64  `json:"id"`
	Nama         string `json:"nama"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"` // tanda "-" : TIDAK PERNAH ikut ke JSON
	Role         string `json:"role"`
}

// Error yang dibuat service/repository. Handler mengubahnya jadi status HTTP.
var (
	ErrTidakDitemukan = errors.New("data tidak ditemukan")      // -> 404
	ErrEmailTerpakai  = errors.New("email sudah terdaftar")     // -> 409
	ErrKredensial     = errors.New("email atau password salah") // -> 401
)

// ErrValidasi dibuat service saat aturan bisnis dilanggar. -> 400
type ErrValidasi struct {
	Field string
	Pesan string
}

func (e ErrValidasi) Error() string { return fmt.Sprintf("%s %s", e.Field, e.Pesan) }
