<script setup>
import axios from 'axios'
import 'bootstrap/dist/css/bootstrap.min.css'
import DataTablesCore from 'datatables.net-bs5'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import Buttons from 'datatables.net-buttons-bs5'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'
import DataTable from 'datatables.net-vue3'
import { onMounted, reactive, ref } from 'vue'

DataTable.use(DataTablesCore)
DataTable.use(Buttons)

const roles = ref([])
const loading = ref(true)
const error = ref(null)

// axios instance
const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

const fetchRoles = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/roles')
    if (Array.isArray(res.data)) roles.value = res.data
    else if (Array.isArray(res.data.roles)) roles.value = res.data.roles
    else if (Array.isArray(res.data.data)) roles.value = res.data.data
    else roles.value = []
  } catch (err) {
    if (err?.response?.status === 401) error.value = 'Silahkan login terlebih dahulu'
    else error.value = err?.response?.data?.message || err.message
  } finally { loading.value = false }
}

/* ----------------------------
   Modal CRUD + Custom Alert
   ---------------------------- */
const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({ id: null, nama_role: '' })

// Custom alert/confirm modal state
const alertModal = reactive({
  show: false,
  type: 'alert', // 'alert' or 'confirm'
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

const openCreate = () => {
  modalMode.value = 'create'
  form.id = null
  form.nama_role = ''
  showModal.value = true
}
const openEdit = (row) => {
  modalMode.value = 'edit'
  form.id = row.id
  form.nama_role = row.nama_role || ''
  showModal.value = true
}

const submit = async () => {
  if (!form.nama_role) {
    showAlert('Nama role wajib diisi!')
    return
  }
  try {
    if (modalMode.value === 'create') await api.post('/roles', { nama_role: form.nama_role })
    else await api.put(`/roles/${form.id}`, { nama_role: form.nama_role })
    showModal.value = false
    await fetchRoles()
    showAlert('Data berhasil disimpan!', 'Berhasil')
  } catch (err) {
    if (err?.response?.status === 401) error.value = 'Silahkan login terlebih dahulu'
    else showAlert(err?.response?.data?.message || err.message, 'Error')
  }
}

const deleteRole = async (id) => {
  const confirmed = await showConfirm('Hapus role ini?', 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/roles/${id}`)
    await fetchRoles()
    showAlert('Role berhasil dihapus!', 'Berhasil')
  } catch (err) {
    if (err?.response?.status === 401) error.value = 'Silahkan login terlebih dahulu'
    else showAlert(err?.response?.data?.message || err.message, 'Error')
  }
}

// Columns
const roleColumns = [
  { title: 'ID', data: 'id' },
  { title: 'Role Name', data: 'nama_role' },
  {
    title: 'Aksi',
    data: null,
    render: () => `
      <button class='btn btn-warning btn-sm me-1 edit-btn'>Edit</button>
      <button class='btn btn-danger btn-sm delete-btn'>Hapus</button>
    `
  }
]

onMounted(fetchRoles)
</script>

<template>
  <div class="container mt-3">
    <div class="d-flex justify-content-between align-items-center mb-3">
      <h4>Data Roles</h4>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah Role</button>
        <button class="btn btn-secondary" @click="fetchRoles">Refresh</button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-4">
      <div class="spinner-border text-primary" role="status"></div>
      <p class="mt-2">Sedang memuat data...</p>
    </div>

    <div v-else-if="error" class="alert alert-danger text-center">{{ error }}</div>

    <div v-else>
      <DataTable
        class="table table-hover table-bordered align-middle text-center w-100"
        :data="roles"
        :columns="roleColumns"
        :options="{
          dom: 'lBfrtip',
          buttons: ['copy','csv','excel','pdf','print'],
          responsive: true,
          pageLength: 5,
          lengthMenu: [[5,10,25,50,100],[5,10,25,50,100]],
          createdRow: function(row,data) {
            const e = row.querySelector('.edit-btn'),
                  del = row.querySelector('.delete-btn')
            if (e) e.addEventListener('click', ()=> openEdit(data))
            if (del) del.addEventListener('click', ()=> deleteRole(data.id))
          }
        }"
      />
    </div>

    <!-- Modal -->
    <div v-if="showModal" class="modal-backdrop"
         style="position:fixed;inset:0;display:flex;align-items:center;justify-content:center;z-index:1050;">
      <div class="card p-3" style="width:360px;">
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h6 class="m-0">{{ modalMode==='create' ? 'Tambah Role' : 'Edit Role' }}</h6>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal=false">×</button>
        </div>
        <div class="mb-2">
          <label class="form-label">Nama Role</label>
          <input class="form-control" v-model="form.nama_role" />
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
table.dataTable th { background:#f5f5f5; font-weight:600;text-align:center }
table.dataTable td { text-align:center; vertical-align: middle; }
.modal-backdrop { background: rgba(0,0,0,0.35) }
.custom-alert-modal { background: rgba(0,0,0,0.5); }
.alert-card {
  background:white; border-radius:12px;
  box-shadow:0 20px 25px -5px rgba(0,0,0,0.1),0 10px 10px -5px rgba(0,0,0,0.04);
  width:400px; max-width:90vw; overflow:hidden;
}
.alert-header { padding:20px 24px 0; border-bottom:none }
.alert-title { margin:0; font-size:1.1rem; font-weight:600; color:#1f2937 }
.alert-body { padding:16px 24px }
.alert-message { margin:0; color:#6b7280; line-height:1.5 }
.alert-footer { padding:0 24px 24px; display:flex; justify-content:flex-end; gap:8px }
.alert-footer .btn { padding:8px 16px; border-radius:6px; font-size:0.875rem; font-weight:500 }
</style>
