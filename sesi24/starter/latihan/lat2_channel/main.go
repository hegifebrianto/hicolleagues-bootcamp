// LATIHAN 2 (9 menit): Channel Hasil
// Hitung kuadrat 1..5 di goroutine, kirim hasilnya lewat channel, cetak di main.
//
// Hasil: 1 4 9 16 25 (urutan tergantung).
// Petunjuk: buffered channel kapasitas 5; close(hasil) SETELAH semua pengirim selesai.
//
//	Jebakan: close di tempat yang salah -> panic "send on closed channel";
//	         tidak pernah close -> `for h := range hasil` menunggu selamanya (deadlock).
//	Pola aman: go func() { wg.Wait(); close(hasil) }()
//
// Bonus: urutkan hasil sebelum dicetak (paket sort).
//
// Jalankan:  go run ./latihan/lat2_channel
package main

import "fmt"

func main() {
	// TODO: hasil := make(chan int, 5)
	// TODO: 5 goroutine (pakai WaitGroup) masing-masing mengirim n*n ke hasil
	// TODO: goroutine penutup: wg.Wait() lalu close(hasil)
	// TODO: for h := range hasil { fmt.Println(h) }

	fmt.Println("TODO: belum dikerjakan")
}
