// Slide "Channel: Komunikasi Aman".
package main

import "fmt"

func main() {
	ch := make(chan string) // unbuffered: pengirim & penerima harus bertemu
	go func() {
		ch <- "Halo dari goroutine" // kirim (menunggu sampai ada penerima)
	}()
	pesan := <-ch // terima (menunggu sampai ada kiriman)
	fmt.Println(pesan)
}
