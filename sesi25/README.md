# Sesi 25 · Workshop Project 3 Backend (Go) — Auth JWT + RBAC

Aplikasi Buku dari Sesi 17/21, sekarang dengan **login, token JWT, middleware auth, role admin**,
**email selamat datang async (goroutine)**, dan **test hijau tanpa race condition**.

```
solution/   versi lengkap (acuan trainer)
starter/    infrastruktur sudah ada; 4 TODO untuk peserta (test dibiarkan MERAH di awal)
```

## Menjalankan (solution / starter)

```bash
cd solution
go mod tidy                 # sekali saja: mengunduh driver MySQL
go run .                    # perlu MySQL + schema.sql
STORE=memory go run .       # tanpa MySQL (data hilang saat berhenti) -> paling aman untuk demo
```
Windows PowerShell: `$env:STORE="memory"; go run .`

MySQL: `mysql -u root -p < schema.sql`, lalu isi `DB_PASSWORD` di `.env`.
**Ganti `JWT_SECRET` di `.env`** (min. 16 karakter; `openssl rand -hex 32`). Tanpa itu server menolak start.
Akun demo (password `rahasia`): `admin@mail.com` (admin), `budi@mail.com` (user). Port default **8089**.

## Peta slide → file

| Slide | File |
|---|---|
| Layered architecture | `handler/` → `service/` → `repository/` (+ `model/`, `respon/`) |
| Anatomi JWT, Membuat & Memverifikasi Token | `auth/token.go` |
| Praktik Aman JWT (secret dari .env, hash password) | `config/config.go`, `.env`, `auth/password.go` |
| Middleware Auth, Memasang Middleware | `middleware/middleware.go` (`Auth`), `handler/router.go` |
| RBAC | `middleware.RequireRole`, rute `DELETE /api/buku/{id}` |
| Async welcome email | `service/auth.go` (`go s.Mailer...`), `service/mailer.go` |
| Unit test service + mock | `service/*_test.go` |
| Test handler dengan httptest | `handler/handler_test.go`, `middleware/middleware_test.go` |

## Demo curl (401 → 403 → 204)

```bash
B=http://localhost:8089
curl -i -X DELETE $B/api/buku/1                         # 401 tanpa token

curl -s -X POST $B/api/login -d '{"email":"budi@mail.com","password":"rahasia"}'
# salin nilai "token"
curl -i -X DELETE -H "Authorization: Bearer <TOKEN_USER>"  $B/api/buku/1   # 403 role user
# login sebagai admin@mail.com, lalu:
curl -i -X DELETE -H "Authorization: Bearer <TOKEN_ADMIN>" $B/api/buku/1   # 204
```
Daftar user baru: `POST /api/register` → balas cepat (201); lihat log server ±2 detik kemudian: `[email] selamat datang ...`. 
<!-- sudah bisa -->
Cek identitas dari token: `GET /api/me`.
<!-- dengan menggunakan token nya kita bisa mendapatkan roles nya -->
## Test

```bash
go test ./...           # semua
go test -race ./...     # tanpa race condition (butuh gcc/cgo; Windows: CGO_ENABLED=1 + MinGW)
go test -cover ./...
```

## Starter: 4 TODO (urutan disarankan)

1. `auth/token.go` — `Buat` dan `Cek` → test `auth` hijau
2. `middleware/middleware.go` — `Auth` dan `RequireRole` → test `middleware` hijau
3. `service/auth.go` — `Login` → test login hijau
4. `service/auth.go` — kirim email di goroutine (`go`) → `TestDaftarEmailAsync` hijau

Awalnya `go test ./...` merah; target akhir hijau penuh. Handler test baru hijau setelah keempatnya selesai.

## Catatan penting (beda dengan slide)

- **JWT** di project ini dibuat dengan standard library (`crypto/hmac`) supaya tanpa dependensi dan strukturnya terlihat.
  Slide memakai library `golang-jwt/jwt`; itu padanan yang dipakai di dunia nyata. Kode slide hanya diuji kompilasinya
  terhadap stub, **bukan** dijalankan dengan library aslinya.
- **Hash password** memakai PBKDF2-SHA256 dari standard library (`crypto/pbkdf2`, butuh **Go 1.24+**), bukan bcrypt seperti di slide.
  Konsepnya sama (salt acak, lambat, satu arah). Untuk produksi pakai bcrypt/argon2.
- Token hanya di-encode, **tidak dienkripsi**: jangan taruh data rahasia di payload.
- Mode MySQL (`repository/mysql.go`) belum pernah dijalankan dengan MySQL asli di lingkungan pembuatan
  (diuji dengan driver palsu). Coba sekali sebelum kelas; kalau ada masalah, pakai `STORE=memory`.
- Hash di `schema.sql` dibuat dengan `go run ./cmd/buathash rahasia`.
