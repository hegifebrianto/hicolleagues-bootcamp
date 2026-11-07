<script setup>
import { ref, computed } from 'vue'
import axios from 'axios'
import { useTugasStore } from '../stores/tugas'

const store = useTugasStore()
const judul = ref('')
const pesan = ref(null) // { jenis: 'sukses' | 'error', teks }
const sedangKirim = ref(false)

// LATIHAN 2: validasi klien -- judul minimal 3 karakter
const errorJudul = computed(() => (judul.value.trim().length >= 3 ? '' : 'Minimal 3 karakter'))
const valid = computed(() => !errorJudul.value)

async function kirim() {
  if (!valid.value || sedangKirim.value) return
  sedangKirim.value = true
  pesan.value = null
  try {
    // Axios otomatis mengubah objek jadi JSON + mengatur Content-Type
    await axios.post('/api/tugas', { judul: judul.value.trim() })
    judul.value = '' // bonus: kosongkan form setelah berhasil
    pesan.value = { jenis: 'sukses', teks: 'Tugas ditambahkan' }
    await store.muat() // segarkan daftar
  } catch (e) {
    const dariServer = e.response && e.response.data && e.response.data.error
    pesan.value = {
      jenis: 'error',
      teks: dariServer || (e.response ? 'Gagal menambah' : 'Server tidak dapat dihubungi'),
    }
  } finally {
    sedangKirim.value = false
  }
}
</script>

<template>
  <section class="card">
    <h2>Tambah Tugas</h2>
    <form class="baris" @submit.prevent="kirim">
      <div class="kolom">
        <input v-model="judul" placeholder="Judul tugas (min. 3 karakter)" />
        <small v-if="judul && errorJudul" class="error-text">{{ errorJudul }}</small>
      </div>
      <button type="submit" :disabled="!valid || sedangKirim">
        {{ sedangKirim ? 'Mengirim...' : 'Kirim' }}
      </button>
    </form>
    <p v-if="pesan" :class="['pesan', pesan.jenis]">{{ pesan.teks }}</p>
    <p class="muted">Validasi klien membantu pengguna, tetapi tidak menggantikan validasi server.</p>
  </section>
</template>
