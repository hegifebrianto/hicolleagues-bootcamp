// LATIHAN 3 (solusi): penghitung yang aman. Selalu 1000, -race bersih.
//
//	go run -race ./latihan/lat3_race
//
// Bonus: versi channel (tanpa Mutex) -- satu goroutine "pemilik" total,
// yang lain hanya mengirim pesan.  "Share memory by communicating".
package main

import (
	"fmt"
	"sync"
)

func hitung(n int) int {
	var (
		mu    sync.Mutex
		wg    sync.WaitGroup
		total int
	)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			total++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return total
}

func hitungDenganChannel(n int) int {
	tambah := make(chan struct{})
	selesai := make(chan int)

	go func() { // satu-satunya yang menyentuh `total`
		total := 0
		for range tambah {
			total++
		}
		selesai <- total
	}()

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tambah <- struct{}{}
		}()
	}
	wg.Wait()
	close(tambah)
	return <-selesai
}

func main() {
	fmt.Println("Mutex  :", hitung(1000))
	fmt.Println("Channel:", hitungDenganChannel(1000))
}
