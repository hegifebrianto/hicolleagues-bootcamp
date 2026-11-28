import { defineStore } from 'pinia'
import { api } from '../api/client'
import { useAuthStore } from './auth'

export const useTugasStore = defineStore('tugas', {
  state: () => ({ daftar: [], loading: false, error: '' }),

  actions: {
    async muat() {
      // Store tugas MEMAKAI store auth (slide "Pinia Modules")
      const auth = useAuthStore()
      console.log('memastikan auth nya ada :',auth.token);
      this.loading = true
      this.error = ''
      try {
        const { data } = await api.get('/api/tugas', {
          headers: { Authorization: 'Bearer ' + auth.token },
        })

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
