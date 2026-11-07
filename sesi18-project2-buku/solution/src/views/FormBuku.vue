<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useBukuStore } from '../stores/buku'

const route = useRoute()
const router = useRouter()
const store = useBukuStore()

// Pola yang sama untuk tambah dan ubah -- bedakan dengan ada/tidaknya :id di URL
const id = route.params.id
const modeUbah = !!id

const form = ref({ judul: '', penulis: '', tahun: '' })
const memuat = ref(modeUbah) // mode ubah: tunggu data lama dimuat dulu
const galatMuat = ref('')
const pesan = ref(null)
const sedangKirim = ref(false)
const dicoba = ref(false) // pesan validasi baru tampil setelah pengguna mencoba menyimpan

// Validasi klien (validasi server tetap wajib -- sudah ada di backend)
const tahunMaks = new Date().getFullYear() + 1
const error = computed(() => {
  const f = form.value
  const tahun = Number(f.tahun)
  return {
    judul: f.judul.trim() ? '' : 'Judul wajib diisi',
    penulis: f.penulis.trim() ? '' : 'Penulis wajib diisi',
    tahun: f.tahun !== '' && Number.isInteger(tahun) && tahun >= 1000 && tahun <= tahunMaks
      ? ''
      : 'Tahun harus berupa angka 1000–' + tahunMaks,
  }
})
const valid = computed(() => !error.value.judul && !error.value.penulis && !error.value.tahun)

onMounted(async () => {
  if (!modeUbah) return
  try {
    const b = await store.ambil(id) // muat data lama ke form
    form.value = { judul: b.judul || '', penulis: b.penulis || '', tahun: b.tahun ?? '' }
  } catch (e) {
    galatMuat.value = !e.response
      ? 'Server tidak dapat dihubungi.'
      : e.response.status === 404
        ? 'Buku tidak ditemukan.'
        : 'Gagal memuat buku.'
  } finally {
    memuat.value = false
  }
})

async function simpan() {
  dicoba.value = true
  if (!valid.value) return // tidak valid -> tidak dikirim

  sedangKirim.value = true
  pesan.value = null
  const payload = {
    judul: form.value.judul.trim(),
    penulis: form.value.penulis.trim(),
    tahun: Number(form.value.tahun),
  }
  try {
    //modeUbah = true ? jika iya maka dia pergi ke statement 1
    if (modeUbah) {
      await store.ubah(id, payload) // PUT
    } else { //false
      await store.tambah(payload) // POST
    }
    router.push('/') // sukses -> kembali ke daftar
  } catch (e) {
    const dariServer = e.response && e.response.data && e.response.data.error
    pesan.value = {
      jenis: 'error',
      teks: dariServer || (e.response ? 'Gagal menyimpan (status ' + e.response.status + ')' : 'Server tidak dapat dihubungi'),
    }
  } finally {
    sedangKirim.value = false
  }
}
</script>

<template>
  <section class="card">
    <h1>{{ modeUbah ? 'Ubah Buku' : 'Tambah Buku' }}</h1>

    <p v-if="memuat" class="pesan info">Memuat data buku...</p>
    <p v-else-if="galatMuat" class="pesan error">{{ galatMuat }}</p>

    <form v-else @submit.prevent="simpan">
      <div class="field">
        <label for="judul">Judul</label>
        <input id="judul" v-model="form.judul" />
        <small v-if="dicoba && error.judul" class="error-text">{{ error.judul }}</small>
      </div>
      <div class="field">
        <label for="penulis">Penulis</label>
        <input id="penulis" v-model="form.penulis" />
        <small v-if="dicoba && error.penulis" class="error-text">{{ error.penulis }}</small>
      </div>
      <div class="field">
        <label for="tahun">Tahun terbit</label>
        <input id="tahun" v-model.number="form.tahun" type="number" />
        <small v-if="dicoba && error.tahun" class="error-text">{{ error.tahun }}</small>
      </div>

      <p v-if="pesan" :class="['pesan', pesan.jenis]">{{ pesan.teks }}</p>

      <div class="aksi-form">
        <button type="submit" :disabled="sedangKirim">{{ sedangKirim ? 'Menyimpan...' : 'Simpan' }}</button>
        <RouterLink class="btn btn-sekunder" to="/">Batal</RouterLink>
      </div>
    </form>

    <p v-if="galatMuat"><RouterLink to="/">← Kembali ke daftar</RouterLink></p>
  </section>
</template>
