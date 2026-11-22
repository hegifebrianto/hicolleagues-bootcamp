// LATIHAN 1: 5 goroutine mencetak nomornya, tunggu semua dengan WaitGroup.
// Hasil: 5 baris tercetak (urutan bisa beda). Bonus: ukur waktu dengan time.Since.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mulai := time.Now()

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("Goroutine", n)
		}(i)
	}
	wg.Wait()

	fmt.Println("Total waktu:", time.Since(mulai)) // bonus
}
