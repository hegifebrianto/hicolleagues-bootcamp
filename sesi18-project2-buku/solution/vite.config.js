import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// SOLUSI 1 untuk CORS (lihat slide): proxy di Vite. Browser hanya bicara ke
// localhost:5173; Vite yang meneruskan semua /api/... ke backend, jadi tidak ada
// request lintas-origin. Ganti target lewat:  API_TARGET=http://localhost:9000 npm run dev
const target = process.env.API_TARGET || 'http://localhost:8089'

export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: { '/api': target },
  },
})
