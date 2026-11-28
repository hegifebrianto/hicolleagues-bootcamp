package auth

import (
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

// Iterasi PBKDF2 untuk demo/pelatihan. Di produksi pakai bcrypt/argon2 atau iterasi >= 600.000.
const iterasi = 100_000

// HashPassword menghasilkan "pbkdf2$iterasi$salt$hash". Salt acak -> hash password yang
// sama menghasilkan teks berbeda. Password asli TIDAK pernah disimpan.
func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	kunci, err := pbkdf2.Key(sha256.New, password, salt, iterasi, 32)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("pbkdf2$%d$%s$%s", iterasi,
		base64.RawStdEncoding.EncodeToString(salt), base64.RawStdEncoding.EncodeToString(kunci)), nil
}

// CekPassword membandingkan password dengan hash tersimpan.
func CekPassword(hash, password string) bool {
	bagian := strings.Split(hash, "$")
	if len(bagian) != 4 || bagian[0] != "pbkdf2" {
		return false
	}
	iter, err := strconv.Atoi(bagian[1])
	if err != nil || iter <= 0 || iter > 5_000_000 {
		return false
	}
	salt, err1 := base64.RawStdEncoding.DecodeString(bagian[2])
	want, err2 := base64.RawStdEncoding.DecodeString(bagian[3])
	if err1 != nil || err2 != nil || len(want) == 0 {
		return false
	}
	got, err := pbkdf2.Key(sha256.New, password, salt, iter, len(want))
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare(got, want) == 1
}
