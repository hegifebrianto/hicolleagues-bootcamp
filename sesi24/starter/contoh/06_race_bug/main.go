// DEMO BUG: penghitung TANPA Mutex. Hasilnya tidak konsisten (hampir selalu < 1000).
//
//	go run ./contoh/06_race_bug          -> jalankan 3x, angkanya berubah-ubah
//	go run -race ./contoh/06_race_bug    -> "WARNING: DATA RACE" + lokasi baris konflik
//
// Penyebab: total++ sebenarnya 3 langkah (baca, tambah, tulis). Dua goroutine bisa
// membaca nilai yang sama lalu saling menimpa -> satu tambahan "hilang".
//
// Catatan: pada komputer cepat, `total++` polos kadang tetap menghasilkan 1000 karena
// tiap goroutine selesai sangat cepat -- bug-nya tetap ada, hanya jarang tampak
// (-race selalu melaporkannya). Di sini 3 langkah itu dipisah dan diberi runtime.Gosched()
// (beri kesempatan goroutine lain jalan) agar bug-nya tampak jelas dan bisa didemokan.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var total int

func tambah() {
	sementara := total    // 1. baca
	runtime.Gosched()     // goroutine lain bisa menyela di sini
	total = sementara + 1 // 3. tulis (menimpa hasil goroutine lain)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); tambah() }()
	}
	wg.Wait()
	fmt.Println(total)
}
