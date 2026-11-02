<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useTugasStore } from '../stores/tugas'

const route = useRoute()
const store = useTugasStore()

// TODO 12 (LATIHAN 3): cari tugas yang id-nya sama dengan route.params.id.
// AWAS: route.params.id berupa STRING, id di store berupa ANGKA --
// ubah dulu dengan Number(...) sebelum dibandingkan (=== ).
// (Kalau lupa, halaman selalu "Tidak ditemukan" -- itulah bug di contoh "5 Why" slide RCA.)
const tugas = computed(() => undefined) // <-- ganti sesuai TODO 12
</script>

<template>
  <section class="card">
    <template v-if="tugas">
      <h1>{{ tugas.judul }}</h1>
      <p class="muted">ID: {{ tugas.id }} · Status: {{ tugas.selesai ? 'Selesai' : 'Belum selesai' }}</p>
      <button type="button" @click="store.toggle(tugas.id)">
        Tandai {{ tugas.selesai ? 'belum selesai' : 'selesai' }}
      </button>
    </template>
    <template v-else>
      <h1>Tidak ditemukan</h1>
      <p class="muted">Tugas dengan id "{{ route.params.id }}" tidak ada.</p>
    </template>
    <p><RouterLink to="/tugas">← Kembali ke daftar</RouterLink></p>
  </section>
</template>
