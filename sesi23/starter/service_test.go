package main

import "testing"

// LATIHAN 3 (8 menit): test BukuService.Tambah dengan REPO PALSU (tanpa database).
//
// Langkah:
//  1. Buat struct repoPalsu yang mengimplementasikan SEMUA method BukuRepo
//     (lihat repo.go: Semua, Ambil, Simpan, Hapus). Tip: buat di file repo_palsu_test.go.
//     Isi method yang tidak dipakai dengan return nilai kosong.
//  2. Tulis test:  judul kosong -> ErrValidasi   |   judul valid -> tersimpan tanpa error
//     (aturan service: judul tidak kosong, tahun > 1900, judul tidak duplikat)
//  3. Bonus: tambahkan field `dipanggil int` di repoPalsu, naikkan di Simpan,
//     lalu pastikan Simpan TIDAK dipanggil saat validasi gagal.
//  4. Jalankan:  go test -v -run TestServiceTambah
//
// Petunjuk memeriksa jenis error:
//
//	var ev ErrValidasi
//	if !errors.As(err, &ev) { t.Fatalf("harusnya ErrValidasi, dapat %v", err) }
//
// Hapus t.Skip setelah Anda mulai menulis test sungguhan.
func TestServiceTambah(t *testing.T) {
	t.Skip("TODO Latihan 3: tulis test service dengan repoPalsu")
}
