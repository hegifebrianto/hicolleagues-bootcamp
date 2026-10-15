# Sesi 8 — Latihan Hands-on (Daftar Tugas)

Kode untuk 3 latihan di slide "Latihan Hands-on · 25 Menit":

| Folder | Latihan | Konsep |
|---|---|---|
| — | Latihan 1 · 9' Daftar Tugas | v-for, @click, v-show, v-if |
| — | Latihan 2 · 8' Filter Daftar | computed |
| — | Latihan 3 · 8' Simpan Otomatis | watch + localStorage |

Satu app Vue yang sama berkembang lewat ketiga latihan ini (bukan 3 proyek
terpisah) — persis alur "pseudocode → kode → uji dengan contoh" di slide.

## Dua versi

- **`starter/`** — kerangka dengan 8 TODO berurutan di `src/App.vue`.
  Template HTML sudah lengkap, tinggal isi logika di bagian `<script>`.
  Aman dijalankan kapan pun meski TODO belum selesai (tidak akan crash).
- **`solution/`** — versi lengkap, termasuk 3 bonus dari slide:
  penghitung jumlah tugas, pesan "Tidak ada hasil", pesan "Tersimpan" sesaat.

## Cara menjalankan

```bash
cd starter      # atau: cd solution
npm install
npm run dev
```

Buka alamat yang muncul di terminal (biasanya `http://localhost:5173`).

**Catatan:** proyek ini belum sempat benar-benar dijalankan `npm install`
di sisi saya (lingkungan kerja saya tidak ada akses internet) — sudah
diperiksa ketat baris demi baris, tapi tetap coba jalankan dulu sebelum
dipakai demo supaya ada waktu kalau perlu perbaikan kecil.

## Cara uji tiap latihan (sesuai tabel di slide)

**Latihan 1:**
- Tambah "Belajar Vue" → muncul di daftar
- Klik Hapus → tugas hilang
- Input kosong lalu Tambah → tidak ditambahkan

**Latihan 2:**
- Ketik "bel" di kotak cari → tampil tugas yang mengandung "bel"
- Kosongkan kotak cari → semua tugas tampil lagi

**Latihan 3:**
- Tambah tugas, lalu refresh halaman → tugas tetap ada
- Hapus tugas, lalu refresh halaman → tugas tetap terhapus
