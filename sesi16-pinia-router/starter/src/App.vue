<script setup>
import { RouterLink, RouterView, useRouter } from 'vue-router'
import { useAuthStore } from './stores/auth'
import { useTugasStore } from './stores/tugas'

const auth = useAuthStore()
const tugas = useTugasStore()
const router = useRouter()

function keluar() {
  auth.keluar()
  router.push('/') // navigasi programatik
}
</script>

<template>
  <div class="app">
    <nav class="menu">
      <RouterLink to="/">Beranda</RouterLink>
      <RouterLink to="/tugas">Tugas <span class="badge">{{ tugas.jumlah }}</span></RouterLink>
      <RouterLink to="/tentang">Tentang</RouterLink>
      <RouterLink to="/admin">Admin</RouterLink>
      <span class="spacer"></span>
      <button v-if="auth.masuk" class="btn-kecil" type="button" @click="keluar">Keluar</button>
      <RouterLink v-else to="/login">Masuk</RouterLink>
    </nav>

    <!-- RouterView = tempat halaman sesuai URL ditampilkan (SPA: tanpa reload) -->
    <RouterView />
  </div>
</template>
