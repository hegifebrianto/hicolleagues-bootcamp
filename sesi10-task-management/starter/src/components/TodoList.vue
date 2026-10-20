<script setup>
import TodoItem from './TodoItem.vue'

/**
 * TodoList -- menerima daftar tugas (SUDAH disaring oleh App.vue)
 * lewat props, lalu merender satu TodoItem untuk tiap tugas.
 * Event dari TodoItem (ubah, hapus) diteruskan lagi ke atas ke
 * App.vue -- TodoList sendiri tidak mengubah data apa pun.
 */
defineProps({
  tugas: { type: Array, required: true },
})

const emit = defineEmits(['ubah', 'hapus'])
</script>

<template>
  <ul class="daftar">
    <li v-if="tugas.length === 0" class="kosong">Tidak ada tugas untuk filter ini.</li>
    <TodoItem
      v-for="t in tugas"
      :key="t.id"
      :tugas="t"
      @ubah="emit('ubah', $event)"
      @hapus="emit('hapus', $event)"
    />
  </ul>
</template>
