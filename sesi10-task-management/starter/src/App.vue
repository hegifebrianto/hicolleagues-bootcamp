<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import TodoInput from './components/TodoInput.vue'
import TodoList from './components/TodoList.vue'
import TodoFilter from './components/TodoFilter.vue'

/**
 * PROJECT 1 FRONTEND -- Aplikasi Vue To-Do (Sesi 10-11) — STARTER
 *
 * Arsitektur (sesuai slide "Rancangan Komponen"):
 *   App.vue (state di sini -- single source of truth)
 *     |-- TodoInput   : emit('tambah', teks)
 *     |-- TodoFilter  : props filterAktif, emit('ubah-filter', nilai)
 *     |-- TodoList    : props tugas (sudah disaring), meneruskan
 *          |-- TodoItem : props tugas, emit('ubah'/'hapus', id)
 *
 * SESI 10 (Workshop 1): kerjakan dulu TODO 1-6 -- komponen, state,
 * tambah, selesai, hapus, dan hitung tugas tersisa.
 * SESI 11 (Workshop 2): lanjutkan TODO 7-9 -- filter dan localStorage.
 *
 * Kerjakan TODO berurutan. Template di bawah sudah lengkap -- fokus
 * isi logikanya di bagian <script> ini saja. Komponen anak
 * (TodoInput, TodoItem) PUNYA TODO SENDIRI -- kerjakan itu dulu
 * (lihat src/components/) sebelum lanjut ke sini, supaya event-nya
 * benar-benar sampai ke App.vue ini.
 */

const KEY_STORAGE = 'tugas'

// TODO 1: ref array kosong untuk tugas -- tiap elemen nanti
// berbentuk { id, teks, selesai }
const tugas = ref([])

// TODO 2: ref filter, nilai awal 'semua' ('semua' | 'aktif' | 'selesai')
const filter = ref('semua')

// TODO 3: lengkapi computed ini supaya mengembalikan tugas yang
// SUDAH DISARING sesuai filter.value:
//   - 'aktif'   -> hanya tugas yang BELUM selesai
//   - 'selesai' -> hanya tugas yang SUDAH selesai
//   - selain itu ('semua') -> tugas.value apa adanya
const tugasTersaring = computed(() => {
  return tugas.value // <-- ganti sesuai TODO 3
})

// TODO 4: computed yang menghitung JUMLAH tugas yang belum selesai
// (pakai .filter(...).length)
const sisaTugas = computed(() => 0) // <-- ganti sesuai TODO 4

function tambahTugas(teksBaru) {
  // TODO 5: trim() teksBaru -- kalau kosong, return (jangan tambah).
  // Kalau tidak kosong, push objek baru ke tugas.value:
  //   { id: Date.now(), teks: <hasil trim>, selesai: false }
}

function hapusTugas(id) {
  // TODO 6a: tugas.value = tugas.value.filter(...) -- buang satu
  // tugas yang id-nya SAMA DENGAN parameter id.
}

function toggleSelesai(id) {
  // TODO 6b: cari tugas yang id-nya cocok (pakai .find), lalu balik
  // nilai t.selesai-nya (t.selesai = !t.selesai). Cek dulu tugasnya
  // ketemu (tidak undefined) sebelum mengubah propertinya.
}

function ubahFilter(nilai) {
  // TODO 7: filter.value = nilai
}

// TODO 8 (Sesi 11): baca localStorage.getItem(KEY_STORAGE) di dalam
// onMounted(() => {...}). Kalau hasilnya ADA, ubah dari JSON jadi
// array lagi dengan JSON.parse(...), lalu isikan ke tugas.value.
onMounted(() => {})

// TODO 9 (Sesi 11): pasang watch() yang memantau "tugas", dengan opsi
// { deep: true } sebagai argumen ketiga. Di dalam callback-nya,
// simpan ke localStorage:
//   localStorage.setItem(KEY_STORAGE, JSON.stringify(tugasBaru))
//
// watch(tugas, (tugasBaru) => {
//   ...
// }, { deep: true })
</script>

<template>
  <div class="app">
    <header>
      <h1>✅ Aplikasi To-Do</h1>
      <p>Project 1 Frontend -- Sesi 10-11 · {{ sisaTugas }} tugas tersisa</p>
    </header>

    <TodoInput @tambah="tambahTugas" />
    <TodoFilter :filter-aktif="filter" @ubah-filter="ubahFilter" />
    <TodoList :tugas="tugasTersaring" @ubah="toggleSelesai" @hapus="hapusTugas" />

    <p class="catatan">
      Sesi 11: coba tambah beberapa tugas, tandai selesai, lalu REFRESH
      halaman ini -- setelah TODO 8-9 selesai, semua akan tetap ada
      karena tersimpan di localStorage browser.
    </p>
  </div>
</template>
