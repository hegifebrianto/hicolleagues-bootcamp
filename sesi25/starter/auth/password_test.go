package auth

import (
	"strings"
	"testing"
)

func TestPasswordHash(t *testing.T) {
	h, err := HashPassword("rahasia")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(h, "rahasia") {
		t.Error("hash tidak boleh memuat password asli")
	}
	if !CekPassword(h, "rahasia") {
		t.Error("password benar harus lolos")
	}
	if CekPassword(h, "salah") {
		t.Error("password salah harus ditolak")
	}
	h2, _ := HashPassword("rahasia")
	if h == h2 {
		t.Error("salt acak: hash dua kali harus berbeda")
	}
	if CekPassword("bukan-hash", "rahasia") {
		t.Error("format hash rusak harus ditolak")
	}
}
