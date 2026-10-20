<script setup>
/**
 * TodoItem -- menampilkan SATU tugas. Dipakai berkali-kali oleh
 * TodoList lewat v-for. Props: tugas (read-only). Emit: 'ubah'
 * (toggle selesai) dan 'hapus'.
 */
defineProps({
  tugas: { type: Object, required: true }, // { id, teks, selesai }
})

// TODO 1: daftarkan dua event yang BOLEH dipancarkan komponen ini:
// 'ubah' dan 'hapus'
const emit = defineEmits([])
</script>

<template>
  <li class="item" :class="{ selesai: tugas.selesai }">
    <!--
      TODO 2: pasang @change di checkbox ini supaya memancarkan
      emit('ubah', tugas.id) setiap kali dicentang/dibatalkan.
    -->
    <input type="checkbox" :checked="tugas.selesai" />
    <span>{{ tugas.teks }}</span>
    <!--
      TODO 3: pasang @click di tombol ini supaya memancarkan
      emit('hapus', tugas.id).
    -->
    <button>Hapus</button>
  </li>
</template>
