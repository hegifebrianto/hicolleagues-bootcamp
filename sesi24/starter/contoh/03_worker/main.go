// Slide "Buffered Channel & Worker Sederhana".
package main

import "fmt"

func main() {
	jobs := make(chan int, 3) // buffered: boleh menampung 3 nilai tanpa penerima
	hasil := make(chan int, 3)

	go func() { // worker: baca jobs sampai ditutup
		for j := range jobs {
			hasil <- j * 2
		}
		close(hasil) // worker selesai -> tutup hasil
	}()

	jobs <- 1
	jobs <- 2
	close(jobs) // tidak ada job lagi -> range di worker berhenti

	for h := range hasil { // berhenti otomatis saat hasil ditutup
		fmt.Println(h) // 2, 4
	}
}
