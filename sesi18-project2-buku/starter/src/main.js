import { createApp } from 'vue'
import { createPinia } from 'pinia'
import axios from 'axios'
import App from './App.vue'
import router from './router'
import './style.css'

// Kosong  -> request relatif ("/api/buku"), diteruskan oleh PROXY Vite (lihat vite.config.js).
// Terisi  -> langsung ke backend lain origin, mis. VITE_API_BASE=http://localhost:8089 npm run dev
axios.defaults.baseURL = import.meta.env.VITE_API_BASE || ''

createApp(App).use(createPinia()).use(router).mount('#app')
