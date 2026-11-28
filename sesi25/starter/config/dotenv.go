package config

import (
	"bufio"
	"os"
	"strings"
)

// MuatDotEnv membaca file .env (KEY=VALUE) ke environment variable.
// Variabel yang sudah diset di terminal TIDAK ditimpa. File tidak ada -> diam saja.
func MuatDotEnv(path string) {
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
		kunci, nilai = strings.TrimSpace(kunci), strings.TrimSpace(nilai)
		if len(nilai) >= 2 && (nilai[0] == '"' || nilai[0] == '\'') && nilai[len(nilai)-1] == nilai[0] {
			nilai = nilai[1 : len(nilai)-1]
		}
		if _, ada := os.LookupEnv(kunci); !ada {
			os.Setenv(kunci, nilai)
		}
	}
}
