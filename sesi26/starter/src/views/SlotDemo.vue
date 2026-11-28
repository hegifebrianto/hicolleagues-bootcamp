<script setup>
import { ref } from 'vue'
import Kartu from '../components/Kartu.vue'
import Modal from '../components/Modal.vue'
import Daftar from '../components/Daftar.vue'

const buka = ref(false)
const tugas = [
  { id: 1, judul: 'Belajar slot', selesai: true },
  { id: 2, judul: 'Belajar scoped slot', selesai: false },
]
</script>

<template>
  <section>
    <h2>Slots: default, named, scoped</h2>

    <h3>Latihan 1 · Kartu (default + named header/footer)</h3>
    <div class="grid2">
      <Kartu>
        <template #header>Profil</template>
        <p>Nama: Budi · Role: user</p>
      </Kartu>
      <Kartu>
        <template #header>Statistik</template>
        <p>12 tugas selesai</p>
        <template #footer>Diperbarui hari ini</template>
      </Kartu>
    </div>
    <Kartu><p>Kartu tanpa header/footer: tidak ada elemen kosong dirender.</p></Kartu>

    <h3>Named slot + scoped slot · Modal</h3>
    <button @click="buka = true">Buka modal</button>
    <Modal :buka="buka" @tutup="buka = false">
      <template #header><b>Judul Modal</b></template>
      <p>Isi modal (slot default).</p>
      <template #footer="{ tutup }">
        <button @click="tutup">Selesai (tutup dari parent)</button>
      </template>
    </Modal>

    <h3>Scoped slot · Daftar</h3>
    <Daftar :items="tugas">
      <template #default="{ item, index }">
        {{ index + 1 }}. {{ item.judul }}
        <b v-if="item.selesai" class="ok">Selesai</b>
      </template>
    </Daftar>
  </section>
</template>
