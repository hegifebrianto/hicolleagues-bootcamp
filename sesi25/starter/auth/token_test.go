package auth

import (
	"encoding/base64"
	"strings"
	"testing"
	"time"
)

const rahasia = "rahasia-untuk-test-123"

func TestBuatLaluCek(t *testing.T) {
	p := NewPenerbit(rahasia, time.Minute)
	token, err := p.Buat(7, "admin")
	if err != nil {
		t.Fatal(err)
	}
	c, err := p.Cek(token)
	if err != nil {
		t.Fatalf("token baru harus valid: %v", err)
	}
	if c.Sub != "7" || c.Role != "admin" {
		t.Errorf("claims salah: %+v", c)
	}
}

func TestCekGagal(t *testing.T) {
	p := NewPenerbit(rahasia, time.Minute)
	token, _ := p.Buat(1, "user")
	bagian := strings.Split(token, ".")
	b64 := base64.RawURLEncoding

	// payload diubah jadi admin, tanda tangan lama dipertahankan
	payloadAdmin := b64.EncodeToString([]byte(`{"sub":"1","role":"admin","iat":0,"exp":99999999999}`))
	dipalsukan := bagian[0] + "." + payloadAdmin + "." + bagian[2]

	// alg none tanpa tanda tangan
	algNone := b64.EncodeToString([]byte(`{"alg":"none","typ":"JWT"}`)) + "." + bagian[1] + "."

	lain := NewPenerbit("kunci-lain-yang-berbeda", time.Minute)
	tokenLain, _ := lain.Buat(1, "user")

	tests := []struct {
		nama  string
		token string
	}{
		{"kosong", ""},
		{"bukan jwt", "abc"},
		{"payload dipalsukan", dipalsukan},
		{"alg none", algNone},
		{"kunci berbeda", tokenLain},
		{"tanda tangan rusak", bagian[0] + "." + bagian[1] + ".@@@"},
	}
	for _, tt := range tests {
		t.Run(tt.nama, func(t *testing.T) {
			if _, err := p.Cek(tt.token); err != ErrTokenTidakValid {
				t.Errorf("harusnya ErrTokenTidakValid, dapat %v", err)
			}
		})
	}
}

func TestTokenKedaluwarsa(t *testing.T) {
	sekarang := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	p := NewPenerbit(rahasia, 15*time.Minute)
	p.Sekarang = func() time.Time { return sekarang }
	token, _ := p.Buat(1, "user")

	sekarang = sekarang.Add(14 * time.Minute)
	if _, err := p.Cek(token); err != nil {
		t.Errorf("menit ke-14 masih valid: %v", err)
	}
	sekarang = sekarang.Add(2 * time.Minute)
	if _, err := p.Cek(token); err != ErrTokenKedaluwarsa {
		t.Errorf("menit ke-16 harus kedaluwarsa, dapat %v", err)
	}
}
