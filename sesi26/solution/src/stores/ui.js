import { defineStore } from 'pinia'

let timer

export const useUiStore = defineStore('ui', {
  state: () => ({ pesan: '' }),
  actions: {
    tampil(pesan, ms = 2500) {
      this.pesan = pesan
      clearTimeout(timer)
      timer = setTimeout(() => (this.pesan = ''), ms)
    },
  },
})
