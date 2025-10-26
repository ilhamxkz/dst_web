<script setup>
import axios from 'axios'
import { onMounted, reactive, ref } from 'vue'

// Bootstrap + DataTables
import 'bootstrap/dist/css/bootstrap.min.css'
import DataTablesCore from 'datatables.net-bs5'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import Buttons from 'datatables.net-buttons-bs5'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'
import DataTable from 'datatables.net-vue3'

DataTable.use(DataTablesCore)
DataTable.use(Buttons)

// state
const images = ref([])
const projects = ref([])
const loading = ref(true)
const error = ref(null)

// axios client
const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

// --- Custom Alert/Confirm State & Helpers ---
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
// ---------------------------------------------

// ambil data
const fetchAll = async () => {
  loading.value = true
  error.value = null
  try {
    const [rImages, rProjects] = await Promise.all([
      api.get('/project-images'),
      api.get('/projects')
    ])
    images.value = Array.isArray(rImages.data.project_images) ? rImages.data.project_images : (Array.isArray(rImages.data) ? rImages.data : [])
    projects.value = Array.isArray(rProjects.data.projects) ? rProjects.data.projects : (Array.isArray(rProjects.data) ? rProjects.data : [])
  } catch (err) {
    error.value = err?.response?.data?.message || err.message
    if (err?.response?.status === 401) error.value = 'Silahkan login terlebih dahulu'
  } finally {
    loading.value = false
  }
}

// modal form
const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({ id: null, project_id: null, image_url: '', caption: '' })

const openCreate = () => {
  modalMode.value = 'create'
  form.id = null
  form.project_id = null
  form.image_url = ''
  form.caption = ''
  showModal.value = true
}
const openEdit = (row) => {
  modalMode.value = 'edit'
  form.id = row.id
  form.project_id = row.project_id
  form.image_url = row.image_url
  form.caption = row.caption
  showModal.value = true
}

const submit = async () => {
  try {
    if (modalMode.value === 'create') {
      await api.post('/project-images', { project_id: form.project_id, image_url: form.image_url, caption: form.caption })
      showModal.value = false
      await fetchAll()
      showAlert('Gambar berhasil ditambahkan!', 'Berhasil')
    } else {
      await api.put(`/project-images/${form.id}`, { project_id: form.project_id, image_url: form.image_url, caption: form.caption })
      showModal.value = false
      await fetchAll()
      showAlert('Gambar berhasil diperbarui!', 'Berhasil')
    }
  } catch (err) {
    showAlert(err?.response?.data?.message || err.message, 'Error')
  }
}

const deleteImage = async (id) => {
  const confirmed = await showConfirm('Hapus gambar ini?', 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/project-images/${id}`)
    await fetchAll()
    showAlert('Gambar berhasil dihapus!', 'Berhasil')
  } catch (err) {
    showAlert(err?.response?.data?.message || err.message, 'Error')
  }
}

onMounted(fetchAll)
</script>

<template>
  <div class="container mt-3">
    <div class="d-flex justify-content-between align-items-center mb-3">
      <h4>Data Project Images</h4>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah Gambar</button>
        <button class="btn btn-secondary" @click="fetchAll">Refresh</button>
      </div>
    </div>

    <div v-if="loading">Loading...</div>
    <div v-else-if="error" class="alert alert-danger">{{ error }}</div>
    <div v-else>
      <DataTable
        class="table table-striped table-hover table-bordered align-middle text-center w-100"
        :data="images"
        :columns="[
          { title: 'ID', data: 'id' },
          { title: 'Project', data: row => (row.project?.title || row.project_id) },
          { title: 'Image URL', data: row => `<a href='${row.image_url}' target='_blank' rel='noopener'>${row.image_url}</a>` },
          { title: 'Caption', data: 'caption' },
          { 
            title: 'Aksi', 
            data: null, 
            render: (d,t,row)=>`
              <button class='btn btn-warning btn-sm me-1 edit-btn'>Edit</button>
              <button class='btn btn-danger btn-sm delete-btn'>Hapus</button>
            `
          }
        ]"
        :options="{
          dom: 'lBfrtip',
          buttons: ['copy','csv','excel','pdf','print'],
          responsive: true,
          pageLength: 5,
          lengthMenu: [[5,10,25,50,100],[5,10,25,50,100]],
          createdRow: function(row,data) {
            const e = row.querySelector('.edit-btn'), del = row.querySelector('.delete-btn')
            if (e) e.addEventListener('click', ()=> openEdit(data))
            if (del) del.addEventListener('click', ()=> deleteImage(data.id))
          }
        }"
      />
    </div>

    <!-- modal -->
    <div v-if="showModal" class="modal-backdrop d-flex align-items-center justify-content-center">
      <div class="card p-3" style="width:420px;">
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h6 class="m-0">{{ modalMode==='create' ? 'Tambah Gambar' : 'Edit Gambar' }}</h6>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal=false">×</button>
        </div>
        <div class="mb-2">
          <label class="form-label">Project</label>
          <select class="form-select" v-model="form.project_id">
            <option :value="null">-- pilih project --</option>
            <option v-for="p in projects" :key="p.id" :value="p.id">{{ p.title }}</option>
          </select>
        </div>
        <div class="mb-2">
          <label class="form-label">Image URL</label>
          <input type="text" class="form-control" v-model="form.image_url" />
        </div>
        <div class="mb-3">
          <label class="form-label">Caption</label>
          <input type="text" class="form-control" v-model="form.caption" />
        </div>
        <div class="d-flex justify-content-end">
          <button class="btn btn-secondary me-2" @click="showModal=false">Batal</button>
          <button class="btn btn-primary" @click="submit">{{ modalMode==='create' ? 'Simpan' : 'Update' }}</button>
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
          <p class="alert-message" v-html="alertModal.message"></p>
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
table.dataTable th { background:#f5f5f5; font-weight:600; text-align:center }
.modal-backdrop { position:fixed; inset:0; background:rgba(0,0,0,0.35); z-index:1050 }

/* Custom alert modal styles */
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
  animation: slideIn 0.18s ease-out;
}
.alert-header { padding: 16px 20px 0 }
.alert-title { margin: 0; font-size: 1.05rem; font-weight: 600; color: #111827 }
.alert-body { padding: 12px 20px }
.alert-message { margin: 0; color: #6b7280; line-height: 1.5; word-break: break-word }
.alert-footer { padding: 0 20px 16px; display: flex; justify-content: flex-end; gap: 8px }
.alert-footer .btn { padding: 6px 14px; border-radius: 6px; font-size: .875rem; font-weight: 500 }
@keyframes slideIn {
  from { opacity: 0; transform: translateY(-12px) }
  to   { opacity: 1; transform: translateY(0) }
}
</style>
