<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import TodoInput from './components/TodoInput.vue'
import TodoList from './components/TodoList.vue'
import TodoFilter from './components/TodoFilter.vue'

/**
 * PROJECT 1 FRONTEND -- Aplikasi Vue To-Do (Sesi 10-11) — SOLUSI
 *
 * Arsitektur (sesuai slide "Rancangan Komponen"):
 *   App.vue (state di sini -- single source of truth)
 *     |-- TodoInput   : emit('tambah', teks)
 *     |-- TodoFilter  : props filterAktif, emit('ubah-filter', nilai)
 *     |-- TodoList    : props tugas (sudah disaring), meneruskan
 *          |-- TodoItem : props tugas, emit('ubah'/'hapus', id)
 *
 * Fitur wajib: tambah, selesai, hapus, hitung tugas (Sesi 10)
 * Fitur lanjutan: filter Semua/Aktif/Selesai, localStorage (Sesi 11)
 */

const KEY_STORAGE = 'tugas'

const tugas = ref([]) // [{ id, teks, selesai }]
const filter = ref('semua') // 'semua' | 'aktif' | 'selesai'

// computed: daftar yang ditampilkan, tersaring sesuai filter aktif
const tugasTersaring = computed(() => {
  if (filter.value === 'aktif') return tugas.value.filter((t) => !t.selesai)
  if (filter.value === 'selesai') return tugas.value.filter((t) => t.selesai)
  return tugas.value
})

// computed: jumlah tugas yang belum selesai
const sisaTugas = computed(() => tugas.value.filter((t) => !t.selesai).length)

function tambahTugas(teksBaru) {
  const bersih = teksBaru.trim() //value ' coba '=== 'coba' 
  if (!bersih) return // input kosong -> tidak ditambahkan
  
  tugas.value.push({ id: Date.now(), teks: bersih, selesai: false })
  console.log(tugas.value)
}

function hapusTugas(id) {
  tugas.value = tugas.value.filter((t) => t.id !== id)
}

function toggleSelesai(id) {
  const t = tugas.value.find((t) => t.id === id)
  if (t) t.selesai = !t.selesai
}

function ubahFilter(nilai) {
  filter.value = nilai
}


//id name type
//01 | anggi | true 


// Muat dari localStorage sekali saat halaman pertama kali dibuka
onMounted(() => {
  const data = localStorage.getItem(KEY_STORAGE)
  console.log('data :',data);
  if (data) tugas.value = JSON.parse(data)
})

// Simpan ke localStorage setiap kali tugas berubah (deep: true supaya
// perubahan ISI array -- push/filter/toggle -- ikut terdeteksi)
watch(
  tugas,
  (tugasBaru) => {
    localStorage.setItem(KEY_STORAGE, JSON.stringify(tugasBaru))
  },
  { deep: true }
)
</script>

<template>
  <div class="app">
    <header>
      <h1>✅ Aplikasi To-Do</h1>
      <p>Project 1 Frontend -- Sesi 10 · {{ sisaTugas }} tugas tersisa</p>
    </header>

    <TodoInput @tambah="tambahTugas" />
    <TodoFilter :filter-aktif="filter" @ubah-filter="ubahFilter" />
    
    
    <TodoList :tugas="tugasTersaring" @ubah="toggleSelesai" @hapus="hapusTugas" />

    <p class="catatan">
      Coba tambah beberapa tugas, tandai selesai, lalu REFRESH halaman ini --
      semua akan tetap ada karena tersimpan di localStorage browser.
    </p>
  </div>
</template>
