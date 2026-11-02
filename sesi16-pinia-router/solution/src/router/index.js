import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import Beranda from '../views/Beranda.vue'
import DaftarTugas from '../views/DaftarTugas.vue'
import DetailTugas from '../views/DetailTugas.vue'
import Tentang from '../views/Tentang.vue'
import Login from '../views/Login.vue'
import NotFound from '../views/NotFound.vue'
import AdminLayout from '../views/AdminLayout.vue'
import DaftarPengguna from '../views/DaftarPengguna.vue'
import Statistik from '../views/Statistik.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    // LATIHAN 2: tiga halaman + navigasi
    { path: '/', component: Beranda },
    { path: '/tugas', component: DaftarTugas, meta: { perluLogin: true } },
    { path: '/tentang', component: Tentang },
    { path: '/login', component: Login },

    // LATIHAN 3: route param :id
    { path: '/tugas/:id', component: DetailTugas, meta: { perluLogin: true } },

    // Nested routes: layout bersama + halaman anak lewat `children`
    {
      path: '/admin',
      component: AdminLayout,
      meta: { perluLogin: true },
      children: [
        { path: '', redirect: '/admin/pengguna' },
        { path: 'pengguna', component: DaftarPengguna },
        { path: 'statistik', component: Statistik },
      ],
    },

    // Bonus Latihan 2: halaman 404 (HARUS paling akhir)
    { path: '/:pathMatch(.*)*', component: NotFound },
  ],
})

// Navigation guard: dijalankan SEBELUM setiap perpindahan rute
router.beforeEach((to) => {
  const auth = useAuthStore()
  // to.matched berisi rute induk + anak, jadi meta induk (/admin) ikut terperiksa
  const perluLogin = to.matched.some((r) => r.meta.perluLogin)
  if (perluLogin && !auth.masuk) {
    // return objek = arahkan ke rute lain; simpan tujuan awal untuk kembali setelah login
    return { path: '/login', query: { redirect: to.fullPath } }
  }
})

export default router
