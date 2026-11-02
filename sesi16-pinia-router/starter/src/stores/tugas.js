import { defineStore } from 'pinia'

// LATIHAN 1: Store Tugas
// Stub di bawah membuat aplikasi tetap bisa jalan sebelum TODO dikerjakan.
export const useTugasStore = defineStore('tugas', {
  // TODO 1: state berupa FUNGSI yang mengembalikan objek { daftar: [...] }.
  // Isi awal boleh array kosong. Tiap tugas: { id, judul, selesai }.
  state: () => ({
    daftar: [],
  }),
  getters: {
    // TODO 2: getter `jumlah` -> jumlah seluruh tugas (daftar.length)
    jumlah: (s) => 0,
    // TODO 3: getter `selesai` -> jumlah tugas yang selesai (pakai .filter lalu .length)
    selesai: (s) => 0,
  },
  actions: {
    tambah(judul) {
      // TODO 4: trim() judul; kalau kosong -> return.
      // Kalau tidak kosong, push { id: Date.now(), judul: <hasil trim>, selesai: false }
      // ke this.daftar.
    },
    hapus(id) {
      // TODO 5: this.daftar = this.daftar.filter(...)  -> buang tugas dengan id tersebut
    },
    toggle(id) {
      // TODO 6: cari tugas dengan id itu (this.daftar.find), lalu balik t.selesai
      // (cek dulu hasil find tidak undefined).
    },
  },
})
