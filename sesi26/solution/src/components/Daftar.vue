<script setup>
defineProps({ items: { type: Array, required: true } })
</script>

<template>
  <ul class="daftar">
    <li v-for="(item, i) in items" :key="item.id">
      <!-- scoped slot: child mengirim item & index, parent yang mendesain tampilannya -->
      <slot :item="item" :index="i">{{ item.judul }}</slot>
    </li>
  </ul>
</template>
