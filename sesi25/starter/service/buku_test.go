package service

import (
	"errors"
	"testing"

	"project3/model"
	"project3/repository"
)

func TestTambahBuku(t *testing.T) {
	tests := []struct {
		nama    string
		input   model.Buku
		wantErr bool
	}{
		{"valid", model.Buku{Judul: "Go Dasar", Penulis: "A", Tahun: 2020}, false},
		{"judul kosong", model.Buku{Judul: "  ", Tahun: 2020}, true},
		{"tahun terlalu lama", model.Buku{Judul: "X", Tahun: 1900}, true},
		{"judul duplikat (beda huruf)", model.Buku{Judul: "LASKAR PELANGI", Tahun: 2005}, true},
	}
	for _, tt := range tests {
		t.Run(tt.nama, func(t *testing.T) {
			repo := repository.NewBukuMemori(model.Buku{Judul: "Laskar Pelangi", Tahun: 2005})
			svc := NewBukuService(repo)
			_, err := svc.Tambah(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tt.wantErr)
			}
			if tt.wantErr {
				var ev model.ErrValidasi
				if !errors.As(err, &ev) {
					t.Errorf("harusnya ErrValidasi, dapat %T", err)
				}
			}
		})
	}
}

func TestHapusTidakAda(t *testing.T) {
	svc := NewBukuService(repository.NewBukuMemori())
	if err := svc.Hapus(99); !errors.Is(err, model.ErrTidakDitemukan) {
		t.Errorf("dapat %v", err)
	}
}
