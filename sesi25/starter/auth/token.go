// Package auth: pembuatan/verifikasi token JWT (HS256) dan hashing password.
//
// Token dibuat dengan standard library (crypto/hmac) agar project jalan tanpa
// dependensi tambahan DAN strukturnya terlihat jelas (lihat slide "Anatomi JWT"):
//
//	base64url(header) . base64url(payload) . base64url(HMAC-SHA256(header.payload, SECRET))
//
// Di aplikasi nyata umumnya dipakai library golang-jwt/jwt (lihat README untuk padanannya).
package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	ErrTokenTidakValid  = errors.New("token tidak valid")
	ErrTokenKedaluwarsa = errors.New("token kedaluwarsa")
)

// Claims = isi (payload) token. Bisa dibaca siapa saja -> JANGAN taruh data rahasia.
type Claims struct {
	Sub  string `json:"sub"` // id user
	Role string `json:"role"`
	Iat  int64  `json:"iat"` // waktu dibuat (detik Unix)
	Exp  int64  `json:"exp"` // waktu kedaluwarsa (detik Unix)
}

// Penerbit membuat dan memeriksa token dengan satu kunci rahasia.
type Penerbit struct {
	Secret   []byte
	TTL      time.Duration
	Sekarang func() time.Time // bisa diganti di test (mis. untuk menguji token kedaluwarsa)
}

func NewPenerbit(secret string, ttl time.Duration) *Penerbit {
	return &Penerbit{Secret: []byte(secret), TTL: ttl, Sekarang: time.Now}
}

var b64 = base64.RawURLEncoding

func (p *Penerbit) tandatangan(data string) []byte {
	m := hmac.New(sha256.New, p.Secret)
	m.Write([]byte(data))
	return m.Sum(nil)
}

// TODO 1a: Buat menerbitkan token HS256.
//
//	header  = base64url(`{"alg":"HS256","typ":"JWT"}`)
//	payload = base64url(json Claims{Sub, Role, Iat, Exp})   Exp = sekarang + p.TTL
//	token   = header + "." + payload + "." + base64url(p.tandatangan(header + "." + payload))
//
// Petunjuk: p.Sekarang() untuk waktu, b64.EncodeToString(...), strconv.FormatInt(id, 10).
func (p *Penerbit) Buat(userID int64, role string) (string, error) {
	return "", errors.New("TODO: belum diimplementasi")
}

// TODO 1b: Cek memverifikasi token. Urutan pemeriksaan:
//  1. dipecah "." harus 3 bagian
//  2. header: alg WAJIB "HS256" (tolak "none")
//  3. tanda tangan: hmac.Equal(sigDariToken, p.tandatangan(bagian0 + "." + bagian1))
//  4. decode payload ke Claims
//  5. kedaluwarsa bila p.Sekarang().Unix() >= c.Exp  -> ErrTokenKedaluwarsa
//
// Semua kegagalan lain -> ErrTokenTidakValid.
func (p *Penerbit) Cek(token string) (*Claims, error) {
	return nil, ErrTokenTidakValid
}

// Pakai agar import tidak "unused" selama TODO belum dikerjakan; hapus setelah selesai.
var _ = []any{hmac.Equal, base64.RawURLEncoding, json.Marshal, strconv.Itoa, strings.Split}
