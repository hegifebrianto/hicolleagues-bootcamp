<script setup>
import { ref, onBeforeMount, onMounted, onUpdated, onUnmounted } from 'vue'
import { pasangInterval } from '../utils/pantau' // TODO: tambahkan lepasInterval

defineProps({
  // true  = membersihkan timer di onUnmounted (benar)
  // false = LUPA membersihkan (memory leak) -- untuk demo
  bersih: { type: Boolean, default: true },
})
const emit = defineEmits(['log'])

const detik = ref(0)
let timer

onBeforeMount(() => emit('log', 'onBeforeMount: belum ada di DOM'))

onMounted(() => {
  emit('log', 'onMounted: tampil di DOM, timer dimulai')
  timer = pasangInterval(() => detik.value++, 1000)
})

onUpdated(() => {
  if (detik.value % 5 === 0) emit('log', `onUpdated: re-render (detik = ${detik.value})`)
})

// TODO (Bagian 2 slide "Membersihkan Efek Samping"): timer di onMounted belum dibersihkan.
// Isi onUnmounted: panggil lepasInterval(timer) lalu emit('log', 'onUnmounted: timer dibersihkan ✔').
// Setelah benar, angka "Timer yang masih hidup" di demo kembali ke 0 saat Jam disembunyikan.
onUnmounted(() => {
  emit('log', 'onUnmounted: (belum ada pembersihan!)')
})
</script>

<template>
  <div class="jam">⏱ {{ detik }} detik</div>
</template>
