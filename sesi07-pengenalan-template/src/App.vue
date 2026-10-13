<script setup>
import { ref } from 'vue'

/**
 * SESI 7: Pengenalan Vue.js & Template Syntax
 * Demo: "Daftar Belanja" versi 1 -- paling dasar.
 *
 * Konsep yang didemokan di file ini:
 *   1. ref()           -> membuat data reaktif
 *   2. {{ }}           -> mustache syntax, menampilkan data
 *   3. v-bind / :attr  -> mengikat atribut HTML ke data
 *   4. v-on / @event   -> menangani event (klik, keyup, dst)
 *   5. v-model         -> data binding dua arah pada input
 *   6. v-if            -> render kondisional (true/false benar2 ada/tidak)
 *   7. v-show          -> tampil/sembunyi kondisional (tetap ada di DOM)
 *   8. v-for           -> perulangan untuk menampilkan daftar
 */

// 1) ref() -- bungkus nilai supaya Vue otomatis mendeteksi perubahannya
const namaToko = ref('Toko Sembako Budi')

// 3) Data untuk v-bind: URL logo (ikon dari iconify, bebas dipakai)
const logoUrl = ref('https://api.iconify.design/mdi/cart-outline.svg?color=white')

// 5) v-model -- dua arah: ketik di input akan otomatis mengubah ini
const itemBaru = ref('')

// 6 & 7) Data untuk v-if / v-show
const tampilkanHarga = ref(true)

// 8) Array untuk v-for
const daftarBelanja = ref([
  { id: 1, nama: 'Beras 5kg', harga: 65000 },
  { id: 2, nama: 'Minyak Goreng 2L', harga: 32000 },
  { id: 3, nama: 'Telur 1kg', harga: 28000 },
])

// 4) Fungsi yang dipanggil lewat v-on / @click
function tambahItem() {
  if (!itemBaru.value.trim()) return
  daftarBelanja.value.push({
    id: Date.now(),
    nama: itemBaru.value,
    harga: 0,
  })
  itemBaru.value = '' // reset input setelah ditambahkan
}

function hapusItem(id) {
  daftarBelanja.value = daftarBelanja.value.filter((item) => item.id !== id)
}
</script>

<template>
  <div class="app">
    <!-- ===== v-bind & mustache syntax ===== -->
    <header>
      <!-- v-bind:src bisa ditulis singkat jadi :src -->
      <img v-bind:src="logoUrl" alt="logo" class="logo" />
      <div>
        <!-- {{ }} = mustache syntax, menyisipkan nilai JS ke HTML -->
        <h1>{{ namaToko }}</h1>
        <p>{{ daftarBelanja.length }} item dalam daftar belanja</p>
      </div>
    </header>

    <!-- ===== v-model & v-on ===== -->
    <section class="tambah card">
      <h3>Tambah Item (v-model + v-on)</h3>
      <div class="tambah">
        <input
          v-model="itemBaru"
          type="text"
          placeholder="Nama item baru..."
          @keyup.enter="tambahItem"
        />
        <!-- v-on:click bisa ditulis singkat jadi @click -->
        <button v-on:click="tambahItem">Tambah</button>
      </div>
      <p class="catatan">
        Coba ketik sesuatu lalu tekan Enter, atau klik tombol Tambah.
        Nilai di kotak input otomatis "terikat" ke variabel itemBaru.
      </p>
    </section>

    <!-- ===== v-if vs v-show ===== -->
    <section class="card">
      <h3>Toggle Tampilan (v-if vs v-show)</h3>
      <label class="toggle">
        <input type="checkbox" v-model="tampilkanHarga" />
        Tampilkan kolom harga
      </label>
      <p class="catatan">
        Checkbox ini pakai v-model juga (untuk tipe checkbox). Kolom harga
        di daftar bawah pakai v-show -- elemennya tetap ADA di HTML, cuma
        disembunyikan lewat CSS display:none. Coba buka DevTools browser
        untuk membuktikannya.
      </p>
    </section>

    <!-- ===== v-for ===== -->
    <section class="card">
      <h3>Daftar Belanja (v-for)</h3>
      <ul class="daftar">
        <!-- v-if di sini: elemen ini BENAR-BENAR tidak ada di HTML bila kondisi salah -->
        <li v-if="daftarBelanja.length === 0">Belanjaan masih kosong.</li>

        <li v-for="item in daftarBelanja" :key="item.id" class="item">
          <span>{{ item.nama }}</span>
          <span v-show="tampilkanHarga" class="harga">
            Rp{{ item.harga.toLocaleString('id-ID') }}
          </span>
          <button class="hapus" @click="hapusItem(item.id)">Hapus</button>
        </li>
      </ul>
    </section>
  </div>
</template>
