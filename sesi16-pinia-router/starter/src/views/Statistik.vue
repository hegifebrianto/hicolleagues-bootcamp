<script setup>
import { storeToRefs } from 'pinia'
import { useTugasStore } from '../stores/tugas'

const { jumlah, selesai } = storeToRefs(useTugasStore())
</script>

<template>
  <div>
    <h2>Statistik</h2>
    <p>Total tugas: <strong>{{ jumlah }}</strong></p>
    <p>Selesai: <strong>{{ selesai }}</strong></p>
    <p class="muted">Halaman anak ini membaca store yang sama dengan halaman lain.</p>
  </div>
</template>
