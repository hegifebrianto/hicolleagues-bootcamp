// Package tugas: TUGAS MANDIRI -- mengambil beberapa data secara konkuren.
//
// Skenario: halaman ringkasan perlu data banyak buku. Tiap pengambilan lambat
// (mis. query ke database/API lain). Berurutan: total waktu = jumlah semua.
// Konkuren: total waktu ~ yang paling lambat.
package tugas

import (
	"context"
	"sync"
)

type Buku struct {
	ID    int
	Judul string
}

// Sumber mengambil satu buku berdasarkan id. Bisa lambat dan bisa gagal.
// Harus menghormati ctx (berhenti bila dibatalkan).
type Sumber func(ctx context.Context, id int) (Buku, error)

// AmbilBerurutan: pembanding (satu per satu).
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

// AmbilKonkuren mengambil semua id secara bersamaan.
//   - hasil TETAP berurutan sesuai `ids`
//   - error pertama membatalkan pekerjaan lainnya (lewat context)
//   - aman dari race: tiap goroutine hanya menulis ke indeks miliknya sendiri
func AmbilKonkuren(ctx context.Context, ambil Sumber, ids []int) ([]Buku, error) {
	ctx, batal := context.WithCancel(ctx)
	defer batal()

	hasil := make([]Buku, len(ids)) // diisi per indeks -> tidak perlu Mutex untuk ini
	var (
		wg         sync.WaitGroup
		mu         sync.Mutex // melindungi errPertama (ditulis banyak goroutine)
		errPertama error
	)

	for i, id := range ids {
		wg.Add(1)
		go func(i, id int) {
			defer wg.Done()
			b, err := ambil(ctx, id)
			if err != nil {
				mu.Lock()
				if errPertama == nil {
					errPertama = err
					batal() // hentikan yang lain
				}
				mu.Unlock()
				return
			}
			hasil[i] = b
		}(i, id)
	}
	wg.Wait()

	if errPertama != nil {
		return nil, errPertama
	}
	return hasil, nil
}
