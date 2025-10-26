<script setup>
import { ref, reactive, onMounted } from 'vue'
import axios from 'axios'

import 'bootstrap/dist/css/bootstrap.min.css'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'

import DataTable from 'datatables.net-vue3'
import DataTablesCore from 'datatables.net-bs5'
import Buttons from 'datatables.net-buttons-bs5'
DataTable.use(DataTablesCore)
DataTable.use(Buttons)

const comments = ref([])
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
api.interceptors.response.use(
  r => r,
  e => Promise.reject(e),
)

const commentColumns = [
  { title: 'ID', data: 'id' },
  { title: 'Name', data: 'name' },
  { title: 'Email', data: 'email' },
  { title: 'Comment', data: 'comment' },
  {
    title: 'Aksi',
    data: null,
    render: (data, type, row) => {
      return `
        <button class='btn btn-warning btn-sm me-1 edit-btn'>Edit</button>
        <button class='btn btn-danger btn-sm delete-btn'>Hapus</button>
      `
    },
  },
]

const normalizeComment = c => ({
  id: c.id ?? null,
  post_id: c.post_id ?? (c.post && c.post.id) ?? null,
  name: c.name ?? '',
  email: c.email ?? '',
  comment: c.comment ?? '',
})

const fetchComments = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/blog-comments')
    const arr = Array.isArray(res.data?.comments) ? res.data.comments : []
    comments.value = arr.map(normalizeComment)
  } catch (e) {
    error.value = e?.response?.data?.message || e.message
  } finally {
    loading.value = false
  }
}

const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({ id: null, name: '', email: '', comment: '' })

const openCreate = () => {
  modalMode.value = 'create'
  form.id = null
  form.name = ''
  form.email = ''
  form.comment = ''
  showModal.value = true
}

const openEdit = row => {
  modalMode.value = 'edit'
  form.id = row.id
  form.name = row.name || ''
  form.email = row.email || ''
  form.comment = row.comment || ''
  showModal.value = true
}

const submit = async () => {
  if (!form.comment) {
    showAlert('Comment wajib diisi', 'Validasi')
    return
  }
  try {
    if (modalMode.value === 'create') {
      await api.post('/blog-comments', {
        name: form.name || null,
        email: form.email || null,
        comment: form.comment,
      })
    } else {
      const payload = {}
      payload.name = form.name || null
      payload.email = form.email || null
      if (form.comment) payload.comment = form.comment
      await api.put(`/blog-comments/${form.id}`, payload)
    }
    showModal.value = false
    await fetchComments()
    showAlert('Data berhasil disimpan!', 'Berhasil')
  } catch (e) {
    showAlert(e?.response?.data?.message || e.message, 'Error')
  }
}

const deleteComment = async id => {
  const confirmed = await showConfirm('Yakin hapus komentar ini?', 'Konfirmasi Hapus')
  if (!confirmed) return
  
  try {
    await api.delete(`/blog-comments/${id}`)
    await fetchComments()
    showAlert('Komentar berhasil dihapus!', 'Berhasil')
  } catch (e) {
    showAlert(e?.response?.data?.message || e.message, 'Error')
  }
}

onMounted(fetchComments)
</script>

<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between mb-3">
      <h3>Blog Comments</h3>
      <div>
        <button
          class="btn btn-primary me-2"
          @click="openCreate"
        >
          + Tambah Komentar
        </button>
        <button
          class="btn btn-secondary"
          @click="fetchComments"
        >
          Refresh
        </button>
      </div>
    </div>

    <div
      v-if="loading"
      class="text-center py-4"
    >
      <div
        class="spinner-border text-primary"
        role="status"
      ></div>
      <p class="mt-2">Memuat data...</p>
    </div>

    <div
      v-else-if="error"
      class="alert alert-danger text-center"
    >
      {{ error }}
    </div>

    <div v-else>
      <DataTable
        class="table table-striped table-hover table-bordered align-middle text-center w-100"
        :data="comments"
        :columns="commentColumns"
        :options="{
          dom: 'lBfrtip',
          buttons: ['copy', 'csv', 'excel', 'pdf', 'print'],
          pageLength: 10,
          lengthMenu: [
            [10, 25, 50],
            [10, 25, 50],
          ],
          responsive: true,
          createdRow: function (row, data) {
            const editBtn = row.querySelector('.edit-btn')
            const delBtn = row.querySelector('.delete-btn')
            if (editBtn) editBtn.addEventListener('click', () => openEdit(data))
            if (delBtn) delBtn.addEventListener('click', () => deleteComment(data.id))
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
      <div
        class="card p-3"
        style="width: 520px; max-width: 95vw"
      >
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h5 class="m-0">{{ modalMode === 'create' ? 'Tambah Komentar' : 'Edit Komentar' }}</h5>
          <button
            class="btn btn-sm btn-outline-secondary"
            @click="showModal = false"
          >
            ×
          </button>
        </div>
        <div class="row g-2">
          <div class="col-md-6">
            <label class="form-label">Name</label>
            <input
              class="form-control"
              v-model="form.name"
            />
          </div>
          <div class="col-md-6">
            <label class="form-label">Email</label>
            <input
              class="form-control"
              v-model="form.email"
            />
          </div>
          <div class="col-12">
            <label class="form-label">Comment</label>
            <textarea
              class="form-control"
              rows="4"
              v-model="form.comment"
            ></textarea>
          </div>
        </div>
        <div class="d-flex justify-content-end mt-3">
          <button
            class="btn btn-secondary me-2"
            @click="showModal = false"
          >
            Batal
          </button>
          <button
            class="btn btn-primary"
            @click="submit"
          >
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