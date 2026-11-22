package matematika

import "testing"

// Test pertama (kita tulis bersama di kelas)
func TestTambah(t *testing.T) {
	hasil := Tambah(2, 3)
	if hasil != 5 {
		t.Errorf("Tambah(2, 3) = %d, ingin 5", hasil)
	}
}

// TODO (latihan di kelas): tulis TestBagi -- pakai t.Fatal bila error tidak muncul saat Bagi(10, 0)
// TODO: tulis TestTambahBerbagaiKasus sebagai table-driven (positif, negatif, nol) dengan t.Run
