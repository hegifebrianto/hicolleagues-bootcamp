import axios from 'axios'
import { mockAdapter } from './mock'

// Satu instance axios untuk seluruh aplikasi.
// adapter = server palsu (lihat mock.js). Untuk backend sungguhan: hapus baris adapter
// dan isi baseURL.
export const api = axios.create({ adapter: mockAdapter })
