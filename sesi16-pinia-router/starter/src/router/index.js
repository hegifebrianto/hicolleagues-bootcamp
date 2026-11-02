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
    { path: '/', component: Beranda },

    // TODO 7 (LATIHAN 2): daftarkan rute berikut -- pola { path: '...', component: ... }
    //   /tugas    -> DaftarTugas   (tambahkan meta: { perluLogin: true })
    //   /tentang  -> Tentang
    //   /login    -> Login
    //
    // TODO 8 (LATIHAN 3): rute dengan parameter dinamis
    //   /tugas/:id -> DetailTugas  (meta: { perluLogin: true })
    //
    // TODO 9 (nested routes): rute /admin memakai AdminLayout dengan `children`:
    //   { path: '', redirect: '/admin/pengguna' }
    //   { path: 'pengguna', component: DaftarPengguna }
    //   { path: 'statistik', component: Statistik }
    //   (beri meta: { perluLogin: true } pada rute induk /admin)
    //
    // TODO 10 (bonus LATIHAN 2): halaman 404, HARUS paling akhir:
    //   { path: '/:pathMatch(.*)*', component: NotFound }
  ],
})

// TODO 11 (bonus LATIHAN 3): navigation guard.
// router.beforeEach((to) => {
//   const auth = useAuthStore()
//   - cek apakah ada rute di to.matched yang meta.perluLogin-nya true (to.matched.some(...))
//   - kalau iya DAN !auth.masuk -> return { path: '/login', query: { redirect: to.fullPath } }
// })

export default router
