# Sesi 13 — REST API Buku (Go + MySQL)

Proyek ini dibuat **persis mengikuti isi slide Sesi 13** (`Sesi013_Go_Database_SQL_MySQL.pptx`):
HTTP server dengan `net/http`, routing dengan `ServeMux`, request/response
JSON, prinsip REST, lalu koneksi ke **MySQL** dan CRUD lewat `database/sql`
— semua jadi satu REST API untuk mengelola data **buku**.

## Endpoint

| Metode | Path | Fungsi | Materi slide |
|---|---|---|---|
| GET | `/api/buku` | Daftar semua buku (`?tahun_min=2020` opsional) | Request: Query, Response: JSON |
| POST | `/api/buku` | Tambah buku baru | Create: Parameter Aman, LastInsertId |
| PUT | `/api/buku/{id}` | Ubah buku (Tugas Mandiri di slide penutup) | Request: Path, Update |
| DELETE | `/api/buku/{id}` | Hapus buku | Update dan Delete |

## Sebelum menjalankan

1. Pastikan MySQL sudah terpasang dan berjalan di komputermu.
2. Jalankan isi `schema.sql` (lewat `mysql` CLI, TablePlus, DBeaver, atau
   phpMyAdmin) untuk membuat database `tugasdb` dan tabel `buku`,
   lengkap dengan 3 data contoh.
3. Sesuaikan DSN di `db.go` (user, password) dengan MySQL-mu:
   ```go
   dsn := "user:rahasia@tcp(localhost:3306)/tugasdb?parseTime=true"
   ```

## Cara menjalankan

```bash
cd solution      # atau: cd starter (setelah TODO dikerjakan)
go mod tidy       # mengunduh driver go-sql-driver/mysql
go run .
```

Server berjalan di `http://localhost:8080`.

**Catatan jujur:** sandbox saya tidak ada Go toolchain maupun akses
internet, jadi saya **tidak bisa menjalankan `go build` untuk
memverifikasi** kode ini sungguhan. Saya sudah audit manual ketat:
kurung dan brace seimbang di semua file, dan saya pastikan setiap
`import` benar-benar dipakai di kode nyata (bukan cuma di komentar
TODO) supaya versi starter tetap bisa di-*compile* walau TODO-nya
belum dikerjakan. Tapi tetap jalankan `go build ./...` sendiri dulu
sebelum dipakai mengajar, untuk jaga-jaga.

## Dua versi

- **`starter/`** — `main.go` (1 TODO: daftarkan 4 route),
  `db.go` (3 TODO: DSN, sql.Open, Ping), `buku.go` (14 TODO: isi
  keempat handler). Helper `tulisJSON`/`tulisError` sudah lengkap.
- **`solution/`** — lengkap, termasuk bonus `?tahun_min=` di GET
  `/api/buku` (sesuai contoh "WHERE tahun > 2020" di slide Latihan 1).

## Cara uji cepat (pakai curl)

```bash
# Daftar semua buku
curl http://localhost:8080/api/buku

# Daftar buku setelah tahun 2020
curl "http://localhost:8080/api/buku?tahun_min=2020"

# Tambah buku
curl -X POST http://localhost:8080/api/buku \
  -H "Content-Type: application/json" \
  -d '{"judul":"Atomic Habits","penulis":"James Clear","tahun":2018}'

# Ubah buku (ganti 1 dengan id yang sesuai)
curl -X PUT http://localhost:8080/api/buku/1 \
  -H "Content-Type: application/json" \
  -d '{"judul":"Laskar Pelangi (Edisi Revisi)","penulis":"Andrea Hirata","tahun":2005}'

# Hapus buku
curl -X DELETE http://localhost:8080/api/buku/1
```

| Skenario | Hasil diharapkan |
|---|---|
| POST tanpa judul | `400` + pesan "judul wajib diisi" |
| PUT/DELETE id yang tidak ada | `404` + pesan "buku tidak ditemukan" |
| POST berhasil | `201` + data buku lengkap dengan id baru |
| DELETE berhasil | `204`, tanpa body |
