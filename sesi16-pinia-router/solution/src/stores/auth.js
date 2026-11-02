import { defineStore } from 'pinia'

// Status login sebagai state global (salah satu contoh "butuh store global" di slide).
// Slide memakai localStorage.getItem('token') langsung di guard; di sini dibungkus
// store supaya tampilan (tombol Masuk/Keluar) ikut reaktif.
export const useAuthStore = defineStore('auth', {
  state: () => ({ token: localStorage.getItem('token') }),
  getters: {
    masuk: (s) => !!s.token,
  },
  actions: {
    login() {
      this.token = 'token-demo'
      localStorage.setItem('token', this.token)
    },
    keluar() {
      this.token = null
      localStorage.removeItem('token')
    },
  },
})
