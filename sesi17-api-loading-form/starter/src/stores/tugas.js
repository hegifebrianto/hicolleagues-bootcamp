import axios from 'axios'
import { defineStore } from 'pinia'

export const useTugasStore = defineStore('tugas', {
  state: () => ({ daftar: [], loading: false, error: null }),
  actions: {
    // LATIHAN 1: GET + loading + error
    async muat() {
      // TODO 1: set this.loading = true dan this.error = null.
      // TODO 2: di dalam try { ... } ambil data:
      //     const { data } = await axios.get('/api/tugas')
      //     this.daftar = data.data        // format Respons dari Go: { "data": [...] }
      // TODO 3: di dalam catch (e) isi this.error dengan pesan yang jelas, mis.:
      //     !e.response            -> 'Server tidak dapat dihubungi. Pastikan backend menyala.'
      //     e.response.status >= 500 -> 'Server bermasalah atau tidak dapat dihubungi (status 5xx)...'
      //                                (lewat proxy Vite, backend yang MATI muncul sebagai status 5xx)
      //     selain itu              -> 'Gagal memuat tugas (status ' + e.response.status + ')'
      // TODO 4: di dalam finally { ... } set this.loading = false
      //     (finally menjamin loading SELALU dimatikan, sukses maupun gagal).
      //
      // Bentuk akhirnya:  try { ... } catch (e) { ... } finally { ... }
    },
  },
})
