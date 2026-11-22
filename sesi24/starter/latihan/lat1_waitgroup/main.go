// LATIHAN 1 (8 menit): Goroutine & WaitGroup
// Jalankan 5 goroutine yang masing-masing mencetak nomornya (1..5),
// lalu tunggu semuanya selesai dengan sync.WaitGroup.
//
// Hasil: 5 baris tercetak (urutan boleh beda tiap dijalankan).
// Petunjuk: wg.Add(1) SEBELUM `go func`, dan `defer wg.Done()` di dalam goroutine.
// Bonus: ukur total waktu dengan time.Now() dan time.Since(mulai).
//
// Jalankan:  go run ./latihan/lat1_waitgroup
package main

import "fmt"

func main() {
	// TODO: deklarasikan var wg sync.WaitGroup  (jangan lupa import "sync")
	// TODO: for i := 1; i <= 5; i++ { wg.Add(1); go func(n int) { ... }(i) }
	// TODO: wg.Wait()

	fmt.Println("TODO: belum dikerjakan")
}
