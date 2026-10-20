# Project 1 Frontend — Aplikasi Vue To-Do (Sesi 10-11)

Kode ini dibuat **persis mengikuti slide "Rancangan Komponen" dan
"Ruang Lingkup Project 1 Vue"** di PPT Sesi 10 yang sudah Anda edit —
bukan proyek generik. Nama komponen, pola props/emit, dan urutan
fitur semuanya menyesuaikan slide tersebut.

## Catatan soal Sesi 12

Sama seperti sebelumnya: **Sesi 12 tidak punya proyek baru** — isinya
murni review Fase 1 dan pengantar Fase 2. Paket ini sudah mencakup
seluruh kebutuhan kode untuk Sesi 10-11.

## Arsitektur (sesuai slide)

```
App.vue (state di sini -- single source of truth)
  |-- TodoInput   : emit('tambah', teks)
  |-- TodoFilter  : props filterAktif, emit('ubah-filter', nilai)
  |-- TodoList    : props tugas (sudah disaring), meneruskan ke:
       |-- TodoItem : props tugas, emit('ubah'/'hapus', id)
```

## Pemetaan ke sesi (sesuai Checklist Workshop di slide)

| Sesi | Checklist Workshop | File yang disentuh |
|---|---|---|
| **10** (Workshop 1) | Proyek & komponen, tambah tugas, selesai & hapus, tampilkan jumlah | `TodoInput.vue`, `TodoItem.vue`, state dasar di `App.vue` |
| **11** (Workshop 2) | Filter, penyimpanan (localStorage), kerapian, gaya dasar | `onMounted` + `watch` di `App.vue`, `TodoFilter.vue` |

## Dua versi

- **`starter/`** — `TodoList.vue` dan `TodoFilter.vue` sudah LENGKAP
  (infrastruktur pendukung). `TodoInput.vue`, `TodoItem.vue`, dan
  `App.vue` punya TODO berurutan (9 TODO total) — kerjakan komponen
  anak dulu (TodoInput, TodoItem), baru lanjut ke App.vue, supaya
  event yang dipancarkan benar-benar sampai dan bisa diuji.
- **`solution/`** — versi lengkap, mencakup fitur wajib (Sesi 10) +
  fitur lanjutan filter & localStorage (Sesi 11).

## Cara menjalankan

```bash
cd solution      # atau: cd starter
npm install
npm run dev
```

**Catatan yang sama seperti proyek-proyek sebelumnya:** sandbox saya
tidak punya akses internet, jadi belum sempat `npm install` secara
nyata di sisi saya — sudah diperiksa ketat baris demi baris dan lolos
audit sintaks otomatis (10 file, semua seimbang), tapi sebaiknya
dicoba dulu sebelum dipakai mengajar.

## Checklist fitur (sesuai "Ruang Lingkup Project 1 Vue")

- [x] Tambah tugas
- [x] Tandai selesai
- [x] Hapus tugas
- [x] Hitung tugas (tersisa)
- [x] Filter Semua / Aktif / Selesai
- [x] Penyimpanan localStorage
- [ ] Edit tugas (disebut di slide sebagai fitur pilihan/bonus —
      belum diimplementasikan, bisa jadi tantangan tambahan untuk
      peserta yang lebih cepat selesai)

## Cara uji cepat (sesuai tabel "Skenario Uji Manual" di slide lama)

| Skenario | Aksi | Hasil diharapkan |
|---|---|---|
| Tambah | Ketik lalu Enter | Tugas muncul |
| Selesai | Klik checkbox | Tercoret, "tugas tersisa" turun |
| Filter aktif | Klik "Aktif" | Hanya tugas belum selesai yang tampil |
| Refresh | F5 | Data tetap ada |
