// BACKEND CADANGAN (tanpa dependency) -- kontrak API-nya SAMA PERSIS dengan backend Go.
// Pakai ini kalau Go bermasalah saat demo:   node mock-server.mjs
// Opsi env: PORT (default 8089), DELAY_MS (default 700), CORS_ORIGIN
import http from 'node:http'

const PORT = process.env.PORT || '8089'
const DELAY_MS = process.env.DELAY_MS !== undefined ? Number(process.env.DELAY_MS) : 700
const CORS_ORIGIN = process.env.CORS_ORIGIN || 'http://localhost:5173'

function koleksi(base, bawaan, data, nextID) {
  return { base, bawaan, data, nextID }
}

const tugas = koleksi('/api/tugas', { selesai: false }, [
  { id: 1, judul: 'Belajar Axios', selesai: true },
  { id: 2, judul: 'Belajar loading state', selesai: false },
  { id: 3, judul: 'Sambungkan Vue ke API Go', selesai: false },
], 4)

const buku = koleksi('/api/buku', {}, [
  { id: 1, judul: 'Laskar Pelangi', penulis: 'Andrea Hirata', tahun: 2005 },
  { id: 2, judul: 'Bumi Manusia', penulis: 'Pramoedya Ananta Toer', tahun: 1980 },
  { id: 3, judul: 'Filosofi Teras', penulis: 'Henry Manampiring', tahun: 2019 },
], 4)

function kirim(res, status, body) {
  if (body === undefined) {
    res.writeHead(status)
    res.end()
    return
  }
  res.writeHead(status, { 'Content-Type': 'application/json' })
  res.end(JSON.stringify(body))
}
const galat = (res, status, pesan) => kirim(res, status, { error: pesan })

function bacaBody(req) {
  return new Promise((resolve) => {
    let teks = ''
    req.on('data', (c) => (teks += c))
    req.on('end', () => {
      try {
        const v = JSON.parse(teks)
        resolve(v && typeof v === 'object' && !Array.isArray(v) ? v : null)
      } catch {
        resolve(null)
      }
    })
  })
}

async function tangani(k, req, res, sisa) {
  const m = req.method
  if (sisa === '') {
    if (m === 'GET') return kirim(res, 200, { data: k.data })
    if (m === 'POST') {
      const masuk = await bacaBody(req)
      if (!masuk) return galat(res, 400, 'body tidak valid')
      if (typeof masuk.judul !== 'string' || masuk.judul.trim() === '') return galat(res, 400, 'judul wajib diisi')
      const baru = { ...k.bawaan, ...masuk, id: k.nextID++ }
      k.data.push(baru)
      return kirim(res, 201, { data: baru })
    }
    return galat(res, 405, 'metode tidak didukung')
  }

  if (!/^-?\d+$/.test(sisa)) return galat(res, 400, 'id tidak valid')
  const id = Number(sisa)
  const i = k.data.findIndex((x) => x.id === id)

  if (m === 'GET') return i < 0 ? galat(res, 404, 'data tidak ditemukan') : kirim(res, 200, { data: k.data[i] })
  if (m === 'PUT') {
    const masuk = await bacaBody(req)
    if (!masuk) return galat(res, 400, 'body tidak valid')
    if (i < 0) return galat(res, 404, 'data tidak ditemukan')
    if ('judul' in masuk && (typeof masuk.judul !== 'string' || masuk.judul.trim() === '')) {
      return galat(res, 400, 'judul wajib diisi')
    }
    const { id: _abaikan, ...sisaField } = masuk
    Object.assign(k.data[i], sisaField)
    return kirim(res, 200, { data: k.data[i] })
  }
  if (m === 'DELETE') {
    if (i < 0) return galat(res, 404, 'data tidak ditemukan')
    k.data.splice(i, 1)
    return kirim(res, 204)
  }
  return galat(res, 405, 'metode tidak didukung')
}

const server = http.createServer(async (req, res) => {
  const mulai = Date.now()
  res.on('finish', () => console.log(req.method, req.url, Date.now() - mulai + 'ms'))

  res.setHeader('Access-Control-Allow-Origin', CORS_ORIGIN)
  res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS')
  res.setHeader('Access-Control-Allow-Headers', 'Content-Type')
  if (req.method === 'OPTIONS') return kirim(res, 204)

  const path = new URL(req.url, 'http://x').pathname

  if (DELAY_MS > 0 && req.method === 'GET' && path.startsWith('/api/') && path !== '/api/ping') {
    await new Promise((r) => setTimeout(r, DELAY_MS))
  }

  if (path === '/') {
    res.writeHead(200, { 'Content-Type': 'text/plain; charset=utf-8' })
    return res.end('Backend demo (cadangan Node) aktif. Coba: /api/ping, /api/tugas, /api/buku\n')
  }
  if (path === '/api/ping') return kirim(res, 200, { data: 'pong' })

  for (const k of [tugas, buku]) {
    if (path === k.base || path.startsWith(k.base + '/')) {
      const sisa = path.slice(k.base.length).replace(/^\/+|\/+$/g, '')
      return tangani(k, req, res, sisa)
    }
  }
  return galat(res, 404, 'tidak ditemukan: ' + path)
})

server.listen(PORT, () => {
  console.log('==========================================')
  console.log(' Backend demo CADANGAN (Node)')
  console.log(' Alamat : http://localhost:' + PORT)
  console.log(' Cek    : http://localhost:' + PORT + '/api/ping')
  console.log('==========================================')
})
