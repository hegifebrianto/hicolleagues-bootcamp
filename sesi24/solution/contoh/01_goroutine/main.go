// Slide "Goroutine: Ringan dan Mudah".
// Jalankan beberapa kali: urutan "Goroutine 1/2" bisa berubah-ubah.
package main

import (
	"fmt"
	"time"
)

func cetak(pesan string) {
	fmt.Println(pesan)
}

func main() {
	go cetak("Goroutine 1")
	go cetak("Goroutine 2")
	// time.Sleep hanya sementara agar main tidak keluar duluan (BUKAN solusi ideal).
	// Coba hapus baris ini: kemungkinan besar tidak ada yang tercetak!
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Selesai")
}
