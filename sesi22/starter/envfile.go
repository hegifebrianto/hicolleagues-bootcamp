package main

import (
	"bufio"
	"os"
	"strings"
)

// muatDotEnv membaca file .env (baris KEY=VALUE) ke environment variable.
// - baris kosong dan yang diawali # diabaikan
// - nilai boleh diberi tanda kutip: DB_PASSWORD="rahasia"
// - variabel yang SUDAH diset di terminal tidak ditimpa (terminal menang)
// - file .env tidak ada -> diam saja (pakai nilai bawaan)
// Ditulis sendiri agar tidak perlu library tambahan (mis. godotenv).
func muatDotEnv(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		baris := strings.TrimSpace(sc.Text())
		if baris == "" || strings.HasPrefix(baris, "#") {
			continue
		}
		baris = strings.TrimPrefix(baris, "export ")
		kunci, nilai, ok := strings.Cut(baris, "=")
		if !ok {
			continue
		}
		kunci = strings.TrimSpace(kunci)
		nilai = strings.TrimSpace(nilai)
		if len(nilai) >= 2 && (nilai[0] == 0x22 || nilai[0] == 0x27) && nilai[len(nilai)-1] == nilai[0] {
			nilai = nilai[1 : len(nilai)-1]
		}
		if _, sudahAda := os.LookupEnv(kunci); !sudahAda {
			os.Setenv(kunci, nilai)
		}
	}
}
