import { defineStore } from 'pinia'
import { api } from '../api/client'
import { useUiStore } from './ui'

const KUNCI = 'sesi26-auth'

function bacaTersimpan() {
  try {
    return JSON.parse(localStorage.getItem(KUNCI)) || {}
  } catch {
    return {}
  }
}

export const useAuthStore = defineStore('auth', {
  state: () => {
    const t = bacaTersimpan()
    return { user: t.user || null, token: t.token || null }
  },

  getters: {
    sudahLogin: (state) => !!state.token,
  },

  actions: {
    async login(email, password) {
      const { data } = await api.post('/api/login', { email, password })
      this.token = data.data.token
      this.user = data.data.user
      // Bonus slide: simpan ke localStorage agar bertahan saat refresh
      localStorage.setItem(KUNCI, JSON.stringify({ token: this.token, user: this.user }))
      // Store boleh memakai store lain
      useUiStore().tampil(`Halo, ${this.user.nama}! Login berhasil.`)
    },

    logout() {
      this.token = null
      this.user = null
      localStorage.removeItem(KUNCI)
    },
  },
})
