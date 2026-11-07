<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useTugasStore } from '../stores/tugas'

const store = useTugasStore()
const pesan = ref(null)

// memuat data saat komponen pertama kali tampil
onMounted(() => store.muat())

function tampilkan(jenis, teks) {
  pesan.value = { jenis, teks }
}

// LATIHAN 3: pesan sesuai status code dari server
function pesanGagal(e) {
  if (!e.response) return 'Server tidak dapat dihubungi. Pastikan backend menyala.'
  if (e.response.status === 404) return 'Tugas tidak ditemukan'
  return 'Terjadi kesalahan'
}

async function hapus(t) {
  console.log(t,'element dari row');
  if (!confirm('Hapus tugas "' + t.judul + '"?')) return // bonus: konfirmasi
  try {
    await axios.delete('/api/tugas/' + t.id) // sukses = 204
    tampilkan('sukses', 'Tugas dihapus')
  } catch (e) {
    tampilkan('error', pesanGagal(e))
  }
  await store.muat() // selalu segarkan, supaya daftar cocok dengan server (termasuk saat 404)
}

// PUT: ubah status selesai
async function toggle(t) {
  try {
    await axios.put('/api/tugas/' + t.id, { judul: t.judul, selesai: !t.selesai })
  } catch (e) {
    tampilkan('error', pesanGagal(e))
  }
  await store.muat()
}
</script>

<template>
  <section class="card">
    <div class="bar">
      <h2>Daftar Tugas</h2>
      <button class="btn-sekunder" type="button" :disabled="store.loading" @click="store.muat()">
        Muat ulang
      </button>
    </div>

    <p v-if="pesan" :class="['pesan', pesan.jenis]">{{ pesan.teks }}</p>

    <!-- LATIHAN 1: tiga kondisi tampilan -->
    <p v-if="store.loading" class="pesan info">Memuat...</p>
    <p v-else-if="store.error" class="pesan error">{{ store.error }}</p>
    <p v-else-if="store.daftar.length === 0" class="kosong">Belum ada tugas.</p>
    <ul v-else class="daftar">
      <li v-for="t in store.daftar" :key="t.id" class="item" :class="{ selesai: t.selesai }">
        <input type="checkbox" :checked="t.selesai" @change="toggle(t)" />
        <span class="judul">{{ t.judul }}</span>
        <button class="btn-hapus" type="button" @click="hapus(t)">Hapus</button>
      </li>
    </ul>
  </section>
</template>
