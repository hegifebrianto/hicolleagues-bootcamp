// Slide "sync.WaitGroup".
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // SEBELUM `go`, bukan di dalam goroutine
		go func(n int) {
			defer wg.Done()
			fmt.Println("Tugas", n)
		}(i) // kirim i sebagai parameter n
	}
	wg.Wait() // tunggu sampai semua Done()
	fmt.Println("Semua selesai")
}
