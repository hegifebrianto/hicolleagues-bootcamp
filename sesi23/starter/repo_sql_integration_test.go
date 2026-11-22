//go:build integration

package main

import (
	"database/sql"
	"errors"
	"testing"
)

// Integration test: menguji repo SQL dengan database MySQL NYATA.
// Hanya jalan bila memakai build tag:   go test -tags=integration ./...
// Memakai database TERPISAH (bukudb_test) -- buat dulu dengan schema_test.sql.

func bukaDBTes(t *testing.T) *sql.DB {
	t.Helper()
	cfg := muatKonfigurasi()
	cfg.DBName = getEnvDefault("DB_NAME_TEST", "bukudb_test")

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		db.Close()
		t.Skipf("MySQL test tidak tersedia (%v) -- jalankan schema_test.sql & cek .env", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

func TestRepoSQL_SemuaTanpaError(t *testing.T) {
	repo := bukuRepoSQL{DB: bukaDBTes(t)}
	if _, err := repo.Semua(); err != nil {
		t.Fatal(err)
	}
}

// Siklus lengkap: simpan -> ambil -> ada di daftar -> hapus -> hilang
func TestRepoSQL_Siklus(t *testing.T) {
	repo := bukuRepoSQL{DB: bukaDBTes(t)}

	b, err := repo.Simpan(Buku{Judul: "Buku Integration Test", Penulis: "Tester", Tahun: 2024})
	if err != nil {
		t.Fatal(err)
	}
	if b.ID == 0 {
		t.Fatal("ID harusnya terisi dari LastInsertId")
	}
	t.Cleanup(func() { repo.Hapus(b.ID) }) // bersihkan walau test gagal

	got, err := repo.Ambil(b.ID)
	if err != nil || got.Judul != "Buku Integration Test" {
		t.Fatalf("Ambil: %+v, %v", got, err)
	}

	semua, err := repo.Semua()
	if err != nil {
		t.Fatal(err)
	}
	ada := false
	for _, x := range semua {
		if x.ID == b.ID {
			ada = true
		}
	}
	if !ada {
		t.Error("buku yang baru disimpan tidak ada di Semua()")
	}

	if err := repo.Hapus(b.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.Ambil(b.ID); !errors.Is(err, ErrTidakDitemukan) {
		t.Errorf("setelah dihapus harusnya ErrTidakDitemukan, dapat %v", err)
	}
	if err := repo.Hapus(b.ID); !errors.Is(err, ErrTidakDitemukan) {
		t.Errorf("hapus kedua harusnya ErrTidakDitemukan, dapat %v", err)
	}
}
