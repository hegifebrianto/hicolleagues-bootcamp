// LATIHAN 2: kuadrat 1..5 dihitung di goroutine, hasil lewat channel (buffered 5).
// Hasil: 1 4 9 16 25 (urutan tergantung). Bonus: diurutkan sebelum dicetak.
package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	hasil := make(chan int, 5)

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			hasil <- n * n
		}(i)
	}

	// Penutup channel dijalankan di goroutine sendiri: menunggu semua pengirim
	// selesai, baru close(hasil). Jika tidak, `range` di bawah menunggu selamanya.
	go func() {
		wg.Wait()
		close(hasil)
	}()

	var semua []int
	for h := range hasil {
		semua = append(semua, h)
	}

	sort.Ints(semua) // bonus: urutkan
	for _, h := range semua {
		fmt.Println(h)
	}
}
