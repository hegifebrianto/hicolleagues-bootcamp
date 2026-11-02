<script setup>
import { useRoute } from 'vue-router'
const route = useRoute()
</script>

<template>
  <section class="card">
    <h1>404 — Halaman tidak ditemukan</h1>
    <p class="muted">Tidak ada halaman untuk <code>{{ route.fullPath }}</code>.</p>
    <RouterLink class="btn" to="/">Ke beranda</RouterLink>
  </section>
</template>
