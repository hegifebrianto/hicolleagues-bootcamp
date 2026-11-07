<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useBukuStore } from '../stores/buku'

const route = useRoute()
const store = useBukuStore()
const buku = ref(null)
const memuat = ref(true)
const galat = ref('')

onMounted(async () => {
  try {
    buku.value = await store.ambil(route.params.id)
  } catch (e) {
    galat.value = !e.response
      ? 'Server tidak dapat dihubungi.'
      : e.response.status === 404
        ? 'Buku tidak ditemukan.'
        : 'Gagal memuat buku.'
  } finally {
    memuat.value = false
  }
})
</script>

<template>
  <section class="card">
    <p v-if="memuat" class="pesan info">Memuat...</p>
    <p v-else-if="galat" class="pesan error">{{ galat }}</p>
    <template v-else>
      <h1>{{ buku.judul }}</h1>
      <p>Penulis: <strong>{{ buku.penulis }}</strong></p>
      <p>Tahun terbit: <strong>{{ buku.tahun }}</strong></p>
      <div class="aksi-form">
        <RouterLink class="btn" :to="'/ubah/' + buku.id">Ubah</RouterLink>
      </div>
    </template>
    <p><RouterLink to="/">← Kembali ke daftar</RouterLink></p>
  </section>
</template>
