<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

// Indikator kecil: apakah frontend bisa menjangkau backend? (klik untuk cek ulang)
const statusConnected = ref('cek') // 'cek' | 'ok' | 'mati'

async function cek() {
  statusConnected.value = 'cek'
  try {
    const dataConnected = await axios.get('/api/ping', { timeout: 3000 }) //satuan milisecond
    const {status,data} = dataConnected
    console.log(status,data.data);

    if(status === 200 & data.data === 'pong'){
      statusConnected.value = 'ok'
    }
    
  } catch (e) {
    statusConnected.value = 'mati'
  }
}

onMounted(cek) //pertama kali di load maka dia akan menjalankan function cek
</script>

<template>
  <button class="statusConnected" :class="statusConnected" type="button" title="Klik untuk cek ulang" @click="cek">
    <span class="titik"></span>
    <span v-if="statusConnected === 'cek'">Mengecek backend...</span>
    <span v-else-if="statusConnected === 'ok'">Backend terhubung</span>
    <span v-else>Backend TIDAK terhubung — jalankan server (lihat README)</span>
  </button>
  <!-- <div style="margin-top: 20px;"></div> -->
</template>
