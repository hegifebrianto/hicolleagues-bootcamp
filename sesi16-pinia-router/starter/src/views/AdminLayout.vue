<template>
  <section class="card">
    <h1>Admin</h1>
    <!-- layout bersama; halaman anak tampil di RouterView di bawah -->
    <div class="sub-menu">
      <RouterLink to="/admin/pengguna">Pengguna</RouterLink>
      <RouterLink to="/admin/statistik">Statistik</RouterLink>
    </div>
    <RouterView />
  </section>
</template>
