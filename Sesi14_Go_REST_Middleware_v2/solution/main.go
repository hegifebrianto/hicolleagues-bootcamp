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

// =====================================================================
// MODEL
// =====================================================================

type Tugas struct {
	ID      int    `json:"id"`
	Judul   string `json:"judul"`
	Selesai bool   `json:"selesai"`
}

type Respons struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func tulisJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func tulisError(w http.ResponseWriter, status int, pesan string) {
	tulisJSON(w, status, Respons{Error: pesan})
}

// =====================================================================
// PENYIMPANAN DI MEMORI
// =====================================================================

var (
	mu          sync.Mutex
	daftarTugas = []Tugas{
		{ID: 1, Judul: "Belajar Go", Selesai: false},
		{ID: 2, Judul: "Belajar REST", Selesai: true},
	}
	idBerikutnya = 3
)

// =====================================================================
// HANDLER /api/tugas
// =====================================================================

func daftarTugasHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	tulisJSON(w, http.StatusOK, Respons{Data: daftarTugas})
}

func detailTugas(w http.ResponseWriter, r *http.Request, idMentah string) {
	id, err := strconv.Atoi(idMentah)
	if err != nil {
		tulisError(w, http.StatusBadRequest, "id tidak valid")
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for _, t := range daftarTugas {
		if t.ID == id {
			tulisJSON(w, http.StatusOK, Respons{Data: t})
			return
		}
	}
	tulisError(w, http.StatusNotFound, "tugas tidak ditemukan")
}

func buatTugas(w http.ResponseWriter, r *http.Request) {
	var t Tugas
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		tulisError(w, http.StatusBadRequest, "body tidak valid")
		return
	}
	if t.Judul == "" {
		tulisError(w, http.StatusBadRequest, "judul wajib diisi")
		return
	}
	mu.Lock()
	t.ID = idBerikutnya
	idBerikutnya++
	daftarTugas = append(daftarTugas, t)
	mu.Unlock()
	tulisJSON(w, http.StatusCreated, Respons{Data: t})
}

func ubahTugas(w http.ResponseWriter, r *http.Request, idMentah string) {
	id, err := strconv.Atoi(idMentah)
	if err != nil {
		tulisError(w, http.StatusBadRequest, "id tidak valid")
		return
	}
	var masukan Tugas
	if err := json.NewDecoder(r.Body).Decode(&masukan); err != nil {
		tulisError(w, http.StatusBadRequest, "body tidak valid")
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for i := range daftarTugas {
		if daftarTugas[i].ID == id {
			if masukan.Judul != "" {
				daftarTugas[i].Judul = masukan.Judul
			}
			daftarTugas[i].Selesai = masukan.Selesai
			tulisJSON(w, http.StatusOK, Respons{Data: daftarTugas[i]})
			return
		}
	}
	tulisError(w, http.StatusNotFound, "tugas tidak ditemukan")
}

func hapusTugas(w http.ResponseWriter, r *http.Request, idMentah string) {
	id, err := strconv.Atoi(idMentah)
	if err != nil {
		tulisError(w, http.StatusBadRequest, "id tidak valid")
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for i, t := range daftarTugas {
		if t.ID == id {
			daftarTugas = append(daftarTugas[:i], daftarTugas[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	tulisError(w, http.StatusNotFound, "tugas tidak ditemukan")
}

// =====================================================================
// MIDDLEWARE
// =====================================================================

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mulai := time.Now()
		next.ServeHTTP(w, r)

		log.Printf("%s %s %v ", r.Method, r.URL.Path, mulai)

		// fmt.Println(w.status)
		// log.Printf("%s ")

	})
}

func apiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/ping" || r.URL.Path == "/" {
			next.ServeHTTP(w, r)
			return
		}
		if r.Header.Get("X-API-Key") != os.Getenv("API_KEY") {
			tulisError(w, http.StatusUnauthorized, "tidak diizinkan")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// =====================================================================
// MAIN -- semua route didaftarkan dan dicetak di sini supaya mudah dicek
// =====================================================================

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089" // sengaja BUKAN 8080 -- port itu sering dipakai
		// proses lain (Docker, dev server lain, dll) sehingga mudah
		// bentrok tanpa disadari. Set env var PORT untuk ganti.
	}

	mux := http.NewServeMux()

	// "/" -- halaman cek cepat: buka http://localhost:PORT/ di browser,
	// kalau ini muncul berarti server BENAR sedang dijalankan dan
	// didengar oleh browser/curl-mu, bukan proses lain di port yang sama.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			tulisError(w, http.StatusNotFound, "halaman tidak ditemukan: "+r.URL.Path)
			return
		}
		fmt.Fprintln(w, "Server Sesi 14 aktif. Coba: /ping, /api/tugas")
	})

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		tulisJSON(w, http.StatusOK, Respons{Data: "pong"})
	})

	mux.HandleFunc("/api/tugas", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			daftarTugasHandler(w, r)
		case http.MethodPost:
			buatTugas(w, r)
		default:
			tulisError(w, http.StatusMethodNotAllowed, "metode tidak didukung")
		}
	})

	mux.HandleFunc("/api/tugas/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/api/tugas/")
		fmt.Println(id, "ID")
		if id == "" {
			tulisError(w, http.StatusBadRequest, "id wajib diisi di URL")
			return
		}
		switch r.Method {
		case http.MethodGet:
			detailTugas(w, r, id)
		case http.MethodPut:
			ubahTugas(w, r, id)
		case http.MethodDelete:
			hapusTugas(w, r, id)
		default:
			tulisError(w, http.StatusMethodNotAllowed, "metode tidak didukung")
		}
	})

	handler := logging(apiKey(mux))

	fmt.Println("==========================================")
	fmt.Println(" Sesi 14 -- REST API Tugas + Middleware")
	fmt.Println("==========================================")
	fmt.Println(" Alamat  : http://localhost:" + port)
	fmt.Println(" Coba    : http://localhost:" + port + "/ping")
	fmt.Println(" Route   : GET/POST /api/tugas, GET/PUT/DELETE /api/tugas/{id}")
	fmt.Println(" API_KEY : " + os.Getenv("API_KEY") + "  (kosong = semua /api/tugas ditolak 401)")
	fmt.Println("==========================================")

	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal("GAGAL MENYALAKAN SERVER: ", err)
	}
}
