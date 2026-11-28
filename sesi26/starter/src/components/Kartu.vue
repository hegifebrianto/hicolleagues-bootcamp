<!-- TODO Latihan 1 (8 menit): Kartu dengan slot
     1. Tambahkan <slot name="header" /> di dalam <header>, tampil HANYA bila diisi: v-if="$slots.header"
     2. Tambahkan slot default (boleh beri teks cadangan) di dalam <div class="isi">
     3. Bonus: slot footer opsional, caranya sama dengan header
     Pemakaiannya sudah disiapkan di views/SlotDemo.vue (Profil & Statistik). -->
<template>
  <div class="kartu">
    <header>Header</header>
    <div class="isi">Isi</div>
  </div>
</template>
