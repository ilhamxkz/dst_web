<script setup>
import { ref, onMounted } from 'vue'
import avatar1 from '@images/avatars/avatar-1.png'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { toast } from 'vue3-toastify'
import 'vue3-toastify/dist/index.css'

const router = useRouter()

// ✅ Helper notif
const notifySuccess = (msg) =>
  toast.success(msg, {
    autoClose: 2000,
    position: 'top-right',
    theme: 'colored',
  })

const notifyError = (msg) =>
  toast.error(msg, {
    autoClose: 2500,
    position: 'top-right',
    theme: 'colored',
  })

// --- state untuk nama & role ---
const userName = ref('Pengguna')
const userRole = ref('—')
const roles = ref([]) // <-- perbaikan: simpan daftar roles di sini

// util: buat header auth
const getAuthHeaders = (token) => ({ Authorization: `Bearer ${token}` })

// util: coba ekstrak dari respons berbagai bentuk
const extractFromResponse = (data) => {
  if (!data) return null

  const candidateObjects = [
    data,
    data.user,
    data.data,
    data.result,
    data.profile,
    data.me,
  ].filter(Boolean)

  for (const obj of candidateObjects) {
    // cari nama
    const name =
      obj.name ||
      obj.username ||
      obj.full_name ||
      obj.fullname ||
      obj.nama ||
      obj.nama_lengkap ||
      obj.displayName ||
      null

    // cari role (bisa string, object, atau id)
    let role = null
    let roleId = null

    if (typeof obj.role === 'string') {
      role = obj.role
    } else if (typeof obj.role === 'number') {
      roleId = obj.role
    } else if (obj.role && typeof obj.role === 'object') {
      role = obj.role.nama_role || obj.role.name || obj.role.role_name || null
      roleId = obj.role.id ?? null
    }

    // beberapa API pakai field role_id atau roleId
    if (!role && (obj.role_id !== undefined || obj.roleId !== undefined)) {
      roleId = obj.role_id ?? obj.roleId
    }

    // jika bukan string tapi ada roles array
    if (!role && Array.isArray(obj.roles) && obj.roles.length) {
      const r = obj.roles[0]
      role = typeof r === 'string' ? r : (r.nama_role || r.name || r.role_name || null)
      if (!role && r.id) roleId = r.id
    }

    // jika kita hanya punya roleId, coba cari nama dari roles.value
    if (!role && roleId != null && Array.isArray(roles.value)) {
      const found = roles.value.find(rr => Number(rr.id) === Number(roleId))
      if (found) role = found.nama_role || found.name || null
    }

    if (name || role) {
      return { name: name || null, role: role || null }
    }
  }

  return null
}

// fallback: decode JWT payload (if token JWT)
const tryDecodeJwt = (token) => {
  try {
    const parts = token.split('.')
    if (parts.length < 2) return null
    const payload = JSON.parse(atob(parts[1].replace(/-/g, '+').replace(/_/g, '/')))
    const name =
      payload.name ||
      payload.username ||
      payload.sub ||
      payload.full_name ||
      payload.nama ||
      null

    let role = null
    let roleId = null
    if (payload.role !== undefined) {
      if (typeof payload.role === 'string') role = payload.role
      else if (typeof payload.role === 'number') roleId = payload.role
      else if (Array.isArray(payload.role) && payload.role.length) role = payload.role[0]
      else if (typeof payload.role === 'object') {
        role = payload.role.nama_role || payload.role.name || null
        roleId = payload.role.id ?? null
      }
    }

    if (!role && (payload.role_id || payload.roleId)) {
      roleId = payload.role_id ?? payload.roleId
    }

    // jika ada roleId, coba mapping lewat roles.value
    if (!role && roleId != null && Array.isArray(roles.value)) {
      const found = roles.value.find(rr => Number(rr.id) === Number(roleId))
      if (found) role = found.nama_role || found.name || null
    }

    return { name, role }
  } catch (e) {
    return null
  }
}

const fetchRoles = async () => {
  try {
    const token = localStorage.getItem('token') || sessionStorage.getItem('token') || null
    const res = await axios.get('http://localhost:8080/api/roles', {
      headers: token ? getAuthHeaders(token) : undefined
    })
    let arr = []
    if (Array.isArray(res.data)) arr = res.data
    else if (Array.isArray(res.data.roles)) arr = res.data.roles
    else if (Array.isArray(res.data.data)) arr = res.data.data
    else if (Array.isArray(res.data.result)) arr = res.data.result
    // normalisasi
    roles.value = arr.map(r => ({ id: r.id, nama_role: r.nama_role ?? r.name ?? r.nama ?? '' }))
    // debug: hapus jika tidak perlu
    // console.debug('roles loaded', roles.value)
  } catch (err) {
    console.warn('Gagal ambil roles:', err)
    roles.value = []
  }
}

const fetchProfile = async () => {
  const token = localStorage.getItem('token') || sessionStorage.getItem('token') || null
  if (!token) {
    router.push('/login')
    return
  }

  const endpoints = ['/profile']
  let found = null

  for (const ep of endpoints) {
    try {
      const res = await axios.get(`http://localhost:8080/api${ep}`, {
        headers: getAuthHeaders(token),
      })
      // debug: lihat dulu struktur respons jika role belum muncul
      // console.debug('profile', ep, res.data)
      const candidate = extractFromResponse(res.data)
      if (candidate) {
        found = candidate
        break
      }
    } catch (err) {
      // abaikan — coba endpoint lainnya
    }
  }

  // jika belum ditemukan, coba GET /users (beberapa backend mengembalikan user di index 0)
  if (!found) {
    try {
      const res = await axios.get('http://localhost:8080/api/users', {
        headers: getAuthHeaders(token),
      })
      // console.debug('users index', res.data)
      if (Array.isArray(res.data) && res.data.length > 0) {
        const candidate = extractFromResponse(res.data[0])
        if (candidate) found = candidate
      } else {
        const candidate = extractFromResponse(res.data)
        if (candidate) found = candidate
      }
    } catch (err) {
      // ignore
    }
  }

  // fallback decode jwt
  if (!found) {
    const decoded = tryDecodeJwt(token)
    if (decoded && (decoded.name || decoded.role)) found = decoded
  }

  if (found) {
    if (found.name) userName.value = found.name
    if (found.role) userRole.value = found.role
  } else {
    userName.value = 'Pengguna'
    userRole.value = '—'
    notifyError('Gagal mengambil profil lengkap. Jika perlu, cek endpoint /profile atau /me di backend.')
  }
}

// handle logout
const handleLogout = async () => {
  try {
    const token = localStorage.getItem('token') || sessionStorage.getItem('token')

    await axios.post(
      'http://localhost:8080/api/logout',
      {},
      { headers: { Authorization: `Bearer ${token}` } }
    )

    localStorage.removeItem('token')
    sessionStorage.removeItem('token')

    notifySuccess('Logout berhasil, sampai jumpa!')
    setTimeout(() => router.push('/login'), 1500)
  } catch (error) {
    console.error('Logout gagal:', error)
    notifyError('Logout gagal, silakan coba lagi.')
  }
}

onMounted(async () => {
  await fetchRoles()    // ambil daftar role dulu supaya bisa mapping role_id -> nama_role
  await fetchProfile()  // lalu ambil profil user dan map ke nama role jika perlu
})
</script>

<template>
  <VBadge
    dot
    location="bottom right"
    offset-x="3"
    offset-y="3"
    color="success"
    bordered
  >
    <VAvatar
      class="cursor-pointer"
      color="primary"
      variant="tonal"
    >
      <VImg :src="avatar1" />

      <!-- SECTION Menu -->
      <VMenu
        activator="parent"
        width="230"
        location="bottom end"
        offset="14px"
      >
        <VList>
          <!-- 👉 User Avatar & Name -->
          <VListItem>
            <template #prepend>
              <VListItemAction start>
                <VBadge
                  dot
                  location="bottom right"
                  offset-x="3"
                  offset-y="3"
                  color="success"
                >
                  <VAvatar
                    color="primary"
                    variant="tonal"
                  >
                    <VImg :src="avatar1" />
                  </VAvatar>
                </VBadge>
              </VListItemAction>
            </template>

            <VListItemTitle class="font-weight-semibold">
              {{ userName }}
            </VListItemTitle>
            <VListItemSubtitle>{{ userRole }}</VListItemSubtitle>
          </VListItem>
          <VDivider class="my-2" />

          <!-- 👉 Logout -->
          <VListItem to="/login" @click="handleLogout">
            <template #prepend>
              <VIcon
                class="me-2"
                icon="ri-logout-box-r-line"
                size="22"
              />
            </template>

            <VListItemTitle>Logout</VListItemTitle>
          </VListItem>
        </VList>
      </VMenu>
      <!-- !SECTION -->
    </VAvatar>
  </VBadge>
</template>
