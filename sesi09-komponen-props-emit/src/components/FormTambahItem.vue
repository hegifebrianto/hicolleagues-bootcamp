<script setup>
import { ref } from 'vue'

/**
 * Komponen anak #1: FormTambahItem
 *
 * Komponen ini TIDAK menyimpan daftar belanja -- dia hanya tahu
 * cara menampilkan form dan "memberi tahu" parent-nya (App.vue)
 * saat pengguna ingin menambah item baru. Itulah fungsi emit:
 * child "memancarkan" event ke atas, parent yang memutuskan apa
 * yang terjadi selanjutnya.
 */

// defineEmits mendaftarkan event apa saja yang BOLEH dipancarkan
// komponen ini -- di sini cuma satu: 'tambah'
const emit = defineEmits(['tambah'])

const nama = ref('')
const harga = ref(0)

function kirim() {
  if (!nama.value.trim()) return
  // emit('tambah', data) -- mengirim data ke parent lewat event 'tambah'
  emit('tambah', { nama: nama.value, harga: Number(harga.value) || 0 })
  nama.value = ''
  harga.value = 0
}
</script>

<template>
  <div class="tambah">
    <input v-model="nama" type="text" placeholder="Nama item..." @keyup.enter="kirim" />
    <input v-model.number="harga" type="number" placeholder="Harga" style="max-width: 120px" />
    <button @click="kirim">Tambah</button>
  </div>
</template>
