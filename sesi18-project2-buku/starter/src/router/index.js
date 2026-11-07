import { createRouter, createWebHistory } from 'vue-router'
import DaftarBuku from '../views/DaftarBuku.vue'
import FormBuku from '../views/FormBuku.vue'
import DetailBuku from '../views/DetailBuku.vue'
import NotFound from '../views/NotFound.vue'

// Rute sesuai slide: /, /tambah, /ubah/:id, /detail/:id
export default createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: DaftarBuku },
    { path: '/tambah', component: FormBuku },
    { path: '/ubah/:id', component: FormBuku }, // komponen yang SAMA dipakai untuk tambah & ubah
    { path: '/detail/:id', component: DetailBuku },
    { path: '/:pathMatch(.*)*', component: NotFound },
  ],
})
