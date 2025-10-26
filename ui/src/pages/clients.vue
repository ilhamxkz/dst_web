<script setup>
import { ref, reactive, onMounted } from 'vue'
import axios from 'axios'

const clients = ref([])
const loading = ref(false)
const error = ref(null)

const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

// === Custom Alert/Confirm State ===
const alertModal = reactive({
  show: false,
  type: 'alert', // 'alert' | 'confirm'
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
  alertModal.onConfirm = () => {
    alertModal.show = false
  }
}

const showConfirm = (message, title = 'Konfirmasi') => {
  return new Promise(resolve => {
    alertModal.show = true
    alertModal.type = 'confirm'
    alertModal.title = title
    alertModal.message = message
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
// ==================================

const fetchClients = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/clients')
    clients.value = res.data.clients || res.data || []
  } catch (err) {
    error.value = err.response?.data?.error || err.message
  } finally {
    loading.value = false
  }
}

const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({ id: null, name: '', website: '', logo: null })

const openCreate = () => {
  modalMode.value = 'create'
  form.id = null
  form.name = ''
  form.website = ''
  form.logo = null
  showModal.value = true
}

const startEdit = row => {
  modalMode.value = 'edit'
  form.id = row.id
  form.name = row.name
  form.website = row.website
  form.logo = null
  showModal.value = true
}

const submit = async () => {
  if (!form.name) {
    showAlert('Nama wajib diisi', 'Validasi')
    return
  }
  const fd = new FormData()
  fd.append('name', form.name)
  fd.append('website', form.website)
  if (form.logo) fd.append('logo', form.logo)

  try {
    if (modalMode.value === 'create') {
      await api.post('/clients', fd, { headers: { 'Content-Type': 'multipart/form-data' } })
    } else {
      await api.put(`/clients/${form.id}`, fd, { headers: { 'Content-Type': 'multipart/form-data' } })
    }
    showModal.value = false
    await fetchClients()
    showAlert('Data client berhasil disimpan!', 'Berhasil')
  } catch (err) {
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

const deleteClient = async row => {
  const confirmed = await showConfirm(`Yakin hapus client ${row.name}?`, 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/clients/${row.id}`)
    await fetchClients()
    showAlert('Client berhasil dihapus!', 'Berhasil')
  } catch (err) {
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

onMounted(fetchClients)
</script>

<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between mb-3">
      <h3>Data Client</h3>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah Client</button>
        <button class="btn btn-secondary" @click="fetchClients">Refresh</button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-4">
      <div class="spinner-border text-primary" role="status"></div>
      <p class="mt-2">Memuat data...</p>
    </div>

    <div v-else-if="error" class="alert alert-danger text-center">
      {{ error }}
    </div>

    <div v-else class="row g-3">
      <div v-for="c in clients" :key="c.id" class="col-md-3 col-sm-6">
        <div class="card text-center h-100">
          <img v-if="c.logo" :src="`http://localhost:8080${c.logo}`" class="card-img-top p-3" style="max-height:80px; object-fit:contain" />
          <div class="card-body">
            <h6 class="card-title">{{ c.name }}</h6>
            <a :href="c.website" target="_blank" rel="noopener" class="d-block text-primary small">
              {{ c.website }}
            </a>
          </div>
          <div class="card-footer d-flex justify-content-center">
            <button class="btn btn-warning btn-sm me-2" @click="startEdit(c)">Edit</button>
            <button class="btn btn-danger btn-sm" @click="deleteClient(c)">Hapus</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <div v-if="showModal" class="modal-backdrop" style="position: fixed; inset: 0; display: flex; align-items: center; justify-content: center; z-index: 1050">
      <div class="card p-3" style="width: 420px; max-width: 95vw">
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h5 class="m-0">{{ modalMode === 'create' ? 'Tambah Client' : 'Edit Client' }}</h5>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal = false">×</button>
        </div>
        <div class="row g-2">
          <div class="col-12">
            <label class="form-label">Nama Client</label>
            <input class="form-control" v-model="form.name" />
          </div>
          <div class="col-12">
            <label class="form-label">Website</label>
            <input class="form-control" v-model="form.website" />
          </div>
          <div class="col-12">
            <label class="form-label">Logo</label>
            <input type="file" class="form-control" @change="e => form.logo = e.target.files[0]" />
          </div>
        </div>
        <div class="d-flex justify-content-end mt-3">
          <button class="btn btn-secondary me-2" @click="showModal = false">Batal</button>
          <button class="btn btn-primary" @click="submit">{{ modalMode === 'create' ? 'Simpan' : 'Update' }}</button>
        </div>
      </div>
    </div>

    <!-- Custom Alert/Confirm Modal -->
    <div
      v-if="alertModal.show"
      class="modal-backdrop custom-alert-modal"
      style="position: fixed; inset: 0; display: flex; align-items: center; justify-content: center; z-index: 1060"
    >
      <div class="alert-card">
        <div class="alert-header">
          <h5 class="alert-title">{{ alertModal.title }}</h5>
        </div>
        <div class="alert-body">
          <p class="alert-message">{{ alertModal.message }}</p>
        </div>
        <div class="alert-footer">
          <template v-if="alertModal.type === 'confirm'">
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
.modal-backdrop {
  background: rgba(0, 0, 0, 0.35);
}
.custom-alert-modal {
  background: rgba(0, 0, 0, 0.5);
}
.alert-card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 20px 25px -5px rgba(0,0,0,.1), 0 10px 10px -5px rgba(0,0,0,.04);
  width: 380px;
  max-width: 90vw;
  overflow: hidden;
  animation: slideIn 0.2s ease-out;
}
.alert-header { padding: 16px 20px 0 }
.alert-title { margin: 0; font-size: 1.1rem; font-weight: 600; color: #1f2937 }
.alert-body { padding: 12px 20px }
.alert-message { margin: 0; color: #6b7280; line-height: 1.5 }
.alert-footer { padding: 0 20px 16px; display: flex; justify-content: flex-end; gap: 8px }
.alert-footer .btn { padding: 6px 14px; border-radius: 6px; font-size: .875rem; font-weight: 500 }
@keyframes slideIn {
  from { opacity: 0; transform: translateY(-16px) }
  to   { opacity: 1; transform: translateY(0) }
}
</style>
