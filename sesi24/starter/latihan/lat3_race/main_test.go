package main

import "testing"

// Jalankan:  go test -race ./latihan/lat3_race
// Awalnya GAGAL (dan -race melapor DATA RACE); setelah diperbaiki harus PASS.
func TestHitung(t *testing.T) {
	if got := hitung(1000); got != 1000 {
		t.Errorf("dapat %d, ingin 1000", got)
	}
}
