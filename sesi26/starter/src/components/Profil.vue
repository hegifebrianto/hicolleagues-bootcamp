<script setup>
import { ref, onMounted, onActivated } from 'vue'
import { hitung } from '../utils/pantau'

const nama = ref('')
onMounted(() => hitung.Profil.mounted++)
onActivated(() => hitung.Profil.activated++)
</script>

<template>
  <div class="tab-isi">
    <h3>Profil</h3>
    <label>Nama <input v-model="nama" placeholder="ketik nama..." /></label>
    <p class="kecil">mounted: {{ hitung.Profil.mounted }}× · activated: {{ hitung.Profil.activated }}×</p>
  </div>
</template>
