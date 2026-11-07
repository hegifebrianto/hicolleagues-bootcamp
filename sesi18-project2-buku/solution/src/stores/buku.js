import axios from 'axios'
import { defineStore } from 'pinia'

// Semua panggilan API ada di store; view tetap ringkas (lihat slide "Store dengan CRUD").
export const useBukuStore = defineStore('buku', {
  state: () => ({ daftar: [], loading: false, error: '' }),
  actions: {
    async muat() {
      this.loading = true
      this.error = ''
      try {
        const { data } = await axios.get('/api/buku')
        // console.log(data,'data buku');

        this.daftar = data.data
        console.log(this.daftar,'isi dari this.daftar');
      } catch (e) {
        // Lewat proxy Vite, backend yang mati muncul sebagai status 5xx (bukan error jaringan)
        this.error = !e.response
          ? 'Server tidak dapat dihubungi. Pastikan backend menyala.'
          : e.response.status >= 500
            ? 'Server bermasalah atau tidak dapat dihubungi (status ' + e.response.status + '). Pastikan backend menyala.'
            : 'Gagal memuat buku (status ' + e.response.status + ')'
      } finally {
        this.loading = false
      }
    },
    // Mengambil SATU buku (untuk form ubah & halaman detail). Error sengaja dilempar
    // ke pemanggil supaya view yang memutuskan pesannya.
    async ambil(id) {
      const { data } = await axios.get('/api/buku/' + id)
      return data.data
    },
    // tambah / ubah / hapus: kirim ke API, lalu muat ulang daftar supaya tampilan selalu sinkron
    async tambah(b) {
      await axios.post('/api/buku', b)
      await this.muat()
    },
    async ubah(id, b) {
      await axios.put('/api/buku/' + id, b)
      await this.muat()
    },
    async hapus(id) {
      await axios.delete('/api/buku/' + id)
      await this.muat()
    },
  },
})
