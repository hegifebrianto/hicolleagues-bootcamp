<script setup>
defineProps({ buka: Boolean })
const emit = defineEmits(['tutup'])
</script>

<template>
  <div v-if="buka" class="modal-latar" @click.self="emit('tutup')">
    <div class="modal">
      <header><slot name="header">Judul</slot></header>
      <main><slot /></main>
      <footer>
        <!-- named slot + scoped: modal mengirim fungsi tutup ke parent -->
        <slot name="footer" :tutup="() => emit('tutup')">
          <button @click="emit('tutup')">Tutup</button>
        </slot>
      </footer>
    </div>
  </div>
</template>
