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

// TODO 12: validasi klien. Kembalikan objek { judul, penulis, tahun } berisi pesan error
// ('' kalau valid):
//   judul   -> 'Judul wajib diisi'     bila form.value.judul.trim() kosong
//   penulis -> 'Penulis wajib diisi'   bila form.value.penulis.trim() kosong
//   tahun   -> 'Tahun harus berupa angka 1000–' + tahunMaks  bila bukan bilangan bulat 1000..tahunMaks
const tahunMaks = new Date().getFullYear() + 1
const error = computed(() => ({ judul: '', penulis: '', tahun: '' })) // <-- ganti sesuai TODO 12
const valid = computed(() => !error.value.judul && !error.value.penulis && !error.value.tahun)

onMounted(async () => {
  if (!modeUbah) return
  // TODO 13 (mode ubah): muat data lama ke form.
  //   try   { const b = await store.ambil(id); form.value = { judul: b.judul || '', penulis: b.penulis || '', tahun: b.tahun ?? '' } }
  //   catch { galatMuat.value = ... }   // tidak ada e.response -> 'Server tidak dapat dihubungi.'
  //                                     // status 404 -> 'Buku tidak ditemukan.'  selain itu -> 'Gagal memuat buku.'
  //   finally { memuat.value = false }
  memuat.value = false // <-- hapus baris ini setelah TODO 13 memakai finally
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
    // TODO 14: mode ubah -> await store.ubah(id, payload)   (PUT)
    //          mode tambah -> await store.tambah(payload)    (POST)
    // TODO 15: sukses -> router.push('/') untuk kembali ke daftar
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
