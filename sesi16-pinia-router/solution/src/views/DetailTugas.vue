<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useTugasStore } from '../stores/tugas'

const route = useRoute()
const store = useTugasStore()

// route.params.id SELALU string, id di store berupa angka.
// Tanpa Number(...) find() memakai === tidak akan pernah cocok -- ini bug
// yang dibedah di contoh "5 Why" slide RCA.
const tugas = computed(() => store.daftar.find((t) => t.id === Number(route.params.id)))
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
