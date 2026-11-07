import axios from 'axios'
import { defineStore } from 'pinia'

// Semua panggilan API ada di store; view tetap ringkas.
// Stub di bawah membuat aplikasi tetap bisa jalan sebelum TODO dikerjakan.
export const useBukuStore = defineStore('buku', {
  state: () => ({ daftar: [], loading: false, error: '' }),
  actions: {
    async muat() {
      // TODO 1: loading = true, error = ''.
      // TODO 2: try { const { data } = await axios.get('/api/buku'); this.daftar = data.data }
      // TODO 3: catch (e) -> isi this.error:
      //     !e.response            -> 'Server tidak dapat dihubungi. Pastikan backend menyala.'
      //     e.response.status >= 500 -> 'Server bermasalah atau tidak dapat dihubungi (status 5xx)...'
      //                                (lewat proxy Vite, backend yang MATI muncul sebagai status 5xx)
      //     selain itu              -> 'Gagal memuat buku (status ' + e.response.status + ')'
      // TODO 4: finally { this.loading = false }
    },
    async ambil(id) {
      // TODO 5: GET satu buku, kembalikan isinya:
      //     const { data } = await axios.get('/api/buku/' + id)
      //     return data.data
      // (JANGAN try/catch di sini -- biarkan error naik ke view yang memutuskan pesannya)
      return {} // <-- ganti sesuai TODO 5
    },
    async tambah(b) {
      // TODO 6: await axios.post('/api/buku', b)  lalu  await this.muat()
    },
    async ubah(id, b) {
      // TODO 7: await axios.put('/api/buku/' + id, b)  lalu  await this.muat()
    },
    async hapus(id) {
      // TODO 8: await axios.delete('/api/buku/' + id)  lalu  await this.muat()
    },
  },
})
