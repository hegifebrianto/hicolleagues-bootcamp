package main

// minMax mengembalikan nilai terkecil dan terbesar dari slice (fungsi dari Sesi 4).
// Slice kosong -> 0, 0 (keputusan desain; ini yang diuji di kasus bonus).
func minMax(angka []int) (int, int) {
	if len(angka) == 0 {
		return 0, 0
	}
	kecil, besar := angka[0], angka[0]
	for _, n := range angka[1:] {
		if n < kecil {
			kecil = n
		}
		if n > besar {
			besar = n
		}
	}
	return kecil, besar
}
