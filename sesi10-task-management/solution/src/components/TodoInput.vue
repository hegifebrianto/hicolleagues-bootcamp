<script setup>
import { ref } from 'vue'

/**
 * TodoInput -- form untuk menambah tugas baru. Komponen ini TIDAK
 * menyimpan daftar tugas sendiri, cuma memancarkan event 'tambah'
 * membawa teksnya; App.vue (parent) yang memutuskan apa yang terjadi
 * selanjutnya (termasuk validasi input kosong).
 */
const emit = defineEmits(['tambah'])

const teks = ref('')

function kirim() {
  console.log(teks.value);
  emit('tambah', teks.value)
  resetForm();
}


function resetForm() {
    teks.value = '' // kosongkan input setelah dikirim

}
</script>

<template>
  <div class="tambah">
    <input v-model="teks" placeholder="Tugas baru..." @keyup.enter="kirim" />
    <button @click="kirim">Tambah</button>
  </div>
</template>
