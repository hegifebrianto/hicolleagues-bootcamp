<script setup>
import { storeToRefs } from 'pinia'
import { useTugasStore } from '../stores/tugas'

const store = useTugasStore()
// storeToRefs: ambil state/getter sebagai ref TANPA kehilangan reaktivitas


const { daftar, jumlah, selesai } = storeToRefs(store)

console.log('daftar tugas:',daftar.value ,'jumlah :',jumlah.value,'selesai :',selesai.value);
</script>

<template>
  <section class="card">
    <h1>Selamat datang 👋</h1>
    <p class="muted">Demo Sesi 16: Pinia (state global) dan Vue Router (SPA).</p>
    <p>Jumlah tugas: <strong>{{ jumlah }}</strong> · selesai: <strong>{{ selesai }}</strong></p>
    <p class="muted">Data yang sama dipakai di menu navigasi, halaman ini, dan halaman Tugas — itulah gunanya store global.</p>
    <RouterLink class="btn" to="/tugas">Buka daftar tugas</RouterLink>
  </section>
</template>
