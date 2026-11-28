<script setup>
import { ref } from 'vue'
import LifecycleDemo from './views/LifecycleDemo.vue'
import SlotDemo from './views/SlotDemo.vue'
import DinamisDemo from './views/DinamisDemo.vue'
import PiniaDemo from './views/PiniaDemo.vue'
import Toast from './components/Toast.vue'

const menu = [
  { id: 'lifecycle', label: '1 · Lifecycle', komponen: LifecycleDemo },
  { id: 'slot', label: '2 · Slots', komponen: SlotDemo },
  { id: 'dinamis', label: '3 · Dinamis', komponen: DinamisDemo },
  { id: 'pinia', label: '4 · Pinia', komponen: PiniaDemo },
]
const aktif = ref('lifecycle')
</script>

<template>
  <div class="container">
    <header>
      <h1>Sesi 26 · Komponen Lanjutan</h1>
    </header>
    <nav>
      <button v-for="m in menu" :key="m.id" :class="{ aktif: aktif === m.id }" @click="aktif = m.id">
        {{ m.label }}
      </button>
    </nav>
    <!-- Tanpa keep-alive di sini: pindah menu = komponen di-unmount (hook unmounted berjalan) -->
    <component :is="menu.find((m) => m.id === aktif).komponen" />
    <Toast />
  </div>
</template>
