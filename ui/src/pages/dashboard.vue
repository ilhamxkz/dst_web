<script setup>
import { ref, onMounted, reactive } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

// CSS
import 'bootstrap/dist/css/bootstrap.min.css'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'

// DataTables
import DataTable from 'datatables.net-vue3'
import DataTablesCore from 'datatables.net-bs5'
import Buttons from 'datatables.net-buttons-bs5'
DataTable.use(DataTablesCore)
DataTable.use(Buttons)

const router = useRouter()

// axios instance
const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

// state
const users = ref([])
const roles = ref([])
const loading = ref(true)
const error = ref(null)

// Custom alert/confirm modal
const alertModal = reactive({
  show: false,
  type: 'alert',
  title: '',
  message: '',
  onConfirm: null,
  onCancel: null
})

const showAlert = (message, title = 'Peringatan') => {
  alertModal.show = true
  alertModal.type = 'alert'
  alertModal.title = title
  alertModal.message = message
  alertModal.onConfirm = () => { alertModal.show = false }
}

const showConfirm = (message, title = 'Konfirmasi') => {
  return new Promise((resolve) => {
    alertModal.show = true
    alertModal.type = 'confirm'
    alertModal.title = title
    alertModal.message = message
    alertModal.onConfirm = () => { alertModal.show = false; resolve(true) }
    alertModal.onCancel = () => { alertModal.show = false; resolve(false) }
  })
}

// error handler
const handleError = (err) => {
  console.error('handleError ->', err)
  const status = err?.response?.status
  if (status === 401) {
    error.value = 'Silahkan login terlebih dahulu'
  } else if (err?.response?.data) {
    error.value = err.response.data.message || JSON.stringify(err.response.data) || err.message
  } else {
    error.value = err?.message || 'Gagal ambil data'
  }
}

// Columns (tambah kolom Role)
const userColumns = [
  { title: 'ID', data: 'id' },
  { title: 'Username', data: 'username' },
  { title: 'Email', data: 'email' },
  { title: 'Role', data: 'role_name' },
  {
    title: 'Aksi',
    data: null,
    render: () => `
      <button class='btn btn-warning btn-sm me-1 edit-btn'>Edit</button>
      <button class='btn btn-danger btn-sm delete-btn'>Hapus</button>
    `
  }
]

// normalize user -> sertakan role_id & role_name
const normalizeUser = (u) => {
  const roleId = u.role_id ?? u.role?.id ?? u.roleId ?? null
  let roleName = u.nama_role ?? u.role?.nama_role ?? u.role?.name ?? u.role_name ?? ''
  if ((!roleName || roleName === '') && roleId != null && Array.isArray(roles.value)) {
    const r = roles.value.find(rr => Number(rr.id) === Number(roleId))
    if (r) roleName = r.nama_role ?? r.name ?? ''
  }
  return {
    id: u.id ?? u.ID ?? u.user_id ?? (u.user && u.user.id) ?? null,
    username: u.username ?? u.Username ?? u.name ?? u.nama ?? '',
    email: u.email ?? u.Email ?? u.email_address ?? '',
    role_id: roleId,
    role_name: roleName || ''
  }
}

// fetch roles
const fetchRoles = async () => {
  try {
    const res = await api.get('/roles')
    let arr = []
    if (Array.isArray(res.data)) arr = res.data
    else if (Array.isArray(res.data.roles)) arr = res.data.roles
    else if (Array.isArray(res.data.data)) arr = res.data.data
    else if (Array.isArray(res.data.result)) arr = res.data.result
    roles.value = arr.map(r => ({ id: r.id, nama_role: r.nama_role ?? r.name ?? r.nama ?? '' }))
  } catch (err) {
    console.warn('Gagal ambil roles:', err)
  }
}

// fetch users
const fetchUsers = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/users')
    let arr = []
    if (Array.isArray(res.data)) arr = res.data
    else if (Array.isArray(res.data.users)) arr = res.data.users
    else if (Array.isArray(res.data.data)) arr = res.data.data
    else if (Array.isArray(res.data.result)) arr = res.data.result
    users.value = arr.map(normalizeUser)
  } catch (err) {
    handleError(err)
  } finally {
    loading.value = false
  }
}

// Modal CRUD
const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({ id: null, username: '', email: '', password: '', role_id: null })

const openCreate = () => {
  modalMode.value = 'create'
  Object.assign(form, {
    id: null,
    username: '',
    email: '',
    password: '',
    role_id: roles.value.length ? roles.value[0].id : null
  })
  showModal.value = true
}

const openEdit = (row) => {
  modalMode.value = 'edit'
  Object.assign(form, {
    id: row.id,
    username: row.username || '',
    email: row.email || '',
    password: '',
    role_id: row.role_id ?? null
  })
  showModal.value = true
}

const submit = async () => {
  error.value = null
  try {
    if (!form.username || !form.email || (modalMode.value === 'create' && !form.password)) {
      showAlert('Harap isi username, email, dan password (untuk create).')
      return
    }
    if (modalMode.value === 'create') {
      await api.post('/users', { username: form.username, email: form.email, password: form.password, role_id: form.role_id })
    } else {
      const payload = { username: form.username, email: form.email, role_id: form.role_id }
      if (form.password) payload.password = form.password
      await api.put(`/users/${form.id}`, payload)
    }
    showModal.value = false
    await fetchUsers()
    showAlert('Data berhasil disimpan!', 'Berhasil')
  } catch (err) {
    showAlert('Error: ' + (err.response?.data?.message || err.message), 'Error')
    handleError(err)
  }
}

const deleteUser = async (id) => {
  const confirmed = await showConfirm('Yakin ingin hapus user ini?', 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/users/${id}`)
    await fetchUsers()
    showAlert('User berhasil dihapus!', 'Berhasil')
  } catch (err) {
    showAlert('Gagal hapus: ' + (err.response?.data?.message || err.message), 'Error')
    handleError(err)
  }
}

onMounted(async () => {
  await fetchRoles()
  await fetchUsers()
})
</script>

<template>
  <div class="container mt-4">
    <!-- DASHBOARD USER -->
    <div class="d-flex justify-content-between mb-3">
      <h3>Data User</h3>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah User</button>
        <button class="btn btn-secondary" @click="fetchUsers">Refresh</button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-4">
      <div class="spinner-border text-primary" role="status"></div>
      <p class="mt-2">Sedang memuat data...</p>
    </div>

    <div v-else-if="error" class="alert alert-danger text-center">
      {{ error }}
      <div v-if="error && error.toLowerCase().includes('login')" class="mt-2">
        <button class="btn btn-sm btn-outline-light" @click="router.push('/login')">Login Sekarang</button>
      </div>
    </div>

    <div v-else>
      <DataTable
        class="table table-striped table-hover table-bordered align-middle text-center w-100"
        :data="users"
        :columns="userColumns"
        :options="{
          dom: 'lBfrtip',
          buttons: ['copy', 'csv', 'excel', 'pdf', 'print'],
          pageLength: 10,
          lengthMenu: [[10, 25, 50], [10, 25, 50]],
          responsive: true,
          createdRow: function (row, data) {
            const editBtn = row.querySelector('.edit-btn')
            const delBtn = row.querySelector('.delete-btn')
            if (editBtn) editBtn.addEventListener('click', () => { openEdit(data) })
            if (delBtn) delBtn.addEventListener('click', () => { deleteUser(data.id) })
          }
        }"
      />
    </div>

    <!-- SIMPLE MODAL -->
    <div v-if="showModal" class="modal-backdrop" style="position:fixed;inset:0;display:flex;align-items:center;justify-content:center;z-index:1050;">
      <div class="card p-3" style="width:420px;">
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h5 class="m-0">{{ modalMode === 'create' ? 'Tambah User' : 'Edit User' }}</h5>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal = false">×</button>
        </div>

        <div class="mb-2">
          <label class="form-label">Username</label>
          <input class="form-control" v-model="form.username" />
        </div>
        <div class="mb-2">
          <label class="form-label">Email</label>
          <input class="form-control" v-model="form.email" />
        </div>
        <div class="mb-2">
          <label class="form-label">Role</label>
          <select class="form-select" v-model="form.role_id">
            <option :value="null">-- Pilih Role --</option>
            <option v-for="r in roles" :key="r.id" :value="r.id">{{ r.nama_role }}</option>
          </select>
        </div>
        <div class="mb-3">
          <label class="form-label">Password <small v-if="modalMode==='edit'">(kosongkan jika tidak ganti)</small></label>
          <input type="password" class="form-control" v-model="form.password" />
        </div>

        <div class="d-flex justify-content-end">
          <button class="btn btn-secondary me-2" @click="showModal=false">Batal</button>
          <button class="btn btn-primary" @click="submit">{{ modalMode==='create' ? 'Simpan' : 'Update' }}</button>
        </div>
      </div>
    </div>

    <!-- Custom Alert/Confirm Modal -->
    <div v-if="alertModal.show" class="modal-backdrop custom-alert-modal"
         style="position: fixed; inset: 0; display: flex; align-items: center; justify-content: center; z-index: 1060">
      <div class="alert-card">
        <div class="alert-header"><h5 class="alert-title">{{ alertModal.title }}</h5></div>
        <div class="alert-body"><p class="alert-message">{{ alertModal.message }}</p></div>
        <div class="alert-footer">
          <template v-if="alertModal.type==='confirm'">
            <button class="btn btn-secondary me-2" @click="alertModal.onCancel">Cancel</button>
            <button class="btn btn-primary" @click="alertModal.onConfirm">OK</button>
          </template>
          <template v-else>
            <button class="btn btn-primary" @click="alertModal.onConfirm">OK</button>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
table.dataTable th {
  background-color: #f8f9fa;
  font-weight: 600;
  text-align: center;
}
table.dataTable td {
  vertical-align: middle;
  text-align: center;
}
table.dataTable tbody tr:hover {
  background-color: #f1f1f1;
  transition: background-color 0.2s;
}
.modal-backdrop { background: rgba(0,0,0,0.35); }

/* Custom Alert Modal */
.custom-alert-modal { background: rgba(0,0,0,0.5); }
.alert-card { background:white; border-radius:12px; box-shadow:0 20px 25px -5px rgba(0,0,0,0.1),0 10px 10px -5px rgba(0,0,0,0.04); width:400px; max-width:90vw; overflow:hidden; }
.alert-header { padding:20px 24px 0; border-bottom:none }
.alert-title { margin:0; font-size:1.1rem; font-weight:600; color:#1f2937 }
.alert-body { padding:16px 24px }
.alert-message { margin:0; color:#6b7280; line-height:1.5 }
.alert-footer { padding:0 24px 24px; display:flex; justify-content:flex-end; gap:8px }
.alert-footer .btn { padding:8px 16px; border-radius:6px; font-size:0.875rem; font-weight:500 }
</style>
