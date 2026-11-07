<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

// Indikator kecil: apakah frontend bisa menjangkau backend? (klik untuk cek ulang)
const status = ref('cek') // 'cek' | 'ok' | 'mati'

async function cek() {
  status.value = 'cek'
  try {
    await axios.get('/api/ping', { timeout: 3000 })
    status.value = 'ok'
  } catch (e) {
    status.value = 'mati'
  }
}

onMounted(cek)
</script>

<template>
  <button class="status" :class="status" type="button" title="Klik untuk cek ulang" @click="cek">
    <span class="titik"></span>
    <span v-if="status === 'cek'">Mengecek backend...</span>
    <span v-else-if="status === 'ok'">Backend terhubung</span>
    <span v-else>Backend TIDAK terhubung — jalankan server (lihat README)</span>
  </button>
</template>
