import { defineStore } from 'pinia'

// LATIHAN 1: Store Tugas -- state daftar, getter jumlah/selesai, action tambah/hapus/toggle
export const useTugasStore = defineStore('tugas', {
  // state = FUNGSI yang mengembalikan objek (supaya tiap app punya salinan sendiri)
  state: () => ({
    daftar: [
      // { id: 1, judul: 'Belajar Pinia', selesai: false },
      // { id: 2, judul: 'Belajar Vue Router', selesai: true },
    ],
  }),
  getters: {
    jumlah: (s) => s.daftar.length,
    selesai: (s) => s.daftar.filter((t) => t.selesai).length,
  },
  actions: {
    // di dalam actions, `this` merujuk ke store
    tambah(judul) {
      const judulBersih = judul.trim()
      if (!judulBersih) return

      this.daftar.push({ id: Date.now(), judul: judulBersih, selesai: false })
    },
    hapus(id) {
      this.daftar = this.daftar.filter((t) => t.id !== id)
    },
    toggle(id) {
      const t = this.daftar.find((t) => t.id === id)
      if (t) t.selesai = !(t.selesai)
    },
  },
})
