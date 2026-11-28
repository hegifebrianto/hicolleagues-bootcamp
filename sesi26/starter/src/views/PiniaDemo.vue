<script setup>
import { ref } from 'vue'
import { useAuthStore } from '../stores/auth'
import { useTugasStore } from '../stores/tugas'
import Daftar from '../components/Daftar.vue'

const auth = useAuthStore()
const tugas = useTugasStore()

const email = ref('demo@mail.com')
const password = ref('')
const pesan = ref('')

async function masuk() {
  pesan.value = ''
  try {
    await auth.login(email.value, password.value)
  } catch {
    pesan.value = 'Email atau password salah'
  }
}
</script>

<template>
  <section>
    <h2>Pinia Modules: auth · tugas · ui</h2>

    <div v-if="!auth.sudahLogin" class="form">
      <label>Email <input v-model="email" /></label>
      <label>Password <input v-model="password" type="password" placeholder="rahasia" /></label>
      <button @click="masuk">Login</button>
      <p v-if="pesan" class="bahaya">{{ pesan }}</p>
      <p class="kecil">Akun demo: <code>demo@mail.com</code> / <code>rahasia</code></p>
    </div>
    <div v-else class="baris">
      <span>Login sebagai <b>{{ auth.user.nama }}</b></span>
      <button @click="auth.logout()">Logout</button>
    </div>

    <h3>Store tugas (memakai token dari store auth)</h3>
    <button :disabled="tugas.loading" @click="tugas.muat()">
      {{ tugas.loading ? 'Memuat...' : 'Muat tugas' }}
    </button>
    <p v-if="tugas.error" class="bahaya">{{ tugas.error }}</p>
    <Daftar :items="tugas.daftar">
      <template #default="{ item }">
        {{ item.judul }} <b v-if="item.selesai" class="ok">Selesai</b>
      </template>
    </Daftar>
    <p class="kecil">
      Coba "Muat tugas" sebelum login → 401. Setelah login → daftar muncul. Refresh halaman: token tetap ada
      (localStorage).
    </p>
  </section>
</template>
