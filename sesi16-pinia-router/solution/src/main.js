import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import { useTugasStore } from './stores/tugas'
import './style.css'

// Ubah ke false untuk MEREPRODUKSI masalah "state Pinia hilang saat refresh"
// -- itu bahan Aktivitas Riset Mandiri (soft skill RCA) di slide.
const SIMPAN_KE_LOCALSTORAGE = true

const app = createApp(App)
const pinia = createPinia()
app.use(pinia) // pinia HARUS dipasang sebelum router (guard memakai store)
app.use(router)

if (SIMPAN_KE_LOCALSTORAGE) {
  const store = useTugasStore()
  try {
    const tersimpan = localStorage.getItem('tugas')
    if (tersimpan) store.daftar = JSON.parse(tersimpan)
  } catch (e) {
    localStorage.removeItem('tugas') // data rusak -> buang
  }
  // $subscribe: dipanggil setiap state store berubah
  store.$subscribe((mutation, state) => {
    localStorage.setItem('tugas', JSON.stringify(state.daftar))
  })
}

app.mount('#app')
