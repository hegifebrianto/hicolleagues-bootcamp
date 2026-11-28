<script setup>
import { ref, onBeforeMount, onMounted, onUpdated, onUnmounted } from 'vue'
import { pasangInterval, lepasInterval } from '../utils/pantau'

const props = defineProps({
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

onUnmounted(() => {
  if (props.bersih) {
    lepasInterval(timer)
    emit('log', 'onUnmounted: timer dibersihkan ✔')
  } else {
    emit('log', 'onUnmounted: timer TIDAK dibersihkan ✘ (bocor!)')
  }
})
</script>

<template>
  <div class="jam">⏱ {{ detik }} detik</div>
</template>
