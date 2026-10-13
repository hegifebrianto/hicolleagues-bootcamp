<script setup>
/**
 * Komponen anak #3: RingkasanBelanja
 *
 * Komponen ini CUMA menerima props, tidak pernah emit apa pun --
 * sengaja dibuat begini untuk menunjukkan bahwa tidak semua
 * komponen anak perlu "bicara balik" ke parent-nya. Kadang
 * komponen cuma bertugas MENAMPILKAN data (presentational).
 */
defineProps({
  totalHarga: { type: Number, required: true },
  jumlahItem: { type: Number, required: true },
  rataRataHarga: { type: Number, required: true },
})
</script>

<template>
  <section class="ringkasan">
    <span>Total: <b>Rp{{ totalHarga.toLocaleString('id-ID') }}</b></span>
    <span>Rata-rata: <b>Rp{{ rataRataHarga.toLocaleString('id-ID') }}</b></span>
    <span>Jumlah: <b>{{ jumlahItem }}</b></span>
  </section>
</template>
