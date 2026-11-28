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

// Buat menerbitkan token untuk user tertentu.
func (p *Penerbit) Buat(userID int64, role string) (string, error) {
	now := p.Sekarang()
	payload, err := json.Marshal(Claims{
		Sub:  strconv.FormatInt(userID, 10),
		Role: role,
		Iat:  now.Unix(),
		Exp:  now.Add(p.TTL).Unix(),
	})
	if err != nil {
		return "", err
	}
	header := b64.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	data := header + "." + b64.EncodeToString(payload)
	return data + "." + b64.EncodeToString(p.tandatangan(data)), nil
}

// Cek memverifikasi tanda tangan, algoritma, dan masa berlaku.
func (p *Penerbit) Cek(token string) (*Claims, error) {
	bagian := strings.Split(token, ".")
	if len(bagian) != 3 {
		return nil, ErrTokenTidakValid
	}

	// 1. Algoritma wajib HS256 (menolak "alg":"none" dan algoritma tak terduga)
	hb, err := b64.DecodeString(bagian[0])
	if err != nil {
		return nil, ErrTokenTidakValid
	}
	var h struct {
		Alg string `json:"alg"`
	}
	if json.Unmarshal(hb, &h) != nil || h.Alg != "HS256" {
		return nil, ErrTokenTidakValid
	}

	// 2. Tanda tangan harus cocok (hmac.Equal = perbandingan waktu-konstan)
	sig, err := b64.DecodeString(bagian[2])
	if err != nil || !hmac.Equal(sig, p.tandatangan(bagian[0]+"."+bagian[1])) {
		return nil, ErrTokenTidakValid
	}

	// 3. Isi token
	pb, err := b64.DecodeString(bagian[1])
	if err != nil {
		return nil, ErrTokenTidakValid
	}
	var c Claims
	if json.Unmarshal(pb, &c) != nil {
		return nil, ErrTokenTidakValid
	}

	// 4. Masa berlaku
	if p.Sekarang().Unix() >= c.Exp {
		return nil, ErrTokenKedaluwarsa
	}
	return &c, nil
}
