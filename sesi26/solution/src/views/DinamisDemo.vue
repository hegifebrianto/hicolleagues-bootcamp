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

    <KeepAlive v-if="pakaiKeepAlive">
      <component :is="komponen[tab]" />
    </KeepAlive>
    <component :is="komponen[tab]" v-else />

    <p class="kecil">
      Isi form, pindah tab, lalu kembali. Dengan keep-alive isian tetap ada dan "mounted" tidak
      bertambah (hanya "activated"). Tanpa keep-alive isian hilang dan "mounted" naik.
    </p>
  </section>
</template>
