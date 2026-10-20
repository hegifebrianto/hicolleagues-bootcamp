<script setup>
/**
 * TodoFilter -- menerima filter yang sedang aktif lewat props,
 * memancarkan 'ubah-filter' saat pengguna memilih tombol lain.
 */
defineProps({
  filterAktif: { type: String, required: true }, // 'semua' | 'aktif' | 'selesai'
})

const emit = defineEmits(['ubah-filter'])
</script>

<template>
  <div class="filter-group">
    <button :class="{ aktif: filterAktif === 'semua' }" @click="emit('ubah-filter', 'semua')">
      Semua
    </button>
    <button :class="{ aktif: filterAktif === 'aktif' }" @click="emit('ubah-filter', 'aktif')">
      Aktif
    </button>
    <button :class="{ aktif: filterAktif === 'selesai' }" @click="emit('ubah-filter', 'selesai')">
      Selesai
    </button>
  </div>
</template>
