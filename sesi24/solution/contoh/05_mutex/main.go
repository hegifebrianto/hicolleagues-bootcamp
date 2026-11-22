// Slide "sync.Mutex: Mencegah Race Condition".  Selalu mencetak 1000.
// Cek:  go run -race ./contoh/05_mutex   (bersih, tanpa WARNING)
package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex
	total int
)

func tambah() {
	mu.Lock()
	defer mu.Unlock()
	total++
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); tambah() }()
	}
	wg.Wait()
	fmt.Println(total) // 10000
}
