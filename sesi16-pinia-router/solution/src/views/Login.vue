<script setup>
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const route = useRoute()
const router = useRouter()
const auth = useAuthStore()

function masuk() {
  auth.login()
  // kembali ke halaman yang tadi dituju (disimpan guard di query.redirect)
  router.push(route.query.redirect || '/tugas')
}
</script>

<template>
  <section class="card">
    <h1>Masuk</h1>
    <p v-if="route.query.redirect" class="pesan info">
      Halaman <code>{{ route.query.redirect }}</code> butuh login. Guard mengarahkanmu ke sini.
    </p>
    <p class="muted">Simulasi: tombol di bawah menyimpan sebuah "token" palsu, tanpa password.</p>
    <button type="button" @click="masuk">Masuk (simulasi)</button>
  </section>
</template>
