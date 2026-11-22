package main

import "testing"

// Jalankan:  go test -race ./latihan/lat3_race
func TestHitung(t *testing.T) {
	for _, k := range []struct {
		nama string
		f    func(int) int
	}{
		{"mutex", hitung},
		{"channel", hitungDenganChannel},
	} {
		t.Run(k.nama, func(t *testing.T) {
			if got := k.f(1000); got != 1000 {
				t.Errorf("dapat %d, ingin 1000", got)
			}
		})
	}
}
