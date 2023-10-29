import { mapGetters } from 'vuex'

export default function ({ store, route, redirect }) {
  // Cek apakah pengguna sudah login atau belum.
  const { isAuthenticated } = mapGetters(['isAuthenticated', 'loggedInUser'])

  // Jika pengguna belum login dan mencoba mengakses halaman beranda,
  // alihkan mereka ke halaman login atau halaman lain yang sesuai.
  if (!isAuthenticated && route.name !== 'login') {
    return redirect('/login')
  }
}
