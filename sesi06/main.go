package main

import (
	"bufio"
	"fmt"
	"os"
	"sesi06/belanja"
	"strconv"
	"strings"
)

func baca(in *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	teks, _ := in.ReadString('\n')
	return strings.TrimSpace(teks)
}

func main() {
	// fmt.Println("sesi06")

	// fmt.Println(belanja.DoPrint())
	// fmt.Println("Program CLI")
	// fmt.Println("masukkan menu :")
	// fmt.Println(input.ReadString('\n'))
	// pilihMenu, _ := input.ReadString('\n')
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Pilih Menu : (1)Tambah (2)Update (3)Lihat (4)Hapus (5)Total (0)Keluar")

		pilihan := baca(input, "Pilih menu: ")
		switch pilihan {
		case "1":
			menuTambah(input)
		case "2":
			menuEdit(input)
		case "3":
			menuShow(input)

		case "4":

		case "0":
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}

}

func menuTambah(in *bufio.Reader) {
	nama := baca(in, "Nama item: ")

	hargaStr := baca(in, "Harga: ")
	harga, err := strconv.ParseFloat(hargaStr, 64)
	if err != nil {
		fmt.Println("Harga harus berupa angka.")
		return
	}

	qtyStr := baca(in, "Qty: ")
	qty, err := strconv.Atoi(qtyStr)
	if err != nil {
		fmt.Println("Qty harus berupa angka bulat.")
		return
	}

	item, err := belanja.Tambah(nama, harga, qty)
	if err != nil {
		fmt.Println("Gagal:", err)
		return
	}
	fmt.Println("Ditambahkan  :", item)
	// fmt.Printf("Ditambahkan: [%d] %s x%d @Rp%.0f\n", it.ID, it.Nama, it.Qty, it.Harga)
}

func menuShow(in *bufio.Reader) {
	fmt.Print(in)
}

func menuEdit(in *bufio.Reader) {
	idStr := baca(in, "ID item: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID harus berupa angka.")
		return
	}

	qtyStr := baca(in, "Qty baru: ")
	qty, err := strconv.Atoi(qtyStr)
	if err != nil {
		fmt.Println("Qty harus berupa angka bulat.")
		return
	}

	if err := belanja.Update(id, qty); err != nil {
		fmt.Println("Gagal:", err)
		return
	}
	fmt.Println("Qty berhasil diubah.")
}
