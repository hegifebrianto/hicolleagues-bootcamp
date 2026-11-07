import { createApp } from 'vue'
import { createPinia } from 'pinia'
import axios from 'axios'
import App from './App.vue'
import './style.css'

// Kosong  -> request relatif ("/api/tugas"), diteruskan oleh PROXY Vite (solusi 1 CORS).
// Terisi  -> request langsung ke backend lain origin, mis. VITE_API_BASE=http://localhost:8089 npm run dev
//            (itu butuh solusi 2: header CORS di backend -- sudah ada di backend demo).
axios.defaults.baseURL = import.meta.env.VITE_API_BASE || ''

createApp(App).use(createPinia()).mount('#app')
