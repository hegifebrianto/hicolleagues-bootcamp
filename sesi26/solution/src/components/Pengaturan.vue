<script setup>
import { ref, onMounted, onActivated } from 'vue'
import { hitung } from '../utils/pantau'

const email = ref('')
const notif = ref(false)
onMounted(() => hitung.Pengaturan.mounted++)
onActivated(() => hitung.Pengaturan.activated++)
</script>

<template>
  <div class="tab-isi">
    <h3>Pengaturan</h3>
    <label>Email <input v-model="email" placeholder="ketik email..." /></label>
    <label class="cek"><input v-model="notif" type="checkbox" /> Kirim notifikasi</label>
    <p class="kecil">mounted: {{ hitung.Pengaturan.mounted }}× · activated: {{ hitung.Pengaturan.activated }}×</p>
  </div>
</template>
