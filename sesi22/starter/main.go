// LATIHAN 3 (8 menit): rakit semua lapisan di main.
// Saat ini hanya /api/ping yang hidup. Selesaikan TODO di bawah.
package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := muatKonfigurasi()

	// TODO 1: kalau cfg.Store == "memory" -> repo := NewBukuRepoMemori()
	//         selain itu: sql.Open("mysql", cfg.DSN()) + db.Ping() -> repo := NewBukuRepoSQL(db)
	//         (tip: coba STORE=memory dulu agar Latihan 2-3 bisa diuji tanpa MySQL)
	// TODO 2: service := NewBukuService(repo)
	// TODO 3: handler := NewBukuHandler(service)
	// Urutan wajib: DB -> Repo -> Service -> Handler

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/ping", func(w http.ResponseWriter, r *http.Request) {
		tulisJSON(w, http.StatusOK, Respons{Data: "pong"})
	})
	// TODO 4: daftarkan route:
	//   GET    /api/buku        -> handler.Semua
	//   GET    /api/buku/{id}   -> handler.Ambil
	//   POST   /api/buku        -> handler.Tambah
	//   DELETE /api/buku/{id}   -> handler.Hapus
	// Bonus: pindahkan perakitan ke fungsi initApp(cfg Config) http.Handler

	log.Println("Server berjalan di http://localhost:" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
