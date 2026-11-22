package matematika

import "testing"

// Test pertama (slide "Test Pertama")
func TestTambah(t *testing.T) {
	hasil := Tambah(4, 1)
	if hasil != 5 {
		t.Errorf("Tambah(2, 3) = %d, ingin 5", hasil)
	}
}

// t.Fatal vs t.Error (slide "t.Error vs t.Fatal")
func TestBagi(t *testing.T) {
	hasil, err := Bagi(1, 0)
	if err == nil {
		t.Fatal("harusnya ada error saat bagi nol") // Fatal: berhenti, tak ada gunanya lanjut
	}
	if hasil != 0 {
		t.Errorf("hasil = %v, ingin 0", hasil) // Error: lanjut
	}
}

func TestBagiNormal(t *testing.T) {
	hasil, err := Bagi(10, 5)
	if err != nil {
		t.Fatalf("tidak diduga error: %v", err)
	}
	if hasil != 2 {
		t.Errorf("Bagi(10, 5) = %v, ingin 2", hasil)
	}
}

// Table-driven (slide "Table-Driven Tests")
func TestTambahBerbagaiKasus(t *testing.T) {
	kasus := []struct {
		nama     string
		a, b     int
		expected int
	}{
		{"positif", 2, 3, 5},
		{"negatif", -1, -1, -2},
		{"nol", 0, 5, 5},
	}
	for _, k := range kasus {
		t.Run(k.nama, func(t *testing.T) {
			if hasil := Tambah(k.a, k.b); hasil != k.expected {
				t.Errorf("dapat %d, ingin %d", hasil, k.expected)
			}
		})
	}
}
