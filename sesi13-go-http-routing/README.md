# Go HTTP Server & Routing — Sesi 10

Contoh project untuk HiColleagues Bootcamp Batch 4. Project menggunakan **Go 1.22+**, `net/http`, `http.ServeMux`, JSON, status code, dan Git dasar. Tidak membutuhkan framework atau database eksternal.

## Struktur folder

```text
go-http-routing-sesi13/
├── cmd/server/main.go
├── internal/
│   ├── handler/
│   │   ├── handler.go
│   │   └── handler_test.go
│   └── tugas/tugas.go
├── go.mod
└── README.md
```

## Menjalankan server

Pastikan Go 1.22 atau lebih baru sudah terpasang.

```bash
go version
go run ./cmd/server
```

Server berjalan di `http://localhost:8080`. Tekan `Ctrl+C` untuk menghentikannya.

## Endpoint

| Method | Path | Keterangan | Status sukses |
|---|---|---|---:|
| GET | `/ping` | Health check | 200 |
| GET | `/waktu` | Waktu server format RFC3339 | 200 |
| GET | `/api/salam?nama=Budi` | Salam dalam Bahasa Indonesia | 200 |
| GET | `/api/salam?nama=Budi&bahasa=en` | Salam Bahasa Inggris | 200 |
| GET | `/api/tugas` | Daftar semua tugas | 200 |
| POST | `/api/tugas` | Membuat tugas | 201 |
| GET | `/api/tugas/{id}` | Detail tugas | 200 |
| DELETE | `/api/tugas/{id}` | Menghapus tugas | 204 |

Data tugas disimpan di memory; data akan hilang saat server dihentikan.

## Contoh request

### 1. Ping

```bash
curl -i http://localhost:8080/ping
```

Response: `200 OK`

```json
{"pesan":"pong"}
```

### 2. Salam

```bash
curl -i 'http://localhost:8080/api/salam?nama=Budi'
curl -i 'http://localhost:8080/api/salam?nama=Budi&bahasa=en'
curl -i http://localhost:8080/api/salam
```

Parameter `nama` wajib diisi. Jika kosong, server mengembalikan `400 Bad Request`.

### 3. Membuat tugas

```bash
curl -i -X POST http://localhost:8080/api/tugas \
  -H 'Content-Type: application/json' \
  -d '{"judul":"Belajar net/http"}'
```

Contoh response `201 Created`:

```json
{"id":1,"judul":"Belajar net/http","selesai":false}
```

JSON rusak atau judul kosong menghasilkan `400 Bad Request`.

### 4. Daftar dan detail tugas

```bash
curl -i http://localhost:8080/api/tugas
curl -i http://localhost:8080/api/tugas/1
```

### 5. Hapus tugas

```bash
curl -i -X DELETE http://localhost:8080/api/tugas/1
```

Berhasil menghapus mengembalikan `204 No Content`. Jika ID tidak ditemukan, server mengembalikan `404 Not Found`.

## Menjalankan test

```bash
go test ./...
```

## Konsep yang dipakai

- **Handler** menerima `http.ResponseWriter` dan `*http.Request`.
- **ServeMux** mendaftarkan route menggunakan pola metode dan path, fitur Go 1.22+.
- **Path parameter** diambil dengan `r.PathValue("id")`.
- **Query parameter** diambil dengan `r.URL.Query().Get("nama")`.
- **JSON request** dibaca menggunakan `json.Decoder`.
- **JSON response** dikirim dengan `json.Encoder`.
- **Status code**: 200, 201, 204, 400, dan 404.
- Repository menggunakan mutex agar akses data in-memory aman saat ada request bersamaan.

## Git dasar

Jalankan dari root project:

```bash
git init
git status
git add .
git commit -m "Tambah HTTP server dan routing"
git branch -M main
git remote add origin https://github.com/USERNAME/NAMA-REPO.git
git push -u origin main
```

Ganti URL remote dengan URL repository GitHub milikmu. Jika repository sudah memiliki remote, periksa dengan `git remote -v` sebelum menambahkannya.

## Tugas mandiri

1. Tambahkan field `Selesai` agar bisa diubah melalui endpoint baru.
2. Tambahkan validasi untuk request JSON.
3. Tambahkan unit test untuk endpoint salam dan pembuatan tugas.
4. Commit dan push perubahan ke GitHub.
