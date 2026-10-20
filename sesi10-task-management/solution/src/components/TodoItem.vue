<script setup>
/**
 * TodoItem -- menampilkan SATU tugas. Dipakai berkali-kali oleh
 * TodoList lewat v-for. Props: tugas (read-only). Emit: 'ubah'
 * (toggle selesai) dan 'hapus'.
 */
defineProps({
  tugas: { type: Object, required: true }, // { id, teks, selesai }
})

const emit = defineEmits(['ubah', 'hapus'])


function changeEmit(id) {
emit('ubah', id)
}

</script>

<template>
  <li class="item" :class="{ selesai: tugas.selesai }">
    <input type="checkbox" :checked="tugas.selesai" @change="changeEmit(tugas.id)" />
    <span>{{ tugas.teks }}</span>
    <button @click="emit('hapus', tugas.id)">Hapus</button>
  </li>
</template>
