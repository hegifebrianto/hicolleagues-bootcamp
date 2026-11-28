// Server PALSU di dalam browser, supaya latihan Pinia jalan tanpa backend.
// Dipasang sebagai "adapter" axios: setiap request axios berakhir di sini, bukan di jaringan.
import { AxiosError } from 'axios'

const AKUN = { email: 'demo@mail.com', password: 'rahasia', nama: 'Demo' }
export const TOKEN_DEMO = 'token-demo-123'

const TUGAS = [
  { id: 1, judul: 'Pelajari lifecycle hooks', selesai: true },
  { id: 2, judul: 'Buat komponen Kartu dengan slot', selesai: true },
  { id: 3, judul: 'Buat tab dinamis + keep-alive', selesai: true },
  { id: 4, judul: 'Pecah store jadi modul', selesai: true },
]

function balas(config, status, body) {
  const response = { data: body, status, statusText: String(status), headers: {}, config, request: {} }
  if (status >= 200 && status < 300) return response
  throw new AxiosError(
    `Request failed with status code ${status}`,
    status >= 500 ? AxiosError.ERR_BAD_RESPONSE : AxiosError.ERR_BAD_REQUEST,
    config, null, response,
  )
}

function ambilHeader(config, nama) {
  const h = config.headers
  return (h && typeof h.get === 'function' ? h.get(nama) : h?.[nama]) || ''
}

export async function mockAdapter(config) {
  await new Promise((r) => setTimeout(r, 300)) // pura-pura lambat
  const metode = (config.method || 'get').toLowerCase()
  const url = config.url

  if (metode === 'post' && url === '/api/login') {
    const b = typeof config.data === 'string' ? JSON.parse(config.data) : config.data || {}
    if (b.email === AKUN.email && b.password === AKUN.password) {
      return balas(config, 200, { data: { token: TOKEN_DEMO, user: { nama: AKUN.nama, email: AKUN.email } } })
    }
    return balas(config, 401, { error: 'email atau password salah' })
  }

  if (metode === 'get' && url === '/api/tugas') {
    if (ambilHeader(config, 'Authorization') !== 'Bearer ' + TOKEN_DEMO) {
      return balas(config, 401, { error: 'token tidak ada atau tidak valid' })
    }
    return balas(config, 200, { data: TUGAS })
  }

  return balas(config, 404, { error: 'tidak ditemukan' })
}
