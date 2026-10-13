<script setup>
import { ref, computed, watch } from 'vue'

/**
 * SESI 8: Vue -- Rendering & Reaktivitas
 * Demo: "Daftar Belanja" versi 2 -- lanjutan dari Sesi 7.
 *
 * Semua konsep Sesi 7 (mustache, v-bind, v-on, v-model, v-if/v-show,
 * v-for) MASIH DIPAKAI di sini. Yang BARU ditambahkan:
 *   1. computed()  -> nilai turunan yang otomatis update & di-cache
 *   2. watch()     -> menjalankan kode saat data tertentu berubah
 *   3. Pencarian/filter yang reaktif (gabungan computed + v-model)
 */

const namaToko = ref('Toko Sembako Budi')
const logoUrl = ref('https://api.iconify.design/mdi/cart-outline.svg?color=white')
const itemBaru = ref('')
const hargaBaru = ref(0)
const tampilkanHarga = ref(true)
const kataKunci = ref('') // untuk fitur pencarian baru

const daftarBelanja = ref([
  { id: 1, nama: 'Beras 5kg', harga: 65000 },
  { id: 2, nama: 'Minyak Goreng 2L', harga: 32000 },
  { id: 3, nama: 'Telur 1kg', harga: 28000 },
  { id: 4, nama: 'Gula Pasir 1kg', harga: 15000 },
])

// ================= 1) COMPUTED =================
// computed() MENGHITUNG ULANG otomatis HANYA ketika data sumbernya
// berubah -- dan hasilnya di-cache (disimpan) selagi belum ada
// perubahan. Beda dari method biasa yang dihitung ulang tiap render.

const totalHarga = computed(() =>
  daftarBelanja.value.reduce((total, item) => total + item.harga, 0)
)

const jumlahItem = computed(() => daftarBelanja.value.length)

const rataRataHarga = computed(() =>
  jumlahItem.value === 0 ? 0 : Math.round(totalHarga.value / jumlahItem.value)
)

// computed juga sering dipakai untuk FILTER data berdasarkan input lain --
// ini contoh search box yang reaktif: ketik di kotak cari, daftar di
// bawah otomatis ter-filter tanpa kode tambahan apa pun.
const daftarTersaring = computed(() => {
  if (!kataKunci.value.trim()) return daftarBelanja.value
  return daftarBelanja.value.filter((item) =>
    item.nama.toLowerCase().includes(kataKunci.value.toLowerCase())
  )
})

// ================= 2) WATCH =================
// watch() dipakai untuk EFEK SAMPING: melakukan sesuatu (bukan
// sekadar menurunkan nilai baru) saat data berubah. Di sini kita
// pakai untuk menampilkan notifikasi singkat setiap kali daftar
// belanja bertambah atau berkurang.
const pesanNotifikasi = ref('')
let timerNotifikasi = null

watch(
  daftarBelanja,
  (daftarBaru, daftarLama) => {
    if (daftarBaru.length > daftarLama.length) {
      pesanNotifikasi.value = '✓ Item baru ditambahkan ke daftar!'
    } else if (daftarBaru.length < daftarLama.length) {
      pesanNotifikasi.value = '✓ Item berhasil dihapus.'
    }
    // notifikasi hilang otomatis setelah 2 detik
    clearTimeout(timerNotifikasi)
    timerNotifikasi = setTimeout(() => {
      pesanNotifikasi.value = ''
    }, 2000)
  },
  { deep: true }
)

function tambahItem() {
  if (!itemBaru.value.trim()) return
  daftarBelanja.value.push({
    id: Date.now(),
    nama: itemBaru.value,
    harga: Number(hargaBaru.value) || 0,
  })
  itemBaru.value = ''
  hargaBaru.value = 0
}

function hapusItem(id) {
  daftarBelanja.value = daftarBelanja.value.filter((item) => item.id !== id)
}
</script>

<template>
  <div class="app">
    <header>
      <img v-bind:src="logoUrl" alt="logo" class="logo" />
      <div>
        <h1>{{ namaToko }}</h1>
        <!-- computed dipakai langsung di template, sama seperti data biasa -->
        <p>{{ jumlahItem }} item · Total Rp{{ totalHarga.toLocaleString('id-ID') }}</p>
      </div>
    </header>

    <!-- notifikasi hasil dari watch() -->
    <p v-if="pesanNotifikasi" class="notif">{{ pesanNotifikasi }}</p>

    <section class="card">
      <h3>Tambah Item</h3>
      <div class="tambah">
        <input v-model="itemBaru" type="text" placeholder="Nama item..." @keyup.enter="tambahItem" />
        <input v-model.number="hargaBaru" type="number" placeholder="Harga" style="max-width: 120px" />
        <button @click="tambahItem">Tambah</button>
      </div>
      <p class="catatan">
        Perhatikan notifikasi hijau yang muncul sesaat setelah kamu
        menambah/menghapus item -- itu hasil kerja watch(), bukan ditulis manual
        di fungsi tambahItem()/hapusItem().
      </p>
    </section>

    <!-- ===== Ringkasan dari COMPUTED ===== -->
    <section class="ringkasan">
      <span>Total: <b>Rp{{ totalHarga.toLocaleString('id-ID') }}</b></span>
      <span>Rata-rata: <b>Rp{{ rataRataHarga.toLocaleString('id-ID') }}</b></span>
      <span>Jumlah: <b>{{ jumlahItem }}</b></span>
    </section>

    <!-- ===== Pencarian reaktif (computed + v-model) ===== -->
    <section class="card" style="margin-top: 1rem">
      <h3>Cari Item (computed + v-model)</h3>
      <input v-model="kataKunci" type="text" placeholder="Ketik untuk mencari, mis. 'minyak'..." />
      <p class="catatan">
        Daftar di bawah otomatis ter-filter mengikuti ketikanmu --
        tidak ada tombol "Cari" sama sekali, karena daftarTersaring
        adalah computed yang bereaksi otomatis terhadap kataKunci.
      </p>
    </section>

    <section class="card">
      <h3>Daftar Belanja</h3>
      <label class="toggle" style="margin-bottom: 0.7rem; display: flex">
        <input type="checkbox" v-model="tampilkanHarga" />
        Tampilkan kolom harga
      </label>
      <ul class="daftar">
        <li v-if="daftarTersaring.length === 0">Tidak ada item yang cocok.</li>
        <li v-for="item in daftarTersaring" :key="item.id" class="item">
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
