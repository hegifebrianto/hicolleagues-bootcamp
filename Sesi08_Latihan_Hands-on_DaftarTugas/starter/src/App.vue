<script setup>
import { ref, computed, watch, onMounted } from 'vue'

/**
 * Sesi 8 — Latihan Hands-on (25 menit) — VERSI STARTER
 *   Latihan 1 (9') Daftar Tugas    -> v-for, @click, v-show, v-if
 *   Latihan 2 (8') Filter Daftar   -> computed
 *   Latihan 3 (8') Simpan Otomatis -> watch + localStorage
 * Alur kerja: pseudocode -> kode -> uji dengan contoh
 *
 * Kerjakan TODO 1-8 secara berurutan. Template di bawah sudah
 * lengkap -- fokus isi logikanya di bagian <script> ini saja.
 */

const KEY_STORAGE = 'daftar-tugas'

// ================= LATIHAN 1: Daftar Tugas =================
// TODO 1: tugas.value akan menyimpan array objek {id, teks}.
// (Sudah disiapkan kosong di bawah -- tidak perlu diubah.)
const tugas = ref([])
const inputBaru = ref('')

function tambahTugas() {
  // TODO 2: ambil inputBaru.value, panggil .trim() -- kalau hasilnya
  // kosong ("") maka langsung "return" (jangan tambahkan apa-apa).

  // TODO 3: kalau tidak kosong, push objek baru ke tugas.value:
  //   { id: Date.now(), teks: <hasil trim tadi> }

  // TODO 4: setelah berhasil ditambahkan, kosongkan inputBaru.value
  // supaya kotak input bersih lagi untuk ketikan berikutnya.
}

function hapusTugas(id) {
  // TODO 5: tugas.value = tugas.value.filter(...) -- buang satu
  // objek yang t.id-nya SAMA DENGAN parameter id di atas.
}

// Bonus Latihan 1 (sudah lengkap, contoh pola computed sederhana):
const jumlahTugas = computed(() => tugas.value.length)

// ================= LATIHAN 2: Filter Daftar =================
const kataKunci = ref('')

// TODO 6: lengkapi computed ini supaya:
//   - bila kataKunci KOSONG (setelah trim), kembalikan semua tugas
//     (tugas.value) apa adanya.
//   - selain itu, kembalikan HANYA tugas yang teksnya MENGANDUNG
//     kataKunci (pakai .toLowerCase() di kedua sisi supaya
//     pencarian tidak peka huruf besar/kecil, lalu .includes()).
const tugasTersaring = computed(() => {
  return tugas.value // <-- ganti baris ini sesuai TODO 6
})

// ================= LATIHAN 3: Simpan Otomatis =================
onMounted(() => {
  // TODO 7: baca localStorage.getItem(KEY_STORAGE). Kalau hasilnya
  // ADA (bukan null), ubah dari teks JSON jadi array lagi dengan
  // JSON.parse(...), lalu isikan ke tugas.value.
})

const pesanTersimpan = ref(false)
let timerPesan = null

// TODO 8: pasang watch() yang memantau "tugas", dengan opsi
// { deep: true } sebagai argumen ketiga (supaya perubahan ISI
// array seperti push/filter ikut terdeteksi). Di dalam callback-nya:
//   - simpan ke localStorage: localStorage.setItem(KEY_STORAGE, JSON.stringify(tugasBaru))
//   - set pesanTersimpan.value = true
//   - clearTimeout(timerPesan), lalu setTimeout 1200ms yang
//     mengembalikan pesanTersimpan.value jadi false (timerPesan = setTimeout(...))
//
// watch(tugas, (tugasBaru) => {
//   ...
// }, { deep: true })
</script>

<template>
  <div class="app">
    <header>
      <h1>📝 Daftar Tugas</h1>
      <p>{{ jumlahTugas }} tugas tersimpan</p>
    </header>

    <!-- v-show: elemen tetap ada di DOM, cuma disembunyikan saat false -->
    <p v-show="pesanTersimpan" class="notif">✓ Tersimpan</p>

    <section class="card">
      <h3>Latihan 1 — Tambah Tugas (v-for, @click, v-show, v-if)</h3>
      <div class="tambah">
        <input
          v-model="inputBaru"
          placeholder="Tugas baru, mis. 'Belajar Vue'..."
          @keyup.enter="tambahTugas"
        />
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
      Latihan 3: setelah TODO 7-8 selesai, coba tambah/hapus tugas lalu
      REFRESH halaman ini -- daftarnya akan tetap ada karena tersimpan
      di localStorage browser.
    </p>
  </div>
</template>
