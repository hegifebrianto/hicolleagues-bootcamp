<script setup>
import { ref } from 'vue'

/**
 * TodoInput -- form untuk menambah tugas baru. Komponen ini TIDAK
 * menyimpan daftar tugas sendiri, cuma memancarkan event 'tambah'
 * membawa teksnya; App.vue (parent) yang memutuskan apa yang terjadi
 * selanjutnya (termasuk validasi input kosong).
 */

// TODO 1: daftarkan event yang BOLEH dipancarkan komponen ini: 'tambah'
const emit = defineEmits([])

const teks = ref('')

function kirim() {
  // TODO 2: pancarkan event 'tambah' membawa isi teks.value
  // (emit('tambah', teks.value))

  // TODO 3: kosongkan teks.value setelah dikirim
}
</script>

<template>
  <div class="tambah">
    <input v-model="teks" placeholder="Tugas baru..." @keyup.enter="kirim" />
    <button @click="kirim">Tambah</button>
  </div>
</template>
