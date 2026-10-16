<script setup>
import { ref, computed } from 'vue'
import FormTambahItem from './components/FormTambahItem.vue'
import ItemBelanja from './components/ItemBelanja.vue'
import RingkasanBelanja from './components/RingkasanBelanja.vue'

/**
 * SESI 9: Komponen Vue -- Props & Emit
 * Demo: "Daftar Belanja" versi 3 -- dipecah jadi komponen.
 *
 * Ini EVOLUSI dari App.vue Sesi 8: logika dan tampilan yang tadinya
 * SEMUA ada di satu file sekarang dipecah jadi 3 komponen anak:
 *
 *   App.vue (PARENT -- "otak" & penyimpan data)
 *     |
 *     |--- FormTambahItem   : emit('tambah', {...})
 *     |--- ItemBelanja (xN) : props item, emit('hapus', id)
 *     |--- RingkasanBelanja : props totalHarga, jumlahItem, dst (read-only)
 *
 * PRINSIP INTI: "Props turun, Event naik"
 *   - Data mengalir DARI App.vue KE komponen anak lewat props.
 *   - Aksi pengguna mengalir DARI komponen anak KE App.vue lewat emit.
 *   - App.vue tetap satu-satunya "pemilik" data daftarBelanja --
 *     komponen anak tidak pernah mengubah data itu sendiri secara
 *     langsung, mereka cuma MEMINTA parent untuk mengubahnya.
 */

const namaToko = ref('Toko Sembako Budi')
const tampilkanHarga = ref(true)

const daftarBelanja = ref([
  { id: 1, nama: 'Beras 5kg', harga: 65000, selesai: false },
  { id: 2, nama: 'Minyak Goreng 2L', harga: 32000, selesai: false },
  { id: 3, nama: 'Telur 1kg', harga: 28000, selesai: true },
])

const totalHarga = computed(() =>
  daftarBelanja.value.reduce((total, item) => total + item.harga, 0)
)
const jumlahItem = computed(() => daftarBelanja.value.length)
const rataRataHarga = computed(() =>
  jumlahItem.value === 0 ? 0 : Math.round(totalHarga.value / jumlahItem.value)
)

// ===== Menangani event dari FormTambahItem =====
// @tambah="tanganiTambah" di template menangkap emit('tambah', data)
function tanganiTambah(dataBaru) {
  console.log('component parent berhasil menangkap value dari child :', dataBaru);

  daftarBelanja.value.push({
    id: daftarBelanja.value.length + 1,
    nama: dataBaru.nama,
    harga: dataBaru.harga,
    selesai: false,
  })
  console.log(daftarBelanja.value, 'terbaru');
}

// ===== Menangani event dari ItemBelanja =====
function tanganiHapus(id) {
  console.log(id,'id');
  daftarBelanja.value = daftarBelanja.value.filter((item) => item.id !== id)

}

function tanganiToggleSelesai(id) {
  const item = daftarBelanja.value.find((i) => i.id === id)
  if (item) item.selesai = !item.selesai
}
</script>

<template>
  <div class="app">
    <header>
      <img src="https://api.iconify.design/mdi/cart-outline.svg?color=white" alt="logo" class="logo" />
      <div>
        <h1>{{ namaToko }}</h1>
        <p>{{ jumlahItem }} item · Total Rp{{ totalHarga.toLocaleString('id-ID') }}</p>
      </div>
    </header>

    <section class="card">
      <h3>Tambah Item (komponen: FormTambahItem)</h3>
      <!--
        @tambah menangkap emit('tambah', ...) dari komponen anak.
        Nama event (tambah) harus SAMA PERSIS dengan yang dipakai
        di emit() dalam FormTambahItem.vue.
      -->
      <FormTambahItem @tambah="tanganiTambah" />
      <p class="catatan">
        Form ini adalah komponen TERPISAH (FormTambahItem.vue). Dia tidak
        tahu-menahu soal daftarBelanja -- begitu pengguna submit, dia cuma
        emit('tambah', data) dan App.vue (parent) yang menentukan apa yang
        terjadi selanjutnya.
      </p>
    </section>

    <!-- RingkasanBelanja hanya menerima props, tidak emit apa-apa -->
    <!-- <RingkasanBelanja :total-harga="totalHarga" :jumlah-item="jumlahItem" :rata-rata-harga="rataRataHarga" /> -->

    <section class="card" style="margin-top: 1rem">
      <h3>Daftar Belanja (komponen: ItemBelanja × {{ jumlahItem }})</h3>
      <label class="toggle" style="margin-bottom: 0.7rem; display: flex">
        <input type="checkbox" v-model="tampilkanHarga" />
        Tampilkan kolom harga
      </label>

      <ul class="daftar">
        <li v-if="daftarBelanja.length === 0">Belanjaan masih kosong.</li>

        <!--
          Satu komponen ItemBelanja dipakai ULANG untuk SETIAP item.
          :item="item" -- mengirim satu objek item sebagai props.
          @hapus dan @toggle-selesai -- menangkap emit dari komponen itu.
        -->
        <ItemBelanja v-for="item in daftarBelanja"  :key="item.id" :item="item" :tampilkan-harga="tampilkanHarga"
          @hapus="tanganiHapus"  @toggle-selesai="tanganiToggleSelesai" />
      </ul>

      <button type="button" v-oncl></button>
    </section>
  </div>
</template>
