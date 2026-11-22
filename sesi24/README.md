# Project — Concurrency di Go · Sesi 24 (Fase 3)

Goroutine, channel, sync (WaitGroup & Mutex), race condition + `-race`, dan context.
Tidak butuh database, internet, atau library tambahan (hanya standard library; tidak perlu `go mod tidy`).

| Folder | Isi |
|---|---|
| `solution/` | Trainer: semua contoh slide, kunci Latihan 1–3, dan kunci Tugas Mandiri (+ test) |
| `starter/` | Peserta: contoh yang sama, Latihan 1–3 berisi `TODO`, tugas mandiri berupa test yang harus dihijaukan |

Setiap program berada di folder sendiri (tiap program butuh `main()` sendiri), dijalankan dengan `go run ./folder`.

## Peta slide → program
| Slide | Jalankan | Hasil yang diharapkan |
|---|---|---|
| Goroutine: Ringan dan Mudah | `go run ./contoh/01_goroutine` | 2 baris Goroutine (urutan bisa beda) + `Selesai` |
| Channel | `go run ./contoh/02_channel` | `Halo dari goroutine` |
| Buffered Channel & Worker | `go run ./contoh/03_worker` | `2` lalu `4` |
| sync.WaitGroup | `go run ./contoh/04_waitgroup` | Tugas 1–3 (urutan bisa beda) + `Semua selesai` |
| sync.Mutex | `go run ./contoh/05_mutex` | `1000` |
| (demo bug) Race Condition | `go run ./contoh/06_race_bug` | angka **berubah-ubah**, bukan 1000 |
| Mendeteksi Race | `go run -race ./contoh/06_race_bug` | `WARNING: DATA RACE` + baris konflik |
| context | `go run ./contoh/07_context` / `... 3s` | `Selesai tepat waktu` / `Timeout: context deadline exceeded` |
| Latihan 1 | `go run ./latihan/lat1_waitgroup` | 5 baris + waktu total (bonus) |
| Latihan 2 | `go run ./latihan/lat2_channel` | 1 4 9 16 25 |
| Latihan 3 | `go run -race ./latihan/lat3_race` | solusi: 1000, tanpa WARNING |
| Tugas Mandiri | `go test -race ./tugas` | PASS |

## Demo trainer (±20 menit)
1. **Goroutine**: jalankan `01_goroutine` 3x (urutan berubah). Lalu hapus `time.Sleep` → tidak ada yang tercetak → alasan butuh WaitGroup.
2. **Channel**: `02_channel`. Tanyakan: "Apa yang terjadi kalau `ch <- ...` tidak ada penerimanya?" → `fatal error: all goroutines are asleep - deadlock!` (demokan dengan menghapus baris `<-ch`).
3. **Worker**: `03_worker`. Pindahkan `close(jobs)` → hapus → deadlock; jelaskan `close` + `range`.
4. **WaitGroup**: `04_waitgroup`; tunjukkan kenapa `n` dikirim sebagai parameter.
5. **Race**: jalankan `06_race_bug` 3–4x (41, 64, 37, ...) → lalu `-race` → baca laporannya (baris mana yang membaca/menulis) → buka `05_mutex` sebagai perbaikan.
6. **context**: jalankan `07_context` tanpa dan dengan argumen `3s`.
7. **Tugas mandiri**: tunjukkan `go test -v ./tugas`: log `berurutan=~300ms konkuren=~60ms`.

## Catatan penting
- **`total++` polos sering tetap menghasilkan 1000** di komputer cepat (goroutine selesai terlalu cepat sehingga jarang bertabrakan), padahal race-nya ada. Karena itu `06_race_bug` memecah `total++` jadi baca → `runtime.Gosched()` → tulis supaya bug-nya terlihat jelas. `-race` selalu melaporkannya, itulah alasan flag ini penting.
- **`-race` butuh cgo (compiler C).** Di macOS dan Linux biasanya langsung jalan. Di **Windows** bila muncul `go: -race requires cgo`, pasang gcc (mis. MSYS2/TDM-GCC) lalu `set CGO_ENABLED=1` (PowerShell: `$env:CGO_ENABLED=1`). Alternatif tanpa gcc: demokan di Mac/Linux atau pakai bukti angka berubah-ubah dari `06_race_bug`.
- Di Windows PowerShell, jika `go test` atau `npm` ditolak karena Execution Policy, lihat solusi `Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy RemoteSigned`.
- Di slide "Solusi Latihan 2" tanda kurung kurawal dan baris kodenya tampak tercampur (hasil ekspor slide), tetapi pola yang dimaksud: kirim hasil di goroutine, dan `go func() { wg.Wait(); close(hasil) }()` untuk menutup channel setelah SEMUA pengirim selesai. `close` tidak boleh dipanggil sebelum pengirim selesai (panic `send on closed channel`), dan tanpa `close` loop `range` menunggu selamanya (deadlock).

## Latihan (starter)
1. **Goroutine & WaitGroup** — `latihan/lat1_waitgroup/main.go`. Bonus: `time.Since`.
2. **Channel hasil** — `latihan/lat2_channel/main.go`. Bonus: `sort.Ints`.
3. **Perbaiki race** — `latihan/lat3_race/main.go`: jalankan dulu (`go run -race`) untuk melihat masalah, lalu tambah `sync.Mutex`. Uji: `go test -race ./latihan/lat3_race`. Bonus: versi channel tanpa Mutex (lihat `hitungDenganChannel` di solution).

## Tugas mandiri: ambil data secara konkuren
`starter/tugas/ambil.go` → isi `AmbilKonkuren` sampai `go test -race ./tugas` lulus. Syarat (sudah dicek test):
urutan hasil terjaga, jauh lebih cepat dari berurutan, error pertama membatalkan yang lain, patuh pada timeout `context`, dan bebas race.
Lalu jalankan `go test -race ./...` pada seluruh project (poin kedua tugas dari slide).
