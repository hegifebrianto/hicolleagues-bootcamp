// LATIHAN 3 (8 menit): Perbaiki Race Condition
// Kode di bawah SENGAJA bermasalah: penghitung diubah banyak goroutine tanpa Mutex.
//
// Langkah:
//  1. Jalankan 3x:   go run ./latihan/lat3_race          -> angka berubah-ubah, bukan 1000
//  2. Lihat laporan: go run -race ./latihan/lat3_race    -> WARNING: DATA RACE (baca baris yang disebut)
//  3. Perbaiki dengan sync.Mutex (Lock/Unlock di sekitar penulisan `total`)
//  4. Uji: go test -race ./latihan/lat3_race             -> harus PASS & bersih dari WARNING
//
// Bonus: tulis versi kedua tanpa Mutex (satu goroutine pemilik total + channel).
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hitung(n int) int {
	var (
		wg    sync.WaitGroup
		total int
		// TODO: var mu sync.Mutex
	)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// TODO: kunci (mu.Lock) sebelum membaca/menulis total, buka (mu.Unlock) sesudahnya
			sementara := total    // baca
			runtime.Gosched()     // (disengaja) beri kesempatan goroutine lain menyela
			total = sementara + 1 // tulis
		}()
	}
	wg.Wait()
	return total
}

func main() {
	fmt.Println(hitung(1000)) // harus selalu 1000
}
