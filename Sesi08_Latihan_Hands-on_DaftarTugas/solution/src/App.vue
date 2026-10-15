<script setup>
import { ref, computed, watch, onMounted } from 'vue'

/**
 * Sesi 8 — Latihan Hands-on (25 menit) — VERSI SOLUSI
 *   Latihan 1 (9') Daftar Tugas    -> v-for, @click, v-show, v-if
 *   Latihan 2 (8') Filter Daftar   -> computed
 *   Latihan 3 (8') Simpan Otomatis -> watch + localStorage
 * Alur kerja: pseudocode -> kode -> uji dengan contoh
 */

const KEY_STORAGE = 'daftar-tugas'

// ================= LATIHAN 1: Daftar Tugas =================
// Simpan array objek {id, teks} -- gunakan :key="t.id" di template.
const tugas = ref([]) //variable array
const inputBaru = ref('') //variable untuk textInput

function tambahTugas() {
  const teks = inputBaru.value.trim()
  if (!teks) {
    return // input kosong -> tidak ditambahkan
  } else {
    console.log(Date.now(), 'datenow');
    tugas.value.push({ id: Date.now(), teks })
    inputBaru.value = ''
  }

}

function test() {

}

// Bonus Latihan 1: penghitung jumlah tugas
const jumlahTugas = computed(() => tugas.value.length)
const jumlahTugasTanpaComputed = (() => tugas.value.length)

function hapusTugas(id) {
  console.log(id);
  // console.log('tempTugas :', tugas.value.filter((t) => t.id !== id))
  const tempTugas = tugas.value.filter((t) => t.id !== id)
  for (let i = 0; i < tugas.value.length; i++) {
    console.log(tugas.value[i].id ,':',id)
    if (tugas.value[i].id === id) {
      
      // tugas.value[i].deleted = true;
    }
    
  }
        console.log(tugas.value)

  // console.log(tugas.value)
  // tugas.value = tugas.value.filter((t) => t.id !== id)

}

//upsert (update dan insert)
// checkout shoppe belanja
// tas nike (2)
// sepatu adidas (1)


// ================= LATIHAN 2: Filter Daftar =================
// computed mengembalikan hasil filter; toLowerCase() supaya pencarian
// tidak peka huruf besar/kecil.
const kataKunci = ref('')
let temp = ref(0);

const tugasTersaring = computed(() => {
  console.log(kataKunci.value)
  // return tugas.value.filter((t) => t.teks.toLowerCase().includes(kunci))

  if (!kataKunci.value.trim()) {
    return tugas.value
  } else {
    const kunci = kataKunci.value.toLowerCase()
    console.log(kunci);
    return tugas.value.filter((t) => t.teks.toLowerCase().includes(kunci))
  }

})

// ================= LATIHAN 3: Simpan Otomatis =================
// Muat dari localStorage sekali saat halaman pertama kali dibuka.
onMounted(() => {
  console.log('halaman terender ', temp.value++)
  console.log(localStorage.getItem(KEY_STORAGE));
  const tersimpan = localStorage.getItem(KEY_STORAGE)
  if (tersimpan) {
    tugas.value = JSON.parse(tersimpan)
  }
})

// Bonus Latihan 3: tampilkan pesan "Tersimpan" sesaat
const pesanTersimpan = ref(false)
let timerPesan

// watch dengan { deep: true } -- tugas adalah array of objects,
// deep:true diperlukan supaya perubahan ISI array (push/filter)
// ikut terdeteksi, bukan cuma penggantian array-nya secara utuh.
watch(
  tugas,
  (tugasBaru) => {
    localStorage.setItem(KEY_STORAGE, JSON.stringify(tugasBaru))
    pesanTersimpan.value = true
    clearTimeout(timerPesan)
    timerPesan = setTimeout(() => (pesanTersimpan.value = false), 100)
  },
  { deep: true }
)
</script>

<template>
  <div class="app">
    <header>
      <h1>📝 Daftar Tugas</h1>
      <p>{{ jumlahTugas }} tugas tersimpan</p>
    </header>
    <!-- true false !true !false -->
    <!-- v-show: elemen tetap ada di DOM, cuma disembunyikan saat false -->
    <p v-show="pesanTersimpan" class="notif">✓ Tersimpan</p>

    <section class="card">
      <h3>Latihan 1 — Tambah Tugas (v-for, @click, v-show, v-if)</h3>
      <div class="tambah">
        <input type="text" v-model="inputBaru" placeholder="Tugas baru, mis. 'Belajar Vue'..."
          @keyup.enter="tambahTugas" />
        <button @click="tambahTugas">Tambah</button>
      </div>
    </section>

    <section class="card">
      <h3>Latihan 2 — Cari Tugas (computed)</h3>
      <input v-model="kataKunci" placeholder="Cari tugas, mis. 'bel'..." />
    </section>

    <section class="card">
      <h3>Daftar Tugas</h3>
      <ul class="daftar">
        <!-- v-if: elemen ini benar-benar TIDAK ADA di DOM bila kondisi salah -->
        <li v-if="tugasTersaring.length === 0">Tidak ada hasil.</li>

        <li v-for="t in tugasTersaring" :key="t.id" class="item">
          <span>{{ t.teks }}</span>
          <button class="hapus" @click="hapusTugas(t.id)">Hapus</button>
        </li>
      </ul>
    </section>

    <p class="catatan">
      Latihan 3: coba tambah/hapus tugas lalu REFRESH halaman ini --
      daftarnya akan tetap ada karena tersimpan di localStorage browser.
    </p>
  </div>
</template>
