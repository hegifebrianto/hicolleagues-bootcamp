package main

import (
	"errors"
	"testing"
)

// Service diuji TANPA server dan TANPA database -- cukup repo memori.
// (Sesi berikutnya membahas testing lebih dalam; ini cuma pembuktian manfaat layering.)
func TestTambah(t *testing.T) {
	kasus := []struct {
		nama    string
		input   Buku
		wantErr bool
	}{
		{"judul kosong", Buku{Judul: "", Penulis: "A", Tahun: 2000}, true},
		{"tahun 1800", Buku{Judul: "Kuno", Penulis: "A", Tahun: 1800}, true},
		{"judul sudah ada", Buku{Judul: "laskar pelangi", Penulis: "A", Tahun: 2000}, true},
		{"data valid", Buku{Judul: "Buku Baru", Penulis: "A", Tahun: 2020}, false},
	}
	for _, k := range kasus {
		t.Run(k.nama, func(t *testing.T) {
			s := NewBukuService(NewBukuRepoMemori())
			_, err := s.Tambah(k.input)
			var ev ErrValidasi
			if k.wantErr && !errors.As(err, &ev) {
				t.Fatalf("harusnya ErrValidasi, dapat %v", err)
			}
			if !k.wantErr && err != nil {
				t.Fatalf("harusnya sukses, dapat %v", err)
			}
		})
	}
}

func TestHapusTidakAda(t *testing.T) {
	s := NewBukuService(NewBukuRepoMemori())
	if err := s.Hapus(999); !errors.Is(err, ErrTidakDitemukan) {
		t.Fatalf("harusnya ErrTidakDitemukan, dapat %v", err)
	}
}
