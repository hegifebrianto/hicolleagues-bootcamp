package main

// repoPalsu: "stub/mock" -- struct yang mengimplementasikan interface BukuRepo
// tapi TIDAK menyentuh database. Karena Service hanya mengenal interface,
// test bisa menyuntikkan repo palsu ini (Sesi 17: dependency injection).
type repoPalsu struct {
	data []Buku

	// Perekam panggilan (mock): berapa kali & dengan apa method dipanggil
	simpanDipanggil int
	tersimpan       []Buku

	// Pemicu kegagalan: isi untuk mensimulasikan error dari lapisan data
	errSemua, errSimpan, errHapus error
}

func (r *repoPalsu) Semua() ([]Buku, error) {
	if r.errSemua != nil { //jika ada error maka tampilkan error
		return nil, r.errSemua
	} //jika tidak maka tampilkan semua
	return r.data, nil
}

func (r *repoPalsu) Ambil(id int64) (Buku, error) {
	for _, b := range r.data {
		if b.ID == id {
			return b, nil
		}
	}
	// Buku = {id:1,judul:"bumi manusia",tahun : 2000}
	return Buku{}, ErrTidakDitemukan
}

func (r *repoPalsu) Simpan(b Buku) (Buku, error) {
	r.simpanDipanggil++     //dari 0 ke 1
	if r.errSimpan != nil { //apakah ada error di simpan ?
		return Buku{}, r.errSimpan
	} //jika tidak ada error maka
	b.ID = int64(len(r.data) + 1)
	// b = {id:1,judul:"bumi manusia",tahun : 2000}
	r.tersimpan = append(r.tersimpan, b)
	return b, nil
}

func (r *repoPalsu) Hapus(id int64) error {

	if r.errHapus != nil { //apakah ada error di r.errorHapus ?
		return r.errHapus //maka retrun error nya

	} //jika tidak error maka running ke sini
	for _, b := range r.data {
		if b.ID == id {
			return nil
		}
	}
	return ErrTidakDitemukan
}
