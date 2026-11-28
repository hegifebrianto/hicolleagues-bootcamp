<script setup>
import { useUiStore } from '../stores/ui'
const ui = useUiStore()
</script>

<template>
  <div v-if="ui.pesan" class="toast">{{ ui.pesan }}</div>
</template>
