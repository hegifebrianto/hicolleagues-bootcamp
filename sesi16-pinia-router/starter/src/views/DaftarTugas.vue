<script setup>
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useTugasStore } from '../stores/tugas'

const store = useTugasStore()
const { daftar, jumlah, selesai } = storeToRefs(store)
const judulBaru = ref('')

function tambah() {
  store.tambah(judulBaru.value)
  judulBaru.value = ''
}
</script>

<template>
  <section class="card">
    <h1>Daftar Tugas</h1>
    <p class="muted">{{ selesai }} dari {{ jumlah }} selesai</p>

    <form class="baris" @submit.prevent="tambah">
      <input v-model="judulBaru" placeholder="Tugas baru..." />
      <button type="submit">Tambah</button>
    </form>

    <p v-if="daftar.length === 0" class="kosong">Belum ada tugas.</p>
    <ul class="daftar">
      <li v-for="t in daftar" :key="t.id" class="item" :class="{ selesai: t.selesai }">
        <input type="checkbox" :checked="t.selesai" @change="store.toggle(t.id)" />
        <!-- RouterLink dengan param dinamis -->
        <RouterLink class="judul" :to="'/tugas/' + t.id">{{ t.judul }}</RouterLink>
        <button class="btn-hapus" type="button" @click="store.hapus(t.id)">Hapus</button>
      </li>
    </ul>
  </section>
</template>
