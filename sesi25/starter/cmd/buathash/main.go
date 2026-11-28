// Pakai: go run ./cmd/buathash rahasia   -> mencetak hash untuk dimasukkan ke tabel users.
package main

import (
	"fmt"
	"os"

	"project3/auth"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("pakai: go run ./cmd/buathash <password>")
		os.Exit(1)
	}
	h, err := auth.HashPassword(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(h)
}
