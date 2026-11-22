package main

import (
	"errors"
	"testing"
)

// Contoh di slide "Mock dan Stub", disesuaikan dengan aturan service kita
// (tahun harus > 1900, jadi Judul saja tidak cukup).
func TestServiceTambah(t *testing.T) {
	repo := &repoPalsu{}
	service := NewBukuService(repo)

	hasil, err := service.Tambah(Buku{Judul: "Go", Penulis: "Tidak Tahu", Tahun: 2020})
	if err != nil {
		t.Fatal(err)
	}
	if hasil.Judul != "Go" {
		t.Errorf("judul salah: %q", hasil.Judul)
	}
	if repo.simpanDipanggil != 1 { // bonus: hitung panggilan Simpan
		t.Errorf("Simpan dipanggil %d kali, ingin 1", repo.simpanDipanggil)
	}
}

// LATIHAN 3 -- kasus validasi (table-driven + mock)
func TestTambah_Validasi(t *testing.T) {
	kasus := []struct {
		nama    string
		input   Buku
		ada     []Buku // isi awal repo
		wantErr bool
	}{
		{"judul kosong", Buku{Judul: "", Tahun: 2020}, nil, true},
		{"judul spasi saja", Buku{Judul: "   ", Tahun: 2020}, nil, true},
		{"tahun 1800", Buku{Judul: "Kuno", Tahun: 1800}, nil, true},
		{"tahun 1900 (batas)", Buku{Judul: "Batas", Tahun: 1900}, nil, true},
		{"tahun terlalu depan", Buku{Judul: "Masa Depan", Tahun: 3000}, nil, true},
		{"judul duplikat", Buku{Judul: "GO", Tahun: 2020}, []Buku{{ID: 1, Judul: "Go", Tahun: 2019}}, true},
		{"data valid", Buku{Judul: "Buku Baru", Penulis: "A", Tahun: 2020}, nil, false},
		{"tahun 1901 (batas valid)", Buku{Judul: "Batas Valid", Tahun: 1901}, nil, false},
	}
	for _, k := range kasus {
		t.Run(k.nama, func(t *testing.T) {
			repo := &repoPalsu{data: k.ada}
			_, err := NewBukuService(repo).Tambah(k.input)

			var ev ErrValidasi
			if k.wantErr {
				if !errors.As(err, &ev) {
					t.Fatalf("harusnya ErrValidasi, dapat %v", err)
				}
				// Penting: bila validasi gagal, data TIDAK boleh sampai ke repo
				if repo.simpanDipanggil != 0 {
					t.Errorf("Simpan dipanggil %d kali padahal validasi gagal", repo.simpanDipanggil)
				}
				return
			}
			if err != nil {
				t.Fatalf("harusnya sukses, dapat %v", err)
			}
			if repo.simpanDipanggil != 1 {
				t.Errorf("Simpan dipanggil %d kali, ingin 1", repo.simpanDipanggil)
			}
		})
	}
}

func TestTambah_Trim(t *testing.T) {
	repo := &repoPalsu{}
	hasil, err := NewBukuService(repo).Tambah(Buku{Judul: "  Go  ", Penulis: " A ", Tahun: 2020})
	if err != nil {
		t.Fatal(err)
	}
	if hasil.Judul != "Go" || hasil.Penulis != "A" {
		t.Errorf("spasi tidak dibersihkan: %+v", hasil)
	}
}

// Error dari repo harus diteruskan APA ADANYA (bukan ErrValidasi)
func TestTambah_ErrorDariRepo(t *testing.T) {
	boom := errors.New("db mati")

	t.Run("Simpan gagal", func(t *testing.T) {
		_, err := NewBukuService(&repoPalsu{errSimpan: boom}).Tambah(Buku{Judul: "X", Tahun: 2020})
		if !errors.Is(err, boom) {
			t.Fatalf("dapat %v, ingin %v", err, boom)
		}
	})
	t.Run("Semua gagal (cek duplikat)", func(t *testing.T) {
		_, err := NewBukuService(&repoPalsu{errSemua: boom}).Tambah(Buku{Judul: "X", Tahun: 2020})
		if !errors.Is(err, boom) {
			t.Fatalf("dapat %v, ingin %v", err, boom)
		}
	})
}

func TestService_Semua(t *testing.T) {
	s := NewBukuService(&repoPalsu{data: []Buku{{ID: 1, Judul: "A"}, {ID: 2, Judul: "B"}}})
	hasil, err := s.Semua()
	if err != nil || len(hasil) != 2 {
		t.Fatalf("dapat %v, %v", hasil, err)
	}
	if _, err := NewBukuService(&repoPalsu{errSemua: errors.New("x")}).Semua(); err == nil {
		t.Error("harusnya error")
	}
}

func TestService_Ambil(t *testing.T) {
	s := NewBukuService(&repoPalsu{data: []Buku{{ID: 1, Judul: "A"}}})
	if b, err := s.Ambil(1); err != nil || b.Judul != "A" {
		t.Errorf("dapat %v, %v", b, err)
	}
	if _, err := s.Ambil(99); !errors.Is(err, ErrTidakDitemukan) {
		t.Errorf("harusnya ErrTidakDitemukan, dapat %v", err)
	}
}

func TestService_Hapus(t *testing.T) {
	s := NewBukuService(&repoPalsu{data: []Buku{{ID: 1, Judul: "A"}}})
	if err := s.Hapus(1); err != nil {
		t.Errorf("harusnya sukses, dapat %v", err)
	}
	if err := s.Hapus(99); !errors.Is(err, ErrTidakDitemukan) {
		t.Errorf("harusnya ErrTidakDitemukan, dapat %v", err)
	}
}
