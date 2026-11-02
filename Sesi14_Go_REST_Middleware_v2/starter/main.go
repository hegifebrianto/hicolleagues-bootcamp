package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	// TODO 0a: tambahkan "strconv" di sini begitu kamu mengerjakan
	// TODO 2/8/11 (butuh strconv.Atoi untuk mengubah idMentah jadi angka).
	// TODO 0b: tambahkan "time" di sini begitu kamu mengerjakan
	// TODO 13/14 (butuh time.Now() dan time.Since() di middleware logging).
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
	// TODO 1: mu.Lock(), defer mu.Unlock(), lalu
	// tulisJSON(w, http.StatusOK, Respons{Data: daftarTugas})
}

func detailTugas(w http.ResponseWriter, r *http.Request, idMentah string) {
	// TODO 2: id, err := strconv.Atoi(idMentah). Error -> tulisError(400, "id tidak valid").
	// TODO 3: mu.Lock()/defer mu.Unlock(), loop "for _, t := range daftarTugas",
	// kalau t.ID == id -> tulisJSON(200, Respons{Data: t}) lalu return.
	// TODO 4: kalau tidak ketemu -> tulisError(404, "tugas tidak ditemukan")
}

func buatTugas(w http.ResponseWriter, r *http.Request) {
	// TODO 5: decode body ke "var t Tugas". Error/judul kosong -> tulisError(400, ...).
	// TODO 6: mu.Lock(), t.ID = idBerikutnya, idBerikutnya++, append, mu.Unlock().
	// TODO 7: tulisJSON(w, http.StatusCreated, Respons{Data: t})
}

func ubahTugas(w http.ResponseWriter, r *http.Request, idMentah string) {
	// TODO 8: id, err := strconv.Atoi(idMentah) (pola sama seperti TODO 2).
	// TODO 9: decode body ke "var masukan Tugas" (pola sama seperti buatTugas).
	// TODO 10: mu.Lock()/defer mu.Unlock(), cari di daftarTugas yang ID-nya cocok,
	// update Judul (kalau masukan.Judul != "") dan Selesai, tulisJSON(200, ...).
	// Tidak ketemu -> tulisError(404, "tugas tidak ditemukan").
}

func hapusTugas(w http.ResponseWriter, r *http.Request, idMentah string) {
	// TODO 11: id, err := strconv.Atoi(idMentah) (pola sama seperti TODO 2).
	// TODO 12: mu.Lock()/defer mu.Unlock(), cari lalu
	// daftarTugas = append(daftarTugas[:i], daftarTugas[i+1:]...),
	// w.WriteHeader(http.StatusNoContent). Tidak ketemu -> tulisError(404, ...).
}

// =====================================================================
// MIDDLEWARE
// =====================================================================

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO 13: mulai := time.Now()
		next.ServeHTTP(w, r) // WAJIB dipanggil -- jangan dihapus
		// TODO 14: log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(mulai))
	})
}

func apiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO 15 (bonus): kalau r.URL.Path == "/ping" (atau "/"),
		// next.ServeHTTP(w, r) lalu return -- lewati pengecekan di bawah.

		// TODO 16: bandingkan r.Header.Get("X-API-Key") dengan
		// os.Getenv("API_KEY"). Kalau TIDAK cocok,
		// tulisError(w, http.StatusUnauthorized, "tidak diizinkan") lalu return.

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
