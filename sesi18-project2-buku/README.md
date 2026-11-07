# Sesi 18 — Hands-on Project 2: Frontend RESTful (Buku)

Mengikuti deck "Hands-on Project 2 Frontend RESTful": daftar, tambah, ubah, hapus buku lewat Vue Router + Pinia + API.
Struktur folder sama dengan slide: `router/`, `stores/buku.js`, `views/`, `components/KartuBuku.vue`.

## Menjalankan (2 terminal)
**Terminal 1 — backend** (pilih SALAH SATU; kontrak API identik):
```bash
cd backend/go && go run .          # backend Go
node backend/mock-server.mjs       # ATAU cadangan Node (tanpa instalasi apa pun)
```
Cek: `curl http://localhost:8089/api/ping` → `{"data":"pong"}`.

**Terminal 2 — frontend:**
```bash
cd solution        # demo lengkap   (starter/ = versi latihan dengan TODO)
npm install
npm run dev        # buka http://localhost:5173
```
> Port bentrok? `PORT=9000 go run .` lalu `API_TARGET=http://localhost:9000 npm run dev`. Jeda GET diatur `DELAY_MS`.

> Catatan: slide memakai `/api/buku` pada Go port 8080 dari Sesi 13 (MySQL). Backend demo ini sengaja
> berdiri sendiri (data di memori, port 8089) supaya demo tidak bergantung MySQL. Kontraknya sama:
> `{ "data": ... }`, field `judul`, `penulis`, `tahun`.

## Peta ke slide
| Slide | Di kode |
|---|---|
| Spesifikasi: daftar, tambah/edit, hapus, router + Pinia | seluruh `src/` |
| Arsitektur: View → Store → API | `views/*.vue` memanggil `stores/buku.js` |
| Router `/`, `/tambah`, `/ubah/:id`, `/detail/:id` | `src/router/index.js` |
| Store dengan CRUD | `stores/buku.js` (`muat`, `ambil`, `tambah`, `ubah`, `hapus`) |
| Form ubah dengan route param (pola sama tambah & ubah) | `views/FormBuku.vue` |
| Checklist Workshop 1 / 2 | daftar+loading+error → form tambah → edit → hapus → kerapian |

## Skenario uji manual (tabel di slide)
| Skenario | Aksi | Hasil |
|---|---|---|
| Muat daftar | buka `/` | "Memuat..." lalu 3 buku dari API |
| Tambah valid | `/tambah`, isi lengkap, Simpan | kembali ke daftar, buku baru muncul |
| Tambah tidak valid | kosongkan judul, Simpan | pesan error di bawah field, tidak terkirim (cek terminal backend: tidak ada POST) |
| Ubah | klik *Ubah* | form terisi data lama; Simpan → PUT → daftar terbarui |
| Hapus | klik *Hapus* | konfirmasi → "Buku ... dihapus" |
| Hapus saat sudah hilang (404) | hapus buku yang sama dari 2 tab | "Buku tidak ditemukan (mungkin sudah dihapus)" |
| Detail tidak ada | buka `/detail/999` | "Buku tidak ditemukan." |
| Backend mati | Ctrl+C backend, *Muat ulang* | pesan error jelas + lampu status merah |

## Starter
`starter/` berisi 15 TODO: store (8), `DaftarBuku` (3), `FormBuku` (4: validasi, muat data lama, simpan). `KartuBuku`, `DetailBuku`, router, dan App sudah lengkap.
Data di memori backend hilang saat restart — disengaja.
