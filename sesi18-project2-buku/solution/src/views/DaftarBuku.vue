<script setup>
import { ref, onMounted } from 'vue'
import { useBukuStore } from '../stores/buku'
import KartuBuku from '../components/KartuBuku.vue'

const store = useBukuStore()
const pesan = ref(null) // { jenis: 'sukses' | 'error', teks }

onMounted(() => store.muat())

async function hapus(b) {
  if (!confirm('Hapus buku "' + b.judul + '"?')) return // konfirmasi sebelum hapus
  try {
    await store.hapus(b.id)
    pesan.value = { jenis: 'sukses', teks: 'Buku "' + b.judul + '" dihapus' }
  } catch (e) {
    if (!e.response) {
      pesan.value = { jenis: 'error', teks: 'Server tidak dapat dihubungi. Buku belum dihapus.' }
    } else if (e.response.status === 404) {
      pesan.value = { jenis: 'error', teks: 'Buku tidak ditemukan (mungkin sudah dihapus).' }
      await store.muat() // sinkronkan daftar dengan server
    } else {
      pesan.value = { jenis: 'error', teks: 'Terjadi kesalahan saat menghapus.' }
    }
  }
}
</script>

<template>
  <section>
    <div class="bar">
      <h1>Daftar Buku</h1>
      <button class="btn-sekunder" type="button" :disabled="store.loading" @click="store.muat()">Muat ulang</button>
    </div>

    <p v-if="pesan" :class="['pesan', pesan.jenis]">{{ pesan.teks }}</p>

    <p v-if="store.loading" class="pesan info">Memuat...</p>
    <p v-else-if="store.error" class="pesan error">{{ store.error }}</p>
    <p v-else-if="store.daftar.length === 0" class="kosong">
      Belum ada buku. <RouterLink to="/tambah">Tambah yang pertama</RouterLink>
    </p>
    <div v-else class="grid-buku">
      <KartuBuku v-for="b in store.daftar" :key="b.id" :buku="b" @hapus="hapus" />
    </div>
  </section>
</template>
