// Package matematika: fungsi murni sederhana untuk belajar unit test.
package matematika

import "errors"

// ErrBagiNol dikembalikan Bagi bila pembagi 0.
var ErrBagiNol = errors.New("tidak bisa membagi dengan nol")

func Tambah(a, b int) int { return a + b }

func Bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrBagiNol
	}
	return a / b, nil
}
