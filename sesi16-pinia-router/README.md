# Sesi 16 — Pinia & Vue Router (demo + latihan)

Mengikuti deck "Vue: Pinia & Vue Router". Tidak butuh backend.

## Menjalankan
```bash
cd solution        # demo lengkap   (starter/ = versi latihan dengan TODO)
npm install
npm run dev        # buka http://localhost:5173
```

## Peta ke slide
| Slide | Di kode |
|---|---|
| Membuat store Pinia / Latihan 1 | `src/stores/tugas.js` (state, getters `jumlah` & `selesai`, actions `tambah`/`hapus`/`toggle`) |
| Memakai store di komponen, `storeToRefs` | `views/Beranda.vue`, `views/DaftarTugas.vue` |
| Setup Vue Router, Latihan 2 | `src/router/index.js`, `App.vue` (`RouterLink` + `RouterView`) |
| Navigasi programatik `router.push` | `App.vue` (tombol Keluar), `views/Login.vue` |
| Route params / Latihan 3 | `/tugas/:id` → `views/DetailTugas.vue` (`Number(route.params.id)`) |
| Nested routes `children` | `/admin` → `AdminLayout.vue` + `DaftarPengguna.vue` / `Statistik.vue` |
| Navigation guard + `meta` | `router.beforeEach` (rute `/tugas`, `/tugas/:id`, `/admin` butuh login) |
| Bonus: 404, localStorage | rute `/:pathMatch(.*)*`; blok `SIMPAN_KE_LOCALSTORAGE` di `main.js` |

## Skrip demo (±12 menit)
1. **State global**: di Beranda perhatikan angka tugas; menu "Tugas" juga punya badge. Buka Tugas, tambah/centang/hapus → angka di menu & Beranda ikut berubah tanpa props/emit.
2. **Guard**: klik "Tugas" saat belum login → diarahkan ke `/login?redirect=/tugas` → klik *Masuk (simulasi)* → otomatis kembali ke `/tugas`. Klik *Keluar* lalu coba `/admin` → dicegat lagi.
3. **Params**: klik judul tugas → `/tugas/<id>`. Ketik manual `/tugas/999` → "Tidak ditemukan". Ketik `/ngawur` → halaman 404.
4. **Nested routes**: `/admin` → ganti sub-menu Pengguna/Statistik; hanya bagian bawah yang berganti, layout tetap.
5. **SPA**: buka tab Network, pindah halaman → tidak ada reload dokumen.
6. **Bahan soft skill RCA**: refresh browser → data tetap (tersimpan di localStorage). Ubah `SIMPAN_KE_LOCALSTORAGE = false` di `main.js`, simpan, tambah beberapa tugas, lalu refresh → data kembali ke 2 tugas awal ("Pinia state hilang saat refresh") → peserta lakukan *5 Why* + riset mandiri (akar masalah: state Pinia hidup di memori halaman, bukan penyimpanan permanen).
7. **Contoh 5 Why di slide**: di `DetailTugas.vue` hapus `Number(...)` → halaman detail selalu kosong; kembalikan lagi.

## Starter
`starter/` berisi 13 TODO (store, rute, guard, DetailTugas, persist). Aplikasi tetap jalan sebelum TODO dikerjakan (stub): menu tampil, data kosong, rute selain `/` belum ada.
