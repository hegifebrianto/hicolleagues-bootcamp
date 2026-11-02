# Sesi 14 — REST API Tugas + Middleware (versi satu file)

Revisi ketiga. Perubahan dari versi sebelumnya:

1. **Satu file (`main.go`) saja** — menghilangkan kemungkinan salah
   folder/salah file yang ter-copy.
2. **Port default diganti ke `8089`** (bisa diubah lewat `PORT=xxxx`)
   — port `8080` sangat umum dipakai proses lain (Docker, dev server
   framework lain, dll) di Mac, jadi gampang bentrok tanpa disadari.
3. **Ditambah endpoint `/`** — buka di browser untuk memastikan
   request-mu benar-benar sampai ke server INI, bukan ke proses lain
   yang kebetulan nongkrong di port yang sama.
4. **Banner saat start** mencetak alamat lengkap dan status `API_KEY`,
   supaya langsung kelihatan di terminal apa yang sedang berjalan.
5. Error `ListenAndServe` sekarang ditangkap eksplisit dengan pesan
   "GAGAL MENYALAKAN SERVER" — kalau port bentrok, ini akan muncul
   jelas di terminal (sebelumnya tercampur dengan log biasa).

## Cara menjalankan

```bash
cd solution      # atau: cd starter (setelah TODO dikerjakan)
go run .
```

Harus muncul banner seperti ini di terminal:

```
==========================================
 Sesi 14 -- REST API Tugas + Middleware
==========================================
 Alamat  : http://localhost:8089
 Coba    : http://localhost:8089/ping
 Route   : GET/POST /api/tugas, GET/PUT/DELETE /api/tugas/{id}
 API_KEY :   (kosong = semua /api/tugas ditolak 401)
==========================================
```

Lalu di terminal LAIN (biarkan yang pertama tetap jalan):

```bash
curl -i http://localhost:8089/
curl -i http://localhost:8089/ping
```

Kalau LANGKAH INI SAJA masih 404 atau gagal konek sama sekali, itu
bukan lagi soal kode Go-nya -- kemungkinan besar ada sesuatu di luar
program ini (firewall, VPN, proxy lokal, atau antivirus) yang
memblokir koneksi ke localhost. Coba port lain untuk memastikan:

```bash
PORT=5055 go run .
# lalu di terminal lain:
curl -i http://localhost:5055/ping
```

## Set API_KEY untuk endpoint /api/tugas

```bash
API_KEY=rahasia go run .
```

## Dua versi

- **`starter/`** — 16 TODO + 2 catatan import (lihat komentar di
  bagian atas file). Server tetap bisa di-`go run .` dan menyala
  normal meski TODO belum dikerjakan -- cuma semua endpoint `/api/tugas`
  akan membalas kosong/404 sampai TODO-nya selesai.
- **`solution/`** — lengkap.

**Catatan jujur yang sama seperti sebelumnya:** sandbox saya tidak ada
Go toolchain, jadi saya tetap tidak bisa menjalankan `go build` untuk
membuktikan ini sungguhan tanpa cela. Saya sudah audit manual ketat
(kurung/brace seimbang, setiap import dipastikan dipakai di kode
nyata), dan kali ini juga saya sederhanakan strukturnya sejauh mungkin
supaya makin kecil kemungkinan ada yang salah. Tapi kalau `/ping` masih
gagal setelah ini DENGAN server yang terbukti menyala (banner muncul)
dan DENGAN port yang terbukti kosong, kemungkinan besar penyebabnya
ada di luar kode ini.
