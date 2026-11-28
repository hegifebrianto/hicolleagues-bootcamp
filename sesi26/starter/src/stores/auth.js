import { defineStore } from 'pinia'
// import { api } from '../api/client'      // TODO: dibutuhkan untuk login
// import { useUiStore } from './ui'        // TODO (bonus): tampilkan toast setelah login

const KUNCI = 'sesi26-auth'

// TODO Latihan 3 (9 menit): Store auth module
//  - getters: sudahLogin  -> true bila token ada
//  - actions.login(email, password):
//      const { data } = await api.post('/api/login', { email, password })
//      simpan data.data.token dan data.data.user ke state
//      Bonus: simpan juga ke localStorage (KUNCI) dan baca lagi di state()
//  - actions.logout(): kosongkan token & user (dan hapus localStorage)
export const useAuthStore = defineStore('auth', {
  state: () => ({ user: null, token: null }),

  getters: {
    sudahLogin: () => false, // TODO
  },

  actions: {
    async login(email, password) {
      // TODO
    },
    logout() {
      // TODO
    },
  },
})
