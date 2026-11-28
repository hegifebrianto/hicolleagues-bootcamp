# Sesi 26 · Vue: Komponen Lanjutan

Satu aplikasi "lab" dengan 4 menu, masing-masing memperagakan satu bagian slide.
Tidak butuh backend: API (login + daftar tugas) dipalsukan di dalam browser (`src/api/mock.js`).

```
solution/   versi lengkap (acuan trainer, untuk demo)
starter/    sama persis, tapi 4 bagian dikosongkan (TODO) untuk dikerjakan peserta
```

## Menjalankan

```bash
cd solution          # atau starter
npm install
npm run dev          # buka http://localhost:5173
```
Windows PowerShell yang memblokir `npm`: pakai `npm.cmd install` / `npm.cmd run dev`
atau `Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy RemoteSigned`.

Akun demo (menu Pinia): `demo@mail.com` / `rahasia`.

## Peta slide → file

| Slide / bagian | Menu | File |
|---|---|---|
| Lifecycle Hooks Utama | 1 · Lifecycle | `components/Jam.vue`, `views/LifecycleDemo.vue` |
| Membersihkan Efek Samping | 1 · Lifecycle | `Jam.vue` (`onUnmounted`), `utils/pantau.js` |
| Slot Default / Named / Scoped | 2 · Slots | `components/Kartu.vue`, `Modal.vue`, `Daftar.vue`, `views/SlotDemo.vue` |
| Dynamic Components, keep-alive | 3 · Dinamis | `views/DinamisDemo.vue`, `components/Profil.vue`, `Pengaturan.vue` |
| Pinia Modules | 4 · Pinia | `stores/auth.js`, `stores/tugas.js`, `stores/ui.js`, `views/PiniaDemo.vue` |

## Skenario demo (±2 menit per menu)

1. **Lifecycle**: hilangkan centang "Bersihkan timer", sembunyikan lalu tampilkan Jam beberapa kali →
   "Timer yang masih hidup" naik terus (memory leak). Centang lagi → kembali ke 0 saat Jam disembunyikan.
   Log di bawah menunjukkan urutan `onBeforeMount → onMounted → onUpdated → onUnmounted`.
2. **Slots**: dua Kartu memakai konten berbeda; kartu tanpa header/footer tidak merender elemen kosong
   (`$slots.header`). Modal: named slot + scoped slot (`#footer="{ tutup }"`). Daftar: scoped slot (`{ item, index }`).
3. **Dinamis**: isi form Pengaturan, pindah tab, kembali. Dengan keep-alive isian tetap & `mounted` tidak bertambah
   (hanya `activated`). Matikan checkbox keep-alive → isian hilang & `mounted` naik.
4. **Pinia**: "Muat tugas" sebelum login → 401. Login → toast (store `ui` dipanggil dari store `auth`) →
   muat tugas berhasil (header `Authorization` dari store `auth`). Refresh: tetap login (localStorage).

## Latihan peserta (di `starter/`) dan kuncinya

| Latihan | Berkas dengan TODO | Selesai bila |
|---|---|---|
| 1 · Kartu dengan Slot (8') | `components/Kartu.vue` | Menu Slots: kartu "Profil" & "Statistik" menampilkan header & isinya sendiri; footer muncul hanya di kartu kedua |
| 2 · Tab Dinamis (8') | `views/DinamisDemo.vue` | Menu Dinamis: isian form bertahan saat pindah tab |
| 3 · Store auth (9') | `stores/auth.js`, `stores/tugas.js` | Menu Pinia: login berhasil, "Muat tugas" menampilkan 4 tugas |
| Bonus slide "Membersihkan Efek Samping" | `components/Jam.vue` | Timer aktif kembali ke 0 setelah Jam disembunyikan |

Kunci jawaban = berkas yang sama di `solution/`. Di starter, checkbox "Bersihkan timer" tidak berpengaruh
(sengaja: pembersihan memang belum ada).

## Catatan jujur

- **Belum dijalankan dengan Vue/Vite sungguhan** di lingkungan pembuatan (tidak ada akses internet untuk `npm install`).
  Yang sudah diperiksa: logika store + API palsu dijalankan di Node (login salah/benar, 401 tanpa token, 4 tugas dengan token,
  logout), dan sintaks semua blok `<script>` dicek. **Template `.vue` hanya ditinjau manual.**
  Jalankan `npm install && npm run dev` sekali sebelum kelas dan klik keempat menu.
- Server palsu hanya tahu `POST /api/login` dan `GET /api/tugas`. Backend Go Sesi 25 tidak punya `/api/tugas`,
  jadi project ini sengaja tidak dihubungkan ke sana.
- Slide memakai `axios.get(...)` langsung; di sini dipakai instance `api` (`api/client.js`) agar bisa memakai server palsu.
  Bentuk pemanggilannya sama.
