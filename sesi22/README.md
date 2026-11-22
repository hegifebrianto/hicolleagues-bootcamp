# Project 2 — Arsitektur Go Berlapis · Sesi 17 (Fase 3)

Refaktor Project 2 (API Buku, Go + MySQL) dari "semua di handler" menjadi
**handler → service → repository → database** dengan dependency injection.

| Folder | Fungsi |
|---|---|
| `sebelum/` | Monolit: SQL + validasi + HTTP menumpuk di handler. Bahan slide "Kode Menumpuk vs Berlapis". |
| `starter/` | Kerangka untuk peserta: Latihan 1–3 berisi `TODO`. Sudah ada `service_test.go` sebagai pengecek Latihan 2. |
| `solution/` | Versi berlapis lengkap + `service_test.go` (untuk trainer). |

## Menjalankan
```bash
cd solution                 # atau sebelum / starter
mysql -u root -p < schema.sql
go mod tidy                 # sekali saja (unduh driver MySQL, butuh internet)
go run .                    # http://localhost:8089/api/ping
```
**Tanpa MySQL** (solution/starter): `STORE=memory go run .` — data di memori, berguna kalau MySQL bermasalah saat demo.
**Kredensial MySQL ada di file `.env`** (bukan di kode), dibaca oleh `envfile.go`:
```
DB_USER=root
DB_PASSWORD=          # isi password MySQL Anda
DB_HOST=localhost
DB_PORT=3306
DB_NAME=bukudb
PORT=8089
```
- Edit `.env` di folder `solution/` (atau `starter/`) sesuai MySQL Anda, lalu `go run .`.
- `.env.example` = contoh yang boleh di-commit; `.env` sudah ada di `.gitignore` (jangan di-commit).
- Variabel yang diset di terminal menang atas `.env`: `PORT=8090 go run .`
- `STORE=memory go run .` tetap bisa dipakai tanpa MySQL.

> Catatan: slide memakai `sql.Open("postgres", ...)`; project ini memakai **MySQL** (sesuai Sesi 13+), jadi `"mysql"` dan placeholder `?`.

## Peta slide → file (solution)
| Slide | File |
|---|---|
| Repository Layer (interface + `bukuRepoSQL`) | `repo.go`, `repo_sql.go` |
| Service Layer (`BukuService`, `ErrValidasi`) | `service.go`, `model.go` |
| Handler Layer (`tulisErrorAPI`) | `handler.go` |
| Dependency Injection (rakit di main) | `main.go` → `initApp` |
| Konfigurasi (`Config`, env var, `.env`) | `config.go`, `envfile.go` |
| Factory (`NewBukuService`, `NewBukuRepoSQL`) | `service.go`, `repo_sql.go` |
| Singleton (satu koneksi DB dibagi) | `initApp`: `sql.Open` sekali |

## Skenario demo (±10 menit)
1. Buka `sebelum/main.go`: tunjukkan SQL, validasi, dan HTTP bercampur; error DB bocor ke klien.
2. Buka `solution/`: telusuri satu request `POST /api/buku` lewat handler → service → repo.
3. **Manfaat layering:** `go test ./...` jalan **tanpa server & tanpa MySQL** (pakai `repo_memory.go`).
4. Tukar repo: `STORE=memory go run .` vs `go run .` — handler & service tidak berubah satu baris pun.

## Uji cepat (curl)
```bash
curl -i localhost:8089/api/buku
curl -i -X POST localhost:8089/api/buku -d '{"judul":"","penulis":"A","tahun":2000}'     # 400 judul wajib diisi
curl -i -X POST localhost:8089/api/buku -d '{"judul":"Kuno","penulis":"A","tahun":1800}' # 400 tahun tidak valid
curl -i -X POST localhost:8089/api/buku -d '{"judul":"Buku Baru","penulis":"A","tahun":2020}' # 201
curl -i -X DELETE localhost:8089/api/buku/9999                                            # 404
```

## Latihan (starter)
1. **Refaktor Repo** — isi `repo_sql.go` (pindahkan query dari `sebelum/main.go`). Bonus: `Hapus`.
2. **Service + Validasi** — isi `service.go`; cek dengan `go test ./...`. Bonus: tolak judul duplikat.
3. **Rakit di main** — isi TODO di `main.go`; cek `go build` dan perilaku endpoint sama seperti `sebelum`. Bonus: `initApp()`.

## Kunci Latihan
Lihat `solution/`. Bonus duplikat judul ada di `service.go`; hasil uji lengkapnya di `service_test.go`.
