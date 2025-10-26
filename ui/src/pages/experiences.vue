<script setup>
import { ref, reactive, onMounted, nextTick, onBeforeUnmount } from 'vue'
import axios from 'axios'
import ClassicEditor from '@ckeditor/ckeditor5-build-classic'

// === CKEditor ===
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

// === API ===
const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

// === Custom Alert / Confirm ===
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
  alertModal.onConfirm = () => (alertModal.show = false)
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

// === Data & CRUD ===
const experiences = ref([])
const loading = ref(false)
const error = ref(null)

const fetchExperiences = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/experiences')
    experiences.value = res.data.experiences || res.data || []
  } catch (err) {
    error.value = err.response?.data?.error || err.message
  } finally {
    loading.value = false
  }
}

// === Modal & Form ===
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
  year: '',
  cover_image: null
})

const openCreate = async () => {
  modalMode.value = 'create'
  Object.assign(form, {
    id: null,
    title: '',
    client_name: '',
    description: '',
    start_date: '',
    end_date: '',
    status: 'ongoing',
    year: '',
    cover_image: null
  })
  showModal.value = true
  await nextTick()
  await initEditor()
}

const startEdit = async (row) => {
  modalMode.value = 'edit'
  Object.assign(form, {
    id: row.id,
    title: row.title,
    client_name: row.client_name,
    description: row.description,
    start_date: row.start_date,
    end_date: row.end_date,
    status: row.status,
    year: row.year || '',
    cover_image: null
  })
  showModal.value = true
  await nextTick()
  await initEditor()
}

// === Submit (create / update) ===
const submit = async () => {
  if (editorInstance) form.description = editorInstance.getData()
  if (!form.title || !form.client_name) {
    showAlert('Title dan Client Name wajib diisi', 'Validasi')
    return
  }

  const fd = new FormData()
  fd.append('title', form.title)
  fd.append('client_name', form.client_name)
  fd.append('description', form.description)
  fd.append('start_date', form.start_date)
  fd.append('end_date', form.end_date)
  fd.append('status', form.status)
  fd.append('year', form.year)
  if (form.cover_image) fd.append('cover_image', form.cover_image)

  try {
    if (modalMode.value === 'create') {
      await api.post('/experiences', fd)
    } else {
      try {
        await api.put(`/experiences/${form.id}`, fd)
      } catch (putErr) {
        console.warn('PUT gagal, fallback ke POST + _method=PUT', putErr)
        fd.append('_method', 'PUT')
        await api.post(`/experiences/${form.id}`, fd)
      }
    }
    showModal.value = false
    await fetchExperiences()
    showAlert('Data experience berhasil disimpan!', 'Berhasil')
  } catch (err) {
    console.error('submit error ->', err)
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

// === Delete ===
const deleteExperience = async row => {
  const confirmed = await showConfirm(`Yakin hapus experience "${row.title}"?`, 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/experiences/${row.id}`)
    await fetchExperiences()
    showAlert('Experience berhasil dihapus!', 'Berhasil')
  } catch (err) {
    console.error('delete error ->', err)
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

onMounted(fetchExperiences)
</script>

<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between mb-3">
      <h3>Data Experiences</h3>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah Experience</button>
        <button class="btn btn-secondary" @click="fetchExperiences">Refresh</button>
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
      <div v-for="exp in experiences" :key="exp.id" class="col-md-4 col-sm-6">
        <div class="card h-100">
          <img
            v-if="exp.cover_image"
            :src="exp.cover_image"
            alt=""
            class="industry-image rounded mb-3 card-img-top"
          />
          <div class="card-body">
            <h5 class="card-title">{{ exp.title }}</h5>
            <p class="mb-1"><strong>Client:</strong> {{ exp.client_name }}</p>
            <p class="text-muted small">
              {{ exp.start_date }} - {{ exp.end_date || 'Present' }}
            </p>
            <p class="text-muted small" v-if="exp.year"><strong>Tahun:</strong> {{ exp.year }}</p>
            <div class="small" v-html="exp.description"></div>
            <span class="badge bg-info text-dark">{{ exp.status }}</span>
          </div>
          <div class="card-footer d-flex justify-content-center">
            <button class="btn btn-warning btn-sm me-2" @click="startEdit(exp)">Edit</button>
            <button class="btn btn-danger btn-sm" @click="deleteExperience(exp)">Hapus</button>
          </div>
        </div>
      </div>
    </div>

    <!-- === Modal Experience === -->
    <div
      v-if="showModal"
      class="modal-backdrop"
      style="position: fixed; inset: 0; display: flex; align-items: center; justify-content: center; z-index: 1050"
    >
      <div
        class="card p-3"
        style="width: 600px; max-width: 95vw; max-height: 90vh; overflow-y: auto;"
      >
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h5 class="m-0">{{ modalMode === 'create' ? 'Tambah Experience' : 'Edit Experience' }}</h5>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal = false">×</button>
        </div>

        <div class="row g-2">
          <div class="col-12">
            <label class="form-label">Title</label>
            <input class="form-control" v-model="form.title" />
          </div>

          <div class="col-12">
            <label class="form-label">Client Name</label>
            <input class="form-control" v-model="form.client_name" />
          </div>

          <div class="col-12">
            <label class="form-label">Description</label>
            <div ref="editorElement" class="editor-placeholder"></div>
          </div>

          <div class="col-6">
            <label class="form-label">Start Date</label>
            <input type="date" class="form-control" v-model="form.start_date" />
          </div>

          <div class="col-6">
            <label class="form-label">End Date</label>
            <input type="date" class="form-control" v-model="form.end_date" />
          </div>

          <div class="col-12">
            <label class="form-label">Status</label>
            <select class="form-select" v-model="form.status">
              <option value="ongoing">Ongoing</option>
              <option value="completed">Completed</option>
              <option value="paused">Paused</option>
            </select>
          </div>

          <!-- ✅ Tambahan Field Year -->
          <div class="col-12">
            <label class="form-label">Tahun</label>
            <input
              type="number"
              min="1900"
              max="2099"
              class="form-control"
              v-model="form.year"
              placeholder="Contoh: 2025"
            />
          </div>

          <div class="col-12">
            <label class="form-label">Cover Image</label>
            <input
              type="file"
              class="form-control"
              @change="e => form.cover_image = e.target.files[0]"
            />
          </div>
        </div>

        <div class="d-flex justify-content-end mt-3">
          <button class="btn btn-secondary me-2" @click="showModal = false">Batal</button>
          <button class="btn btn-primary" @click="submit">
            {{ modalMode === 'create' ? 'Simpan' : 'Update' }}
          </button>
        </div>
      </div>
    </div>

    <!-- === Custom Alert / Confirm === -->
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
.editor-placeholder {
  min-height: 150px;
  border: 1px solid #ced4da;
  padding: 6px;
  border-radius: 4px;
  background: #fff;
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

/* Card style */
.card {
  width: 100%;
  max-width: 380px;
  height: 320px;
  margin: 0 auto;
  border-radius: 16px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.card-img-top {
  width: 100%;
  height: 200px;
  object-fit: contain;
  background-color: #fff;
  padding: 10px;
}
</style>
