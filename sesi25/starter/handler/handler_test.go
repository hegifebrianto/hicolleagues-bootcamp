package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"project3/auth"
	"project3/middleware"
	"project3/model"
	"project3/repository"
	"project3/service"
)

type mailerDiam struct{}

func (mailerDiam) KirimSelamatDatang(string) {}

// siapkan merakit aplikasi UTUH (service asli + repo memori) dengan 2 akun seed.
func siapkan(t *testing.T) http.Handler {
	t.Helper()
	tokens := auth.NewPenerbit("rahasia-test-1234567", time.Minute)
	users := repository.NewUserMemori()
	for _, u := range []struct{ email, role string }{{"admin@mail.com", "admin"}, {"budi@mail.com", "user"}} {
		h, _ := auth.HashPassword("rahasia")
		users.Simpan(model.User{Nama: u.role, Email: u.email, PasswordHash: h, Role: u.role})
	}
	buku := repository.NewBukuMemori(model.Buku{Judul: "Laskar Pelangi", Tahun: 2005})
	return NewRouter(
		&BukuHandler{Svc: service.NewBukuService(buku)},
		&AuthHandler{Svc: service.NewAuthService(users, tokens, mailerDiam{})},
		&middleware.Middleware{Tokens: tokens},
	)
}

func kirim(h http.Handler, metode, url, body, token string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(metode, url, strings.NewReader(body))
	if token != "" {
		r.Header.Set("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w
}

func login(t *testing.T, h http.Handler, email, pass string) string {
	t.Helper()
	w := kirim(h, "POST", "/api/login", `{"email":"`+email+`","password":"`+pass+`"}`, "")
	if w.Code != 200 {
		t.Fatalf("login %s: %d %s", email, w.Code, w.Body)
	}
	var res struct{ Data struct{ Token string } }
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Data.Token == "" {
		t.Fatal("token kosong")
	}
	return res.Data.Token
}

func TestLogin(t *testing.T) {
	h := siapkan(t)
	login(t, h, "admin@mail.com", "rahasia")
	if w := kirim(h, "POST", "/api/login", `{"email":"admin@mail.com","password":"salah"}`, ""); w.Code != 401 {
		t.Errorf("password salah: %d, mau 401", w.Code)
	}
	if w := kirim(h, "POST", "/api/login", `bukan json`, ""); w.Code != 400 {
		t.Errorf("json rusak: %d, mau 400", w.Code)
	}
}

// Skenario slide "Uji dengan curl": tanpa token 401, user 403, admin 204.
func TestHapusButuhAdmin(t *testing.T) {
	h := siapkan(t)
	if w := kirim(h, "DELETE", "/api/buku/1", "", ""); w.Code != 401 {
		t.Errorf("tanpa token: %d, mau 401", w.Code)
	}
	user := login(t, h, "budi@mail.com", "rahasia")
	if w := kirim(h, "DELETE", "/api/buku/1", "", user); w.Code != 403 {
		t.Errorf("token user: %d, mau 403", w.Code)
	}
	admin := login(t, h, "admin@mail.com", "rahasia")
	if w := kirim(h, "DELETE", "/api/buku/1", "", admin); w.Code != 204 {
		t.Errorf("token admin: %d, mau 204", w.Code)
	}
	if w := kirim(h, "DELETE", "/api/buku/1", "", admin); w.Code != 404 {
		t.Errorf("hapus ulang: %d, mau 404", w.Code)
	}
}

func TestTambahButuhLogin(t *testing.T) {
	h := siapkan(t)
	body := `{"judul":"Go Dasar","penulis":"A","tahun":2020}`
	if w := kirim(h, "POST", "/api/buku", body, ""); w.Code != 401 {
		t.Errorf("tanpa token: %d, mau 401", w.Code)
	}
	user := login(t, h, "budi@mail.com", "rahasia")
	if w := kirim(h, "POST", "/api/buku", body, user); w.Code != 201 {
		t.Errorf("user login boleh tambah: %d, mau 201", w.Code)
	}
	if w := kirim(h, "POST", "/api/buku", `{"judul":"","tahun":2020}`, user); w.Code != 400 {
		t.Errorf("validasi: %d, mau 400", w.Code)
	}
}

func TestPublikTanpaToken(t *testing.T) {
	h := siapkan(t)
	for _, url := range []string{"/api/ping", "/api/buku", "/api/buku/1"} {
		if w := kirim(h, "GET", url, "", ""); w.Code != 200 {
			t.Errorf("GET %s: %d, mau 200", url, w.Code)
		}
	}
	if w := kirim(h, "GET", "/api/buku/abc", "", ""); w.Code != 400 {
		t.Errorf("id bukan angka: %d, mau 400", w.Code)
	}
	if w := kirim(h, "GET", "/api/buku/99", "", ""); w.Code != 404 {
		t.Errorf("id tak ada: %d, mau 404", w.Code)
	}
}

func TestRegisterDanMe(t *testing.T) {
	h := siapkan(t)
	w := kirim(h, "POST", "/api/register", `{"nama":"Ani","email":"ani@mail.com","password":"rahasia1"}`, "")
	if w.Code != 201 {
		t.Fatalf("register: %d %s", w.Code, w.Body)
	}
	if strings.Contains(w.Body.String(), "pbkdf2") || strings.Contains(strings.ToLower(w.Body.String()), "password") {
		t.Errorf("response register tidak boleh membocorkan hash: %s", w.Body)
	}
	if w := kirim(h, "POST", "/api/register", `{"nama":"Ani","email":"ani@mail.com","password":"rahasia1"}`, ""); w.Code != 409 {
		t.Errorf("email dobel: %d, mau 409", w.Code)
	}
	token := login(t, h, "ani@mail.com", "rahasia1")
	w = kirim(h, "GET", "/api/me", "", token)
	if w.Code != 200 || !strings.Contains(w.Body.String(), `"role":"user"`) {
		t.Errorf("/api/me: %d %s", w.Code, w.Body)
	}
	if w := kirim(h, "GET", "/api/me", "", ""); w.Code != 401 {
		t.Errorf("/api/me tanpa token: %d, mau 401", w.Code)
	}
}
