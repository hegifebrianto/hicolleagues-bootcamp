<script setup>
import { ref } from 'vue'
import Profil from '../components/Profil.vue'
import Pengaturan from '../components/Pengaturan.vue'

const tab = ref('profil')
const pakaiKeepAlive = ref(true)
const komponen = { profil: Profil, pengaturan: Pengaturan }
</script>

<template>
  <section>
    <h2>Dynamic Components &amp; keep-alive</h2>
    <div class="baris">
      <button :class="{ aktif: tab === 'profil' }" @click="tab = 'profil'">Profil</button>
      <button :class="{ aktif: tab === 'pengaturan' }" @click="tab = 'pengaturan'">Pengaturan</button>
      <label class="cek"><input v-model="pakaiKeepAlive" type="checkbox" /> keep-alive</label>
    </div>

    <!-- TODO Latihan 2 (8 menit): bungkus <component> dengan <KeepAlive> supaya isian form tidak hilang
         saat pindah tab. Hubungkan dengan checkbox: <KeepAlive v-if="pakaiKeepAlive"> ... dan
         <component ... v-else /> untuk versi tanpa keep-alive. -->
    <component :is="komponen[tab]" />

    <p class="kecil">
      Isi form, pindah tab, lalu kembali. Dengan keep-alive isian tetap ada dan "mounted" tidak
      bertambah (hanya "activated"). Tanpa keep-alive isian hilang dan "mounted" naik.
    </p>
  </section>
</template>
