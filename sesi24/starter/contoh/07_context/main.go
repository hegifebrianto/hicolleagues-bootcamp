// Slide "context untuk Timeout & Cancellation".
//
//	go run ./contoh/07_context        -> pekerjaan 1s, timeout 2s  => "Selesai tepat waktu"
//	go run ./contoh/07_context 3s     -> pekerjaan 3s, timeout 2s  => "Timeout: context deadline exceeded"
package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

// waktuKerja mensimulasikan pekerjaan yang butuh `lama` waktu.
// Channel dibuat berkapasitas 1 supaya goroutine tidak macet selamanya
// bila tidak ada yang membaca hasilnya (mencegah goroutine leak).
func waktuKerja(ctx context.Context, lama time.Duration) <-chan string {
	out := make(chan string, 1)
	go func() {
		select {
		case <-time.After(lama):
			out <- "hasil pekerjaan"
		case <-ctx.Done(): // dibatalkan -> berhenti lebih awal
		}
	}()
	return out
}

func main() {
	lama := 1 * time.Second
	if len(os.Args) > 1 {
		if d, err := time.ParseDuration(os.Args[1]); err == nil {
			lama = d
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // selalu panggil cancel untuk melepas resource

	select {
	case <-waktuKerja(ctx, lama):
		fmt.Println("Selesai tepat waktu")
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}
}
