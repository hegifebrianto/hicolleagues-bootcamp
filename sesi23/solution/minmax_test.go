package main

import "testing"

// LATIHAN 1: versi sederhana (satu kasus per fungsi)
func TestMinMaxNormal(t *testing.T) {
	kecil, besar := minMax([]int{4, 9, 2, 7})
	if kecil != 2 || besar != 9 {
		t.Errorf("dapat %d,%d ingin 2,9", kecil, besar)
	}
}

func TestMinMaxSatuElemen(t *testing.T) {
	kecil, besar := minMax([]int{5})
	if kecil != 5 || besar != 5 {
		t.Errorf("dapat %d,%d ingin 5,5", kecil, besar)
	}
}

// LATIHAN 2: table-driven (nama sub-test: TestMinMax/normal, TestMinMax/negatif, ...)
func TestMinMax(t *testing.T) {
	kasus := []struct {
		nama         string
		in           []int
		kecil, besar int
	}{
		{"normal", []int{4, 9, 2, 7}, 2, 9},
		{"satu elemen", []int{5}, 5, 5},
		{"negatif", []int{-3, 0, 8}, -3, 8},
		{"duplikat", []int{3, 3, 1, 1}, 1, 3},   // bonus
		{"semua sama", []int{7, 7, 7, 7}, 7, 7}, // kasus tepi
		{"kosong", []int{}, 0, 0},               // bonus: slice kosong
		{"nil", nil, 0, 0},                      // kasus tepi
	}
	for _, k := range kasus {
		t.Run(k.nama, func(t *testing.T) {
			kecil, besar := minMax(k.in)
			if kecil != k.kecil || besar != k.besar {
				t.Errorf("dapat %d,%d ingin %d,%d", kecil, besar, k.kecil, k.besar)
			}
		})
	}
}
