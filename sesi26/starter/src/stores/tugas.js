import { defineStore } from 'pinia'
import { api } from '../api/client'
import { useAuthStore } from './auth'

export const useTugasStore = defineStore('tugas', {
  state: () => ({ daftar: [], loading: false, error: '' }),

  actions: {
    async muat() {
      // Store tugas MEMAKAI store auth (slide "Pinia Modules")
      const auth = useAuthStore()
      this.loading = true
      this.error = ''
      try {
        // TODO Latihan 3: kirim header Authorization: 'Bearer ' + auth.token
        // (tanpa header server palsu membalas 401)
        const { data } = await api.get('/api/tugas')
        this.daftar = data.data
      } catch (e) {
        this.daftar = []
        if (e.response?.status === 401) {
          this.error = '401: belum login atau token tidak valid'
          auth.logout()
        } else {
          this.error = 'Gagal memuat tugas'
        }
      } finally {
        this.loading = false
      }
    },
  },
})
