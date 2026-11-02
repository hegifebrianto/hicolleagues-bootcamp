import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './style.css'

const app = createApp(App)
const pinia = createPinia()
app.use(pinia) // pinia HARUS dipasang sebelum router (guard memakai store)
app.use(router)

// TODO 13 (bonus LATIHAN 1): simpan store ke localStorage.
//  - import { useTugasStore } from './stores/tugas' di atas
//  - const store = useTugasStore()
//  - saat start: baca localStorage.getItem('tugas'); kalau ada -> store.daftar = JSON.parse(...)
//  - store.$subscribe((mutation, state) => { localStorage.setItem('tugas', JSON.stringify(state.daftar)) })
// Tanpa ini, data HILANG setiap halaman di-refresh (bahan Aktivitas Riset Mandiri).

app.mount('#app')
