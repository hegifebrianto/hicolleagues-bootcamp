<script setup>
/**
 * Komponen anak #2: ItemBelanja
 *
 * Satu komponen ini dipakai BERKALI-KALI oleh App.vue -- satu kali
 * untuk tiap item di daftar belanja (lewat v-for). Setiap pemakaian
 * menerima data item YANG BERBEDA lewat props.
 */

// defineProps mendeklarasikan data apa saja yang BOLEH diterima
// komponen ini dari parent-nya. Props bersifat read-only -- komponen
// ini TIDAK BOLEH mengubah item.harga dkk secara langsung.
const props = defineProps({
  item: { type: Object, required: true },
  tampilkanHarga: { type: Boolean, default: true },
})

// Event yang bisa dipancarkan: 'hapus' (saat tombol hapus diklik)
// dan 'toggle-selesai' (saat item dicentang selesai/belum).
const emit = defineEmits(['hapus', 'toggle-selesai'])
</script>

<template>
  <li class="item" :class="{ selesai: item.selesai }">
    <input type="checkbox" :checked="item.selesai" @change="emit('toggle-selesai', item.id)" />
    <span>{{ item.nama }}</span>
    <span v-show="tampilkanHarga" class="harga">
      Rp{{ item.harga.toLocaleString('id-ID') }}
    </span>
    <!-- emit('hapus', item.id) -- mengirim ID item yang mau dihapus ke parent -->
    <button class="hapus" @click="emit('hapus', item.id)">Hapus</button>
  </li>
</template>
