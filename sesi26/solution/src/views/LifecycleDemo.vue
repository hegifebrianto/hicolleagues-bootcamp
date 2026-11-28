<script setup>
import { ref } from 'vue'
import Jam from '../components/Jam.vue'
import { pantau, lepasSemua } from '../utils/pantau'

const tampil = ref(true)
const bersih = ref(true)
const log = ref([])

function catat(pesan) {
  log.value.unshift(pesan)
  log.value = log.value.slice(0, 12)
}
</script>

<template>
  <section>
    <h2>Lifecycle Hooks &amp; Membersihkan Efek Samping</h2>
    <div class="baris">
      <button @click="tampil = !tampil">{{ tampil ? 'Sembunyikan' : 'Tampilkan' }} komponen Jam</button>
      <label class="cek"><input v-model="bersih" type="checkbox" /> Bersihkan timer saat unmount</label>
    </div>

    <Jam v-if="tampil" :bersih="bersih" @log="catat" />

    <p class="status" :class="{ bahaya: pantau.aktif > (tampil ? 1 : 0) }">
      Timer yang masih hidup: <b>{{ pantau.aktif }}</b>
      <button class="kecil-btn" @click="lepasSemua">Matikan semua</button>
    </p>
    <p class="kecil">
      Coba: hilangkan centang, sembunyikan Jam, tampilkan lagi — ulangi. Angka timer terus naik
      (memory leak). Dengan centang, angka kembali ke 0 saat Jam disembunyikan.
    </p>

    <h3>Log hook</h3>
    <ol class="log"><li v-for="(l, i) in log" :key="i">{{ l }}</li></ol>
  </section>
</template>
