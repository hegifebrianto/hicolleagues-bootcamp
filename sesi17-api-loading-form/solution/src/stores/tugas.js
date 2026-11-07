import axios from 'axios'
import { defineStore } from 'pinia'

// Semua pemuatan data lewat SATU tempat ini. Tiga kondisi tampilan:
// loading, error, atau data -- dipakai komponen lewat v-if / v-else-if / v-else.
export const useTugasStore = defineStore('tugas', {
  state: () => ({ daftar: [], loading: false, error: null }),
  actions: {
    // LATIHAN 1: GET + loading + error
    async muat() {
      this.loading = true
      this.error = null
      try {
        const { data } = await axios.get('/api/tugas')
        console.log(data,'data')
        this.daftar = data.data // format Respons dari Go: { "data": [...] }
      } catch (e) {
        // Lewat proxy Vite, backend yang mati muncul sebagai status 5xx (bukan error jaringan)
        this.error = !e.response
          ? 'Server tidak dapat dihubungi. Pastikan backend menyala.'
          : e.response.status >= 500
            ? 'Server bermasalah atau tidak dapat dihubungi (status ' + e.response.status + '). Pastikan backend menyala.'
            : 'Gagal memuat tugas (status ' + e.response.status + ')'
      } finally {
        this.loading = false // finally menjamin loading SELALU dimatikan
      }
    },
  },
})
