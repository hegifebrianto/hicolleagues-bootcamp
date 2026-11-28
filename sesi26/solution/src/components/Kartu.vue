<template>
  <div class="kartu">
    <header v-if="$slots.header"><slot name="header" id="slot-content-id"/></header>
    <div class="isi"><slot class="slot-content">Konten default bila kosong</slot></div>
    <footer v-if="$slots.footer"><slot name="footer" /></footer>
  </div>
</template>
