<script setup>
import StatusBackend from './components/StatusBackend.vue'
</script>

<template>
  <div class="app">
    <nav class="menu">
      <RouterLink to="/">📚 Daftar Buku</RouterLink>
      <RouterLink to="/tambah">+ Tambah Buku</RouterLink>
    </nav>
    <StatusBackend />
    <!-- :key membuat halaman dibuat ulang saat URL berubah (mis. /ubah/1 -> /ubah/2),
         supaya form selalu memuat data yang benar -->
    <RouterView :key="$route.fullPath" />
  </div>
</template>
