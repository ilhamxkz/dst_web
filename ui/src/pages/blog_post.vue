<script setup>
import axios from 'axios'
import { onBeforeUnmount, onMounted, reactive, ref } from 'vue'

import 'bootstrap/dist/css/bootstrap.min.css'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'

import DataTablesCore from 'datatables.net-bs5'
import Buttons from 'datatables.net-buttons-bs5'
import DataTable from 'datatables.net-vue3'
DataTable.use(DataTablesCore)
DataTable.use(Buttons)

// CKEditor
import ClassicEditor from '@ckeditor/ckeditor5-build-classic'

const posts = ref([])
const loading = ref(true)
const error = ref(null)

// Custom Alert/Confirm State
const alertModal = reactive({
  show: false,
  type: 'alert', // 'alert' or 'confirm'
  title: '',
  message: '',
  onConfirm: null,
  onCancel: null
})

// Custom Alert Function
const showAlert = (message, title = 'Peringatan') => {
  alertModal.show = true
  alertModal.type = 'alert'
  alertModal.title = title
  alertModal.message = message
  alertModal.onConfirm = () => {
    alertModal.show = false
  }
}

// Custom Confirm Function
const showConfirm = (message, title = 'Konfirmasi') => {
  return new Promise((resolve) => {
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

const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})
api.interceptors.response.use(r => r, e => Promise.reject(e))

const postColumns = [
  { title: 'ID', data: 'id' },
  { title: 'Title', data: 'title' },
  { title: 'Slug', data: 'slug' },
  { title: 'Author ID', data: 'author_id' },
  {
    title: 'Aksi',
    data: null,
    render: () => {
      return `
        <button class='btn btn-warning btn-sm me-1 edit-btn'>Edit</button>
        <button class='btn btn-danger btn-sm delete-btn'>Hapus</button>
      `
    },
  },
]

const normalizePost = p => ({
  id: p.id ?? p.ID ?? null,
  title: p.title ?? '',
  slug: p.slug ?? '',
  content: p.content ?? '',
  cover_image: p.cover_image ?? '',
  author_id: p.author_id ?? (p.author && p.author.id) ?? null,
})

const fetchPosts = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/blog-posts')
    const arr = Array.isArray(res.data?.posts) ? res.data.posts : []
    posts.value = arr.map(normalizePost)
  } catch (e) {
    error.value = e?.response?.data?.message || e.message
  } finally {
    loading.value = false
  }
}

const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({
  id: null,
  title: '',
  slug: '',
  content: '',
  cover_image: '',
  author_id: null,
})

const editorElement = ref(null)
let editorInstance = null

const initEditor = async (data = '') => {
  if (editorInstance) {
    await editorInstance.destroy()
    editorInstance = null
  }
  try {
    const config = {
      toolbar: {
        items: [
          'undo', 'redo', '|',
          'bold', 'italic', 'underline', 'link', '|',
          'insertTable', 'bulletedList', 'numberedList', '|',
          'blockQuote', 'removeFormat'
        ],
        shouldNotGroupWhenFull: true
      },
      placeholder: 'Select a template above or start typing...',
      table: { contentToolbar: ['tableColumn', 'tableRow', 'mergeTableCells'] }
    }
    editorInstance = await ClassicEditor.create(editorElement.value, config)
    editorInstance.setData(data)
    editorInstance.model.document.on('change:data', () => {
      form.content = editorInstance.getData()
    })
  } catch (err) {
    console.error('CKEditor init failed:', err)
    showAlert('Editor initialization failed: ' + (err.message || err), 'Error')
  }
}

const openCreate = async () => {
  modalMode.value = 'create'
  form.id = null
  form.title = ''
  form.slug = ''
  form.content = ''
  form.cover_image = ''
  form.author_id = null
  showModal.value = true
  setTimeout(() => initEditor(''), 50)
}

const openEdit = async row => {
  modalMode.value = 'edit'
  form.id = row.id
  form.title = row.title || ''
  form.slug = row.slug || ''
  form.content = row.content || ''
  form.cover_image = row.cover_image || ''
  form.author_id = row.author_id || null
  showModal.value = true
  setTimeout(() => initEditor(form.content), 50)
}

const submit = async () => {
  if (editorInstance) {
    form.content = editorInstance.getData()
  }

  if (!form.title || !form.content) {
    showAlert('Title dan Content wajib diisi', 'Validasi')
    return
  }

  try {
    if (modalMode.value === 'create') {
      await api.post('/blog-posts', {
        title: form.title,
        slug: form.slug || null,
        content: form.content,
        cover_image: form.cover_image || null,
        author_id: form.author_id || null,
      })
    } else {
      const payload = {}
      if (form.title) payload.title = form.title
      if (form.slug !== '') payload.slug = form.slug || null
      if (form.content) payload.content = form.content
      if (form.cover_image !== '') payload.cover_image = form.cover_image || null
      if (form.author_id != null) payload.author_id = form.author_id
      await api.put(`/blog-posts/${form.id}`, payload)
    }
    showModal.value = false
    await fetchPosts()
    showAlert('Data berhasil disimpan!', 'Berhasil')
  } catch (e) {
    showAlert(e?.response?.data?.message || e.message, 'Error')
  }
}

const deletePost = async id => {
  const confirmed = await showConfirm('Yakin hapus post ini?', 'Konfirmasi Hapus')
  if (!confirmed) return
  
  try {
    await api.delete(`/blog-posts/${id}`)
    await fetchPosts()
    showAlert('Post berhasil dihapus!', 'Berhasil')
  } catch (e) {
    showAlert(e?.response?.data?.message || e.message, 'Error')
  }
}

onMounted(fetchPosts)

onBeforeUnmount(() => {
  if (editorInstance) {
    editorInstance.destroy().catch(err => console.warn('Error destroying editor:', err))
    editorInstance = null
  }
})
</script>

<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between mb-3">
      <h3>Blog Posts</h3>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah Post</button>
        <button class="btn btn-secondary" @click="fetchPosts">Refresh</button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-4">
      <div class="spinner-border text-primary" role="status"></div>
      <p class="mt-2">Memuat data...</p>
    </div>

    <div v-else-if="error" class="alert alert-danger text-center">
      {{ error }}
    </div>

    <div v-else>
      <DataTable
        class="table table-striped table-hover table-bordered align-middle text-center w-100"
        :data="posts"
        :columns="postColumns"
        :options="{
          dom: 'lBfrtip',
          buttons: ['copy', 'csv', 'excel', 'pdf', 'print'],
          pageLength: 10,
          lengthMenu: [[10, 25, 50],[10, 25, 50]],
          responsive: true,
          createdRow: function (row, data) {
            const editBtn = row.querySelector('.edit-btn')
            const delBtn = row.querySelector('.delete-btn')
            if (editBtn) editBtn.addEventListener('click', () => openEdit(data))
            if (delBtn) delBtn.addEventListener('click', () => deletePost(data.id))
          },
        }"
      />
    </div>

    <!-- Form Modal -->
    <div
      v-if="showModal"
      class="modal-backdrop"
      style="position: fixed; inset: 0; display: flex; align-items: center; justify-content: center; z-index: 1050"
    >
      <div class="card p-3" style="width: 520px; max-width: 95vw">
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h5 class="m-0">{{ modalMode === 'create' ? 'Tambah Post' : 'Edit Post' }}</h5>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal = false">×</button>
        </div>
        <div class="row g-2">
          <div class="col-12">
            <label class="form-label">Title</label>
            <input class="form-control" v-model="form.title" />
          </div>
          <div class="col-md-6">
            <label class="form-label">Slug</label>
            <input class="form-control" v-model="form.slug" />
          </div>
          <div class="col-md-6">
            <label class="form-label">Author ID</label>
            <input class="form-control" v-model.number="form.author_id" type="number" />
          </div>
          <div class="col-12">
            <label class="form-label">Cover Image URL</label>
            <input class="form-control" v-model="form.cover_image" />
          </div>
          <div class="col-12">
            <label class="form-label">Content</label>
            <div ref="editorElement" class="editor-placeholder"></div>
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
            <button
              class="btn btn-secondary me-2"
              @click="alertModal.onCancel"
            >
              Cancel
            </button>
            <button
              class="btn btn-primary"
              @click="alertModal.onConfirm"
            >
              OK
            </button>
          </template>
          <template v-else>
            <button
              class="btn btn-primary"
              @click="alertModal.onConfirm"
            >
              OK
            </button>
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
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
  width: 400px;
  max-width: 90vw;
  overflow: hidden;
  animation: slideIn 0.2s ease-out;
}

.alert-header {
  padding: 20px 24px 0;
  border-bottom: none;
}

.alert-title {
  margin: 0;
  font-size: 1.1rem;
  font-weight: 600;
  color: #1f2937;
}

.alert-body {
  padding: 16px 24px;
}

.alert-message {
  margin: 0;
  color: #6b7280;
  line-height: 1.5;
}

.alert-footer {
  padding: 0 24px 24px;
  display: flex;
  justify-content: flex-end;
  gap: 8px;
}

.alert-footer .btn {
  padding: 8px 16px;
  border-radius: 6px;
  font-size: 0.875rem;
  font-weight: 500;
}

.editor-placeholder {
  border: 1px solid #e5e7eb;
  min-height: 280px;
  border-radius: 8px;
  padding: 12px;
  background: white;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.03);
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(-16px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>