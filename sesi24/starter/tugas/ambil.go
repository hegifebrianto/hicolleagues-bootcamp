// Package tugas: TUGAS MANDIRI -- mengambil beberapa data secara konkuren.
//
// Skenario: halaman ringkasan perlu data banyak buku, tiap pengambilan lambat.
// Berurutan: total waktu = jumlah semua. Konkuren: total waktu ~ yang paling lambat.
package tugas

import (
	"context"
	"errors"
)

type Buku struct {
	ID    int
	Judul string
}

// Sumber mengambil satu buku berdasarkan id. Bisa lambat dan bisa gagal.
// Harus menghormati ctx (berhenti bila dibatalkan).
type Sumber func(ctx context.Context, id int) (Buku, error)

// AmbilBerurutan: sudah jadi, sebagai pembanding (satu per satu).
func AmbilBerurutan(ctx context.Context, ambil Sumber, ids []int) ([]Buku, error) {
	hasil := make([]Buku, 0, len(ids))
	for _, id := range ids {
		b, err := ambil(ctx, id)
		if err != nil {
			return nil, err
		}
		hasil = append(hasil, b)
	}
	return hasil, nil
}

// AmbilKonkuren: TODO -- kerjakan!
//
// Syarat (semuanya diperiksa oleh ambil_test.go):
//  1. Ambil semua id BERSAMAAN (satu goroutine per id) -> jauh lebih cepat.
//  2. Hasil TETAP berurutan sesuai `ids`.
//     Petunjuk: hasil := make([]Buku, len(ids)); tiap goroutine menulis hasil[i] miliknya sendiri.
//  3. Error pertama dikembalikan, dan pekerjaan lain dibatalkan.
//     Petunjuk: ctx, batal := context.WithCancel(ctx); panggil batal() saat ada error.
//     Variabel penampung error ditulis banyak goroutine -> lindungi dengan sync.Mutex.
//  4. Harus lulus:  go test -race ./tugas
func AmbilKonkuren(ctx context.Context, ambil Sumber, ids []int) ([]Buku, error) {
	return nil, errors.New("TODO: AmbilKonkuren belum dikerjakan")
}
