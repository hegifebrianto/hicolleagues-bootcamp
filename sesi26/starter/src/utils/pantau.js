// Membungkus setInterval/clearInterval agar kita bisa MELIHAT berapa timer yang masih hidup.
// Dipakai untuk mendemokan kebocoran memori (lifecycle).
import { reactive } from 'vue'

const hidup = new Set()
export const pantau = reactive({ aktif: 0 })

export function pasangInterval(fn, ms) {
  const id = setInterval(fn, ms)
  hidup.add(id)
  pantau.aktif = hidup.size
  return id
}

export function lepasInterval(id) {
  clearInterval(id)
  hidup.delete(id)
  pantau.aktif = hidup.size
}

// Tombol darurat di demo: matikan semua timer yang "bocor".
export function lepasSemua() {
  for (const id of hidup) clearInterval(id)
  hidup.clear()
  pantau.aktif = 0
}

// Hitung berapa kali sebuah komponen dibuat (mounted) / diaktifkan (activated).
export const hitung = reactive({
  Profil: { mounted: 0, activated: 0 },
  Pengaturan: { mounted: 0, activated: 0 },
})
