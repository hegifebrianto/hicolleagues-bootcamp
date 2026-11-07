<script setup>
import StatusBackend from './components/StatusBackend.vue'
import FormTugas from './components/FormTugas.vue'
import DaftarTugas from './components/DaftarTugas.vue'
</script>

<template>
  <div class="app">
    <h1>Sesi 17 · Vue + API Go</h1>
    <p class="muted">Loading state, error state, form + validasi, dan GET / POST / PUT / DELETE.</p>
    <StatusBackend />
    <FormTugas />
    <DaftarTugas />
  </div>
</template>
