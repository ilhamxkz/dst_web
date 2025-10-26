<script setup>
import axios from 'axios'
import { onMounted, reactive, ref } from 'vue'

// ====== Import Bootstrap & DataTables ======
import 'bootstrap/dist/css/bootstrap.min.css'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'

import DataTablesCore from 'datatables.net-bs5'
import Buttons from 'datatables.net-buttons-bs5'
import DataTable from 'datatables.net-vue3'
DataTable.use(DataTablesCore)
DataTable.use(Buttons)

// ====== State ======
const profiles = ref([])
const loading = ref(true)
const error = ref(null)
const showModal = ref(false)
const modalMode = ref('create')

// ====== Alert/Confirm Modal ======
const alertModal = reactive({
  show: false,
  type: 'alert',
  title: '',
  message: '',
  onConfirm: null,
  onCancel: null
})

const showAlert = (msg, title = 'Peringatan') => {
  alertModal.show = true
  alertModal.type = 'alert'
  alertModal.title = title
  alertModal.message = msg
  alertModal.onConfirm = () => (alertModal.show = false)
}

const showConfirm = (msg, title = 'Konfirmasi') => {
  return new Promise(resolve => {
    alertModal.show = true
    alertModal.type = 'confirm'
    alertModal.title = title
    alertModal.message = msg
    alertModal.onConfirm = () => {
      alertModal.show = false
      resolve(true)
    }
    alertModal.onCancel = () => {
      alertModal.show = false
      resolve(false)
    }
  })
}

// ====== API Config ======
const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

// ====== DataTable Columns ======
const profileColumns = [
  { title: 'ID', data: 'id' },
  { title: 'Nama Perusahaan', data: 'company_name' },
  { title: 'Tagline', data: 'tagline' },
  { title: 'Deskripsi', data: 'description' },
  { title: 'Alamat', data: 'address' },
  { title: 'Telepon', data: 'phone' },
  { title: 'Email', data: 'email' },
  { title: 'Website', data: 'website' },
  {
    title: 'Logo',
    data: 'logo',
    render: data =>
      data
        ? `<img src="${data}" alt="logo" style="height:40px; border-radius:4px;">`
        : '-'
  },
  {
    title: 'Aksi',
    data: null,
    render: () =>
      `<button class='btn btn-warning btn-sm me-1 edit-btn'>Edit</button>
       <button class='btn btn-danger btn-sm delete-btn'>Hapus</button>`
  }
]

// ====== CRUD State ======
const form = reactive({
  id: null,
  company_name: '',
  tagline: '',
  description: '',
  address: '',
  phone: '',
  email: '',
  website: '',
  logo: ''
})

const normalizeProfile = p => ({
  id: p.id ?? null,
  company_name: p.company_name ?? '',
  tagline: p.tagline ?? '',
  description: p.description ?? '',
  address: p.address ?? '',
  phone: p.phone ?? '',
  email: p.email ?? '',
  website: p.website ?? '',
  logo: p.logo ?? ''
})

// ====== Fetch Data ======
const fetchProfiles = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/company_profile')
    // bisa 'data' atau 'data.data' tergantung API kamu
    const arr =
      Array.isArray(res.data?.data) ? res.data.data
      : Array.isArray(res.data) ? res.data
      : []
    profiles.value = arr.map(normalizeProfile)
  } catch (err) {
    error.value = err.response?.data?.message || err.message
  } finally {
    loading.value = false
  }
}

// ====== CRUD Actions ======
const openCreate = () => {
  modalMode.value = 'create'
  Object.assign(form, {
    id: null,
    company_name: '',
    tagline: '',
    description: '',
    address: '',
    phone: '',
    email: '',
    website: '',
    logo: ''
  })
  showModal.value = true
}

const openEdit = row => {
  modalMode.value = 'edit'
  Object.assign(form, row)
  showModal.value = true
}

const submit = async () => {
  if (!form.company_name || !form.address) {
    showAlert('Nama perusahaan dan alamat wajib diisi', 'Validasi')
    return
  }

  try {
    if (modalMode.value === 'create') {
      await api.post('/company_profile', form)
    } else {
      await api.put(`/company_profile/${form.id}`, form)
    }
    showModal.value = false
    await fetchProfiles()
    showAlert('✅ Data berhasil disimpan', 'Berhasil')
  } catch (err) {
    showAlert(err.response?.data?.message || err.message, 'Error')
  }
}

const deleteProfile = async id => {
  const confirm = await showConfirm('Yakin hapus data ini?', 'Konfirmasi')
  if (!confirm) return
  try {
    await api.delete(`/company_profile/${id}`)
    await fetchProfiles()
    showAlert('🗑️ Data berhasil dihapus', 'Berhasil')
  } catch (err) {
    showAlert(err.response?.data?.message || err.message, 'Error')
  }
}

onMounted(fetchProfiles)
</script>

<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between mb-3">
      <h3>Company Profile</h3>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah</button>
        <button class="btn btn-secondary" @click="fetchProfiles">Refresh</button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-5">
      <div class="spinner-border text-primary" role="status"></div>
      <p class="mt-2">Memuat data...</p>
    </div>

    <div v-else-if="error" class="alert alert-danger">{{ error }}</div>

    <div v-else>
      <DataTable
        class="table table-striped table-bordered align-middle text-center w-100"
        :data="profiles"
        :columns="profileColumns"
        :options="{
          dom: 'lBfrtip',
          buttons: ['copy', 'csv', 'excel', 'pdf', 'print'],
          pageLength: 10,
          lengthMenu: [[10, 25, 50], [10, 25, 50]],
          responsive: true,
          createdRow: (row, data) => {
            const editBtn = row.querySelector('.edit-btn')
            const delBtn = row.querySelector('.delete-btn')
            if (editBtn) editBtn.addEventListener('click', () => openEdit(data))
            if (delBtn) delBtn.addEventListener('click', () => deleteProfile(data.id))
          }
        }"
      />
    </div>

    <!-- Modal -->
    <div
      v-if="showModal"
      class="modal-backdrop"
      style="position:fixed;inset:0;display:flex;align-items:center;justify-content:center;z-index:1050"
    >
      <div class="card p-3" style="width:520px;max-width:95vw;">
        <div class="d-flex justify-content-between align-items-center mb-3">
          <h5 class="m-0">{{ modalMode === 'create' ? 'Tambah Profil' : 'Edit Profil' }}</h5>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal=false">×</button>
        </div>

        <div class="row g-2">
          <div class="col-12">
            <label class="form-label">Nama Perusahaan</label>
            <input v-model="form.company_name" class="form-control" />
          </div>
          <div class="col-12">
            <label class="form-label">Tagline</label>
            <input v-model="form.tagline" class="form-control" />
          </div>
          <div class="col-12">
            <label class="form-label">Deskripsi</label>
            <textarea v-model="form.description" class="form-control" rows="3"></textarea>
          </div>
          <div class="col-12">
            <label class="form-label">Alamat</label>
            <textarea v-model="form.address" class="form-control" rows="2"></textarea>
          </div>
          <div class="col-md-6">
            <label class="form-label">Telepon</label>
            <input v-model="form.phone" class="form-control" />
          </div>
          <div class="col-md-6">
            <label class="form-label">Email</label>
            <input v-model="form.email" class="form-control" />
          </div>
          <div class="col-12">
            <label class="form-label">Website</label>
            <input v-model="form.website" class="form-control" />
          </div>
          <div class="col-12">
            <label class="form-label">Logo URL</label>
            <input v-model="form.logo" class="form-control" placeholder="https://..." />
          </div>
        </div>

        <div class="d-flex justify-content-end mt-3">
          <button class="btn btn-secondary me-2" @click="showModal=false">Batal</button>
          <button class="btn btn-primary" @click="submit">
            {{ modalMode === 'create' ? 'Simpan' : 'Update' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Custom Alert/Confirm -->
    <div
      v-if="alertModal.show"
      class="modal-backdrop"
      style="position:fixed;inset:0;display:flex;align-items:center;justify-content:center;z-index:1060;background:rgba(0,0,0,0.4)"
    >
      <div class="alert-card bg-white rounded-3 p-4 shadow-lg" style="width:400px;max-width:90vw;">
        <h5 class="fw-bold mb-3">{{ alertModal.title }}</h5>
        <p>{{ alertModal.message }}</p>
        <div class="d-flex justify-content-end mt-3">
          <template v-if="alertModal.type === 'confirm'">
            <button class="btn btn-secondary me-2" @click="alertModal.onCancel">Batal</button>
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
.modal-backdrop {
  background: rgba(0, 0, 0, 0.3);
}
table.dataTable tbody tr:hover {
  background-color: #f1f1f1;
  transition: background-color 0.2s;
}
.modal-backdrop { background: rgba(0,0,0,0.35); }
</style>
