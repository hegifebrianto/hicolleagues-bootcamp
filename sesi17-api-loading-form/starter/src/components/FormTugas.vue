<script setup>
import { ref, computed } from 'vue'
import axios from 'axios'
import { useTugasStore } from '../stores/tugas'

const store = useTugasStore()
const judul = ref('')
const pesan = ref(null) // { jenis: 'sukses' | 'error', teks }
const sedangKirim = ref(false)

// LATIHAN 2: validasi klien
// TODO 5: errorJudul -> '' kalau judul.value.trim().length >= 3,
//         selain itu 'Minimal 3 karakter'  (pakai computed)
const errorJudul = computed(() => '') // <-- ganti sesuai TODO 5
const valid = computed(() => !errorJudul.value)

async function kirim() {
  if (!valid.value || sedangKirim.value) return
  sedangKirim.value = true
  pesan.value = null
  try {
    // TODO 6: kirim POST ke '/api/tugas' dengan body { judul: judul.value.trim() }
    //         (axios otomatis mengubahnya jadi JSON dan mengatur Content-Type)
    // TODO 7 (bonus): kosongkan form -> judul.value = ''
    // TODO 8: pesan.value = { jenis: 'sukses', teks: 'Tugas ditambahkan' }
    //         lalu segarkan daftar: await store.muat()
  } catch (e) {
    // TODO 9: pesan error. Server membalas { "error": "..." } -> e.response.data.error
    //         Kalau e.response tidak ada = server tidak terjangkau.
    pesan.value = { jenis: 'error', teks: 'Gagal menambah' } // <-- sempurnakan
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
