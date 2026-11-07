# Sesi 17 — Vue: API, Loading & Form (demo + latihan)

Mengikuti deck "Vue: API, Loading & Form". Frontend memanggil backend `/api/tugas`.

## Menjalankan (2 terminal)
**Terminal 1 — backend** (pilih SALAH SATU; kontrak API-nya identik):
```bash
cd backend/go && go run .          # backend Go
# ATAU, kalau Go bermasalah saat demo:
node backend/mock-server.mjs       # cadangan Node, tanpa instalasi apa pun
```
Harus muncul banner `Alamat : http://localhost:8089`. Cek: `curl http://localhost:8089/api/ping` → `{"data":"pong"}`.

**Terminal 2 — frontend:**
```bash
cd solution        # demo lengkap   (starter/ = versi latihan dengan TODO)
npm install
npm run dev        # buka http://localhost:5173
```
Lampu "Backend terhubung" di atas halaman harus hijau.

> Port bentrok? `PORT=9000 go run .` lalu `API_TARGET=http://localhost:9000 npm run dev`.
> Backend menunda request GET 700 ms supaya "Memuat..." terlihat; ubah dengan `DELAY_MS=0` / `DELAY_MS=2000`.

## Peta ke slide
| Slide | Di kode |
|---|---|
| GET dengan Axios, `data.data` | `src/stores/tugas.js` → `muat()` |
| CORS solusi 1 (proxy) | `vite.config.js` |
| CORS solusi 2 (header di backend) | `cors()` di `backend/go/main.go` + `VITE_API_BASE` di `main.js` |
| try/catch/finally di action Pinia | `stores/tugas.js` |
| Loading & error di template, `onMounted` | `components/DaftarTugas.vue` |
| Form + validasi klien, POST / Latihan 2 | `components/FormTugas.vue` |
| PUT | centang tugas (`toggle` di `DaftarTugas.vue`) |
| DELETE + pesan per status / Latihan 3 | `hapus()` + `pesanGalat()` di `DaftarTugas.vue` |

## Skrip demo (±15 menit)
1. **Loading state**: muat ulang halaman → "Memuat..." (0,7 dtk) → daftar tampil. Klik *Muat ulang* untuk mengulang.
2. **Validasi klien**: ketik `ab` → pesan merah "Minimal 3 karakter", tombol nonaktif. Ketik `Belajar Vue` → *Kirim* → "Tugas ditambahkan", form kosong, daftar bertambah (POST).
3. **PUT**: centang tugas → di terminal backend muncul `PUT /api/tugas/..`; tampilan tercoret setelah daftar dimuat ulang.
4. **DELETE 204**: klik Hapus → konfirmasi → "Tugas dihapus".
5. **Status 404**: buka aplikasi di DUA tab. Tab 2: hapus tugas X. Tab 1 (belum di-refresh): hapus tugas X yang sama → "Tugas tidak ditemukan".
6. **Server mati**: hentikan backend (Ctrl+C) → *Muat ulang* → pesan error merah + lampu status merah. Nyalakan lagi → *Muat ulang* → pulih. (Lewat proxy Vite, server mati terlihat sebagai status 500 — itu normal.)
7. **CORS** (opsional, nyambung ke slide "CORS"): hentikan frontend, jalankan `VITE_API_BASE=http://localhost:8089 npm run dev` → tetap jalan karena backend mengirim header CORS. Lalu restart backend dengan `CORS_ORIGIN=http://salah.com` → di Console browser muncul error CORS. Itu yang ditunjukkan slide.

## Starter
`starter/` berisi 12 TODO (store `muat`, validasi + POST di `FormTugas`, `onMounted`/`pesanGalat`/DELETE di `DaftarTugas`). Aplikasi tetap tampil sebelum TODO dikerjakan.
Data backend ada di MEMORI — hilang saat backend di-restart (itu disengaja).
