<script setup>

// Komponen kecil yang dipakai ulang untuk tiap buku. Props turun, event naik.
defineProps({
  buku: { type: Object, required: true },
})
  // const {id ,judul, penulis,tahun } = buku;

const emit = defineEmits(['hapus'])

</script>

<template>
  <article class="kartu-buku">
    <h3>{{ buku.judul }}</h3>
    <p class="muted">{{ buku.penulis }} · {{ buku.tahun }}</p>
    <div class="aksi">
      <RouterLink class="btn btn-sekunder" :to="'/detail/' + buku.id">Detail</RouterLink>
      <RouterLink class="btn btn-sekunder" :to="'/ubah/' + buku.id">Ubah</RouterLink>
      <button class="btn-hapus" type="button" @click="emit('hapus', buku)">Hapus</button>
    </div>
  </article>
</template>
