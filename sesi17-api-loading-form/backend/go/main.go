// Backend demo Sesi 17 & 18 -- hanya standard library, data di MEMORI
// (hilang saat server di-restart). Dua resource: /api/tugas dan /api/buku.
//
// Jalankan:  go run .
// Opsi (environment variable):
//   PORT=8089        port server (default 8089)
//   DELAY_MS=700     jeda buatan untuk request GET supaya "Memuat..." terlihat
//   CORS_ORIGIN=...  origin frontend yang diizinkan (default http://localhost:5173)
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type item map[string]any

type koleksi struct {
	mu     sync.Mutex
	base   string // contoh: "/api/tugas"
	wajib  string // field yang wajib diisi (tidak boleh kosong)
	bawaan item   // nilai default untuk field yang tidak dikirim saat POST
	data   []item
	nextID int
}

func tulisJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func tulisError(w http.ResponseWriter, status int, pesan string) {
	tulisJSON(w, status, map[string]any{"error": pesan})
}

func (k *koleksi) cari(id int) int {
	for i, it := range k.data {
		if it["id"] == id {
			return i
		}
	}
	return -1
}

func (k *koleksi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sisa := strings.Trim(strings.TrimPrefix(r.URL.Path, k.base), "/")

	// /api/tugas  (tanpa id)
	if sisa == "" {
		switch r.Method {
		case http.MethodGet:
			k.daftar(w)
		case http.MethodPost:
			k.buat(w, r)
		default:
			tulisError(w, http.StatusMethodNotAllowed, "metode tidak didukung")
		}
		return
	}

	// /api/tugas/{id}
	id, err := strconv.Atoi(sisa)
	if err != nil {
		tulisError(w, http.StatusBadRequest, "id tidak valid")
		return
	}
	switch r.Method {
	case http.MethodGet:
		k.detail(w, id)
	case http.MethodPut:
		k.ubah(w, r, id)
	case http.MethodDelete:
		k.hapus(w, id)
	default:
		tulisError(w, http.StatusMethodNotAllowed, "metode tidak didukung")
	}
}

func (k *koleksi) daftar(w http.ResponseWriter) {
	k.mu.Lock()
	defer k.mu.Unlock()
	hasil := make([]item, len(k.data))
	copy(hasil, k.data)
	tulisJSON(w, http.StatusOK, map[string]any{"data": hasil})
}

func (k *koleksi) detail(w http.ResponseWriter, id int) {
	k.mu.Lock()
	defer k.mu.Unlock()
	i := k.cari(id)
	if i < 0 {
		tulisError(w, http.StatusNotFound, "data tidak ditemukan")
		return
	}
	tulisJSON(w, http.StatusOK, map[string]any{"data": k.data[i]})
}

func (k *koleksi) buat(w http.ResponseWriter, r *http.Request) {
	var masuk item
	if err := json.NewDecoder(r.Body).Decode(&masuk); err != nil || masuk == nil {
		tulisError(w, http.StatusBadRequest, "body tidak valid")
		return
	}
	if s, _ := masuk[k.wajib].(string); strings.TrimSpace(s) == "" {
		tulisError(w, http.StatusBadRequest, k.wajib+" wajib diisi")
		return
	}

	k.mu.Lock()
	defer k.mu.Unlock()
	for kunci, nilai := range k.bawaan {
		if _, ada := masuk[kunci]; !ada {
			masuk[kunci] = nilai
		}
	}
	masuk["id"] = k.nextID // abaikan id dari klien
	k.nextID++
	k.data = append(k.data, masuk)
	tulisJSON(w, http.StatusCreated, map[string]any{"data": masuk})
}

func (k *koleksi) ubah(w http.ResponseWriter, r *http.Request, id int) {
	var masuk item
	if err := json.NewDecoder(r.Body).Decode(&masuk); err != nil || masuk == nil {
		tulisError(w, http.StatusBadRequest, "body tidak valid")
		return
	}

	k.mu.Lock()
	defer k.mu.Unlock()
	i := k.cari(id)
	if i < 0 {
		tulisError(w, http.StatusNotFound, "data tidak ditemukan")
		return
	}
	if v, ada := masuk[k.wajib]; ada {
		if s, _ := v.(string); strings.TrimSpace(s) == "" {
			tulisError(w, http.StatusBadRequest, k.wajib+" wajib diisi")
			return
		}
	}
	for kunci, nilai := range masuk {
		if kunci == "id" {
			continue
		}
		k.data[i][kunci] = nilai
	}
	tulisJSON(w, http.StatusOK, map[string]any{"data": k.data[i]})
}

func (k *koleksi) hapus(w http.ResponseWriter, id int) {
	k.mu.Lock()
	defer k.mu.Unlock()
	i := k.cari(id)
	if i < 0 {
		tulisError(w, http.StatusNotFound, "data tidak ditemukan")
		return
	}
	k.data = append(k.data[:i], k.data[i+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

// ---------------------------------------------------------------- middleware

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mulai := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(mulai))
	})
}

func cors(next http.Handler) http.Handler {
	origin := os.Getenv("CORS_ORIGIN")
	if origin == "" {
		origin = "http://localhost:5173"
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// jeda buatan khusus GET (selain /api/ping) supaya indikator "Memuat..." sempat terlihat saat demo
func jeda(next http.Handler) http.Handler {
	ms := 700
	if v := os.Getenv("DELAY_MS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			ms = n
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ms > 0 && r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/api/") && r.URL.Path != "/api/ping" {
			time.Sleep(time.Duration(ms) * time.Millisecond)
		}
		next.ServeHTTP(w, r)
	})
}

// ---------------------------------------------------------------------- main

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}

	tugas := &koleksi{
		base:   "/api/tugas",
		wajib:  "judul",
		bawaan: item{"selesai": false},
		nextID: 4,
		data: []item{
			{"id": 1, "judul": "Belajar Axios", "selesai": true},
			{"id": 2, "judul": "Belajar loading state", "selesai": false},
			{"id": 3, "judul": "Sambungkan Vue ke API Go", "selesai": false},
		},
	}
	buku := &koleksi{
		base:   "/api/buku",
		wajib:  "judul",
		bawaan: item{},
		nextID: 4,
		data: []item{
			{"id": 1, "judul": "Laskar Pelangi", "penulis": "Andrea Hirata", "tahun": 2005},
			{"id": 2, "judul": "Bumi Manusia", "penulis": "Pramoedya Ananta Toer", "tahun": 1980},
			{"id": 3, "judul": "Filosofi Teras", "penulis": "Henry Manampiring", "tahun": 2019},
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			tulisError(w, http.StatusNotFound, "tidak ditemukan: "+r.URL.Path)
			return
		}
		fmt.Fprintln(w, "Backend demo aktif. Coba: /api/ping, /api/tugas, /api/buku")
	})
	mux.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		tulisJSON(w, http.StatusOK, map[string]any{"data": "pong"})
	})
	mux.Handle("/api/tugas", tugas)
	mux.Handle("/api/tugas/", tugas)
	mux.Handle("/api/buku", buku)
	mux.Handle("/api/buku/", buku)

	handler := logging(cors(jeda(mux)))

	fmt.Println("==========================================")
	fmt.Println(" Backend demo Sesi 17 & 18 (Go)")
	fmt.Println(" Alamat : http://localhost:" + port)
	fmt.Println(" Cek    : http://localhost:" + port + "/api/ping")
	fmt.Println("==========================================")

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("GAGAL MENYALAKAN SERVER: ", err)
	}
}
