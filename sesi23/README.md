# Project 2 — Testing di Go · Sesi 23 (Fase 3)

Menguji kode API Buku berlapis dari Sesi 23: fungsi murni, table-driven test, mock/stub, dan integration test.

| Folder | Fungsi |
|---|---|
| `starter/` | Untuk peserta: kode aplikasi lengkap + test Latihan 1–3 berstatus `TODO` (SKIP). |
| `solution/` | Untuk trainer: semua test lengkap + bonus (httptest, siklus integration test). |

## Persiapan
```bash
cd solution                 # atau starter
go mod tidy                 # sekali saja (unduh driver MySQL, butuh internet)
go test ./...
```
Unit test **tidak butuh MySQL**. Hanya integration test yang butuh (lihat bagian bawah).
`go test ./matematika` bisa dijalankan bahkan tanpa `go mod tidy`.

## Peta slide → file (solution)
| Slide | File |
|---|---|
| Test Pertama (`TestTambah`) | `matematika/matematika_test.go` |
| `t.Error` vs `t.Fatal` (`TestBagi`) | `matematika/matematika_test.go` |
| Table-Driven Tests | `matematika_test.go` (`TestTambahBerbagaiKasus`), `minmax_test.go` (`TestMinMax`) |
| Latihan 1 & 2 (minMax) | `minmax.go`, `minmax_test.go` |
| Mock dan Stub (`repoPalsu`) | `repo_palsu_test.go`, `service_test.go` |
| Latihan 3 + bonus hitung `Simpan` | `service_test.go` (`simpanDipanggil`) |
| Integration Test + build tag | `repo_sql_integration_test.go`, `schema_test.sql` |

## Perintah (slide "Coverage & Perintah Berguna")
```bash
go test ./...                       # semua test
go test -v ./...                    # detail tiap test & sub-test
go test -cover ./...                # persentase cakupan
go test -run TestMinMax -v .        # satu test (dan sub-test-nya)
go test -run 'TestMinMax/negatif' -v .   # satu sub-test saja
go test -tags=integration ./...     # termasuk integration test
```

### Tentang cakupan (tugas mandiri: >70% untuk service)
Slide menyebut "package service", tapi di project ini service ada di **package main** bersama file lain
(`main.go`, `repo_sql.go`, dst.), jadi `go test -cover` melaporkan angka package keseluruhan (±48%).
Untuk melihat cakupan per fungsi/file service:
```bash
go test -coverprofile=c.out .
go tool cover -func=c.out | grep service.go     # solution: semua fungsi 100%
go tool cover -html=c.out                       # tampilan HTML: baris merah = belum teruji
```

## Catatan penyesuaian dari slide
- Slide "Mock dan Stub" memakai `repoPalsu` dengan 2 method (`Semua`, `Simpan`). Interface `BukuRepo` kita punya **4 method**
  (`Semua`, `Ambil`, `Simpan`, `Hapus`), dan Go mewajibkan semuanya ada, jadi `repoPalsu` mengimplementasikan keempatnya.
- Slide menguji `Tambah(Buku{Judul: "Go"})` lalu mengharapkan sukses. Service kita menolak `tahun <= 1900`,
  jadi contoh dibuat `Buku{Judul: "Go", Penulis: "Andi", Tahun: 2020}`.
- `minMax` untuk slice kosong dibuat mengembalikan `0, 0` (kasus bonus).
- Memakai **MySQL** (bukan PostgreSQL) sesuai Sesi 13 dst.

## Demo trainer (±15 menit)
1. **Test pertama**: `cd matematika && go test -v` → PASS. Lalu ubah `Tambah` jadi `a - b` → **FAIL** dengan pesan `Tambah(2, 3) = -1, ingin 5`. Kembalikan.
2. **Fatal vs Error**: di `TestBagi`, ganti `t.Fatal` jadi `t.Error` dan hapus pengaman → tunjukkan bedanya saat error tidak muncul.
3. **Table-driven**: tambah satu baris kasus di `TestMinMax`, jalankan `go test -v -run TestMinMax`, lihat sub-test baru muncul otomatis.
4. **Test menangkap bug**: di `service.go` ubah `b.Tahun <= 1900` jadi `< 1900` → `TestTambah_Validasi/tahun_1900_(batas)` GAGAL. Kembalikan.
5. **Mock**: tunjukkan `service_test.go` jalan tanpa MySQL & tanpa server; jelaskan bahwa `repoPalsu` bisa disuntik karena Service hanya mengenal interface (Sesi 17).
6. **Integration**: `mysql -u root -p < schema_test.sql`, isi `.env`, lalu `go test -tags=integration -v -run RepoSQL .`
   Tanpa tag → tidak jalan. Dengan tag tapi MySQL mati → **SKIP** (bukan gagal).

## Latihan (starter)
1. **Test fungsi murni** — `minmax_test.go`: `[4,9,2,7]`→2,9; `[5]`→5,5. Bonus: slice kosong.
2. **Table-driven** — ubah jadi tabel ≥4 kasus + `t.Run`. Bonus: duplikat nilai.
3. **Test dengan mock** — `repo_palsu_test.go` + `service_test.go`: judul kosong → `ErrValidasi`; judul valid → tersimpan. Bonus: hitung panggilan `Simpan`.

Hapus `t.Skip` di file starter saat mulai mengisi. Kunci jawaban: file yang sama di `solution/`.

## Integration test: butuh database terpisah
Test ini membuat dan menghapus data, jadi **jangan** diarahkan ke `bukudb` yang berisi data asli.
Gunakan `bukudb_test` (dibuat oleh `schema_test.sql`). Kredensial dibaca dari `.env`;
nama database test bisa diganti lewat `DB_NAME_TEST`.

## Tugas mandiri
Tulis test untuk semua fungsi service (sudah dicontohkan di `solution/service_test.go`), capai cakupan >70% di `service.go`,
dan terapkan satu masukan dari sesi code review.
