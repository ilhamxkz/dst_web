<script setup>
import axios from 'axios'
import { onBeforeUnmount, onMounted, reactive, ref, nextTick } from 'vue'

// CKEditor manual
import ClassicEditor from '@ckeditor/ckeditor5-build-classic'

let editorInstance = null
const editorElement = ref(null)

const initEditor = async () => {
  if (!editorElement.value) return
  try {
    if (editorInstance) {
      await editorInstance.destroy()
      editorInstance = null
    }
    editorInstance = await ClassicEditor.create(editorElement.value, {
      toolbar: [
        'undo', 'redo', '|',
        'bold', 'italic', 'link', '|',
        'insertTable', '|',
        'bulletedList', 'numberedList', '|',
        'blockQuote'
      ]
    })
    editorInstance.model.document.on('change:data', () => {
      form.description = editorInstance.getData()
    })
    if (form.description) {
      editorInstance.setData(form.description)
    }
  } catch (e) {
    console.error('CKEditor init error:', e)
    showAlert('Gagal inisialisasi editor: ' + (e.message || e), 'Error')
  }
}

onBeforeUnmount(() => {
  if (editorInstance) {
    editorInstance.destroy().catch(console.error)
    editorInstance = null
  }
})

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

// ambil data projects
const fetchAll = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/projects')
    projects.value = Array.isArray(res.data.projects) ? res.data.projects : (Array.isArray(res.data) ? res.data : [])
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
const form = reactive({
  id: null,
  title: '',
  client_name: '',
  description: '',
  start_date: '',
  end_date: '',
  status: 'ongoing',
  cover_image: ''
})

const openCreate = async () => {
  modalMode.value = 'create'
  Object.assign(form, {
    id: null, title: '', client_name: '', description: '',
    start_date: '', end_date: '', status: 'ongoing', cover_image: ''
  })
  showModal.value = true
  await nextTick()
  await initEditor()
}

const openEdit = async (row) => {
  modalMode.value = 'edit'
  Object.assign(form, {
    id: row.id,
    title: row.title,
    client_name: row.client_name,
    description: row.description,
    start_date: row.start_date ? row.start_date.split('T')[0] : '',
    end_date: row.end_date ? row.end_date.split('T')[0] : '',
    status: row.status || 'ongoing',
    cover_image: row.cover_image || ''
  })
  showModal.value = true
  await nextTick()
  await initEditor()
}

const submit = async () => {
  try {
    if (editorInstance) {
      form.description = editorInstance.getData()
    }

    // Format tanggal ke ISO
    const formatDate = (d) => d ? `${d}T00:00:00Z` : '';
    const payload = {
      title: form.title,
      client_name: form.client_name,
      description: form.description,
      start_date: formatDate(form.start_date),
      end_date: formatDate(form.end_date),
      status: form.status,
      cover_image: form.cover_image
    };

    if (!payload.title) {
      showAlert('Judul project wajib diisi', 'Validasi')
      return
    }

    if (modalMode.value === 'create') {
      await api.post('/projects', payload)
      showModal.value = false
      await fetchAll()
      showAlert('Project berhasil ditambahkan!', 'Berhasil')
    } else {
      await api.put(`/projects/${form.id}`, payload)
      showModal.value = false
      await fetchAll()
      showAlert('Project berhasil diperbarui!', 'Berhasil')
    }
  } catch (err) {
    showAlert(err?.response?.data?.message || err.message, 'Error')
  }
}

const deleteProject = async (id) => {
  const confirmed = await showConfirm('Hapus project ini?', 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/projects/${id}`)
    await fetchAll()
    showAlert('Project berhasil dihapus!', 'Berhasil')
  } catch (err) {
    showAlert(err?.response?.data?.message || err.message, 'Error')
  }
}

onMounted(fetchAll)
</script>

<template>
  <div class="container mt-3">
    <div class="d-flex justify-content-between align-items-center mb-3">
      <h4>Data Projects</h4>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah Project</button>
        <button class="btn btn-secondary" @click="fetchAll">Refresh</button>
      </div>
    </div>

    <div v-if="loading">Loading...</div>
    <div v-else-if="error" class="alert alert-danger">{{ error }}</div>
    <div v-else>
      <DataTable
        class="table table-striped table-hover table-bordered align-middle text-center w-100"
        :data="projects"
        :columns="[
          { title: 'ID', data: 'id' },
          { title: 'Judul', data: 'title' },
          { title: 'Client', data: 'client_name' },
          { title: 'Status', data: 'status' },
          { title: 'Mulai', data: row => row.start_date ? row.start_date.split('T')[0] : '' },
          { title: 'Selesai', data: row => row.end_date ? row.end_date.split('T')[0] : '' },
          { 
            title: 'Aksi', 
            data: null, 
            render: (d,t,row)=>`
              <button class='btn btn-warning btn-sm me-1 edit-btn' data-id='${row.id}'>Edit</button>
              <button class='btn btn-danger btn-sm delete-btn' data-id='${row.id}'>Hapus</button>
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
            if (del) del.addEventListener('click', ()=> deleteProject(data.id))
          }
        }"
      />
    </div>

    <!-- modal -->
    <div v-if="showModal" class="modal-backdrop d-flex align-items-center justify-content-center">
      <div class="card p-3" style="width:520px;">
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h6 class="m-0">{{ modalMode==='create' ? 'Tambah Project' : 'Edit Project' }}</h6>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal=false">×</button>
        </div>

        <div class="mb-2">
          <label class="form-label">Judul</label>
          <input type="text" class="form-control" v-model="form.title" />
        </div>
        <div class="mb-2">
          <label class="form-label">Client</label>
          <input type="text" class="form-control" v-model="form.client_name" />
        </div>
        <div class="mb-2">
          <label class="form-label">Deskripsi</label>
          <div ref="editorElement" class="editor-placeholder"></div>
        </div>
        <div class="mb-2">
          <label class="form-label">Tanggal Mulai</label>
          <input type="date" class="form-control" v-model="form.start_date" />
        </div>
        <div class="mb-2">
          <label class="form-label">Tanggal Selesai</label>
          <input type="date" class="form-control" v-model="form.end_date" />
        </div>
        <div class="mb-2">
          <label class="form-label">Status</label>
          <select class="form-select" v-model="form.status">
            <option value="ongoing">Ongoing</option>
            <option value="completed">Completed</option>
            <option value="pending">Pending</option>
          </select>
        </div>
        <div class="mb-3">
          <label class="form-label">Cover Image (URL)</label>
          <input type="text" class="form-control" v-model="form.cover_image" />
        </div>
        
        <div class="d-flex justify-content-end">
          <button class="btn btn-secondary me-2" @click="showModal=false">Batal</button>
          <button class="btn btn-primary" @click="submit">
            {{ modalMode==='create' ? 'Simpan' : 'Update' }}
          </button>
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
.editor-placeholder { min-height:150px; border:1px solid #ced4da; padding:6px; border-radius:4px; background:#fff }

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
