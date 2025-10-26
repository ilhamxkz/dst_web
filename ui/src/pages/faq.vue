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

const faqs = ref([])
const loading = ref(true)
const error = ref(null)

const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})
api.interceptors.response.use(r => r, e => Promise.reject(e))

const faqColumns = [
  { title: 'ID', data: 'id' },
  { title: 'Question', data: 'question' },
  {
    title: 'Answer',
    data: 'answer',
    render: (data) => {
      if (!data) return ''
      // tampilkan ringkasan plain text (maks 120 char)
      const txt = data.replace(/<[^>]*>/g, '') // hapus tag HTML
      return txt.length > 120 ? txt.slice(0, 120) + '...' : txt
    },
  },
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

const normalizeFaq = f => ({
  id: f.id ?? f.ID ?? null,
  question: f.question ?? '',
  answer: f.answer ?? '',
})

const fetchFaqs = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/faqs')
    const arr = Array.isArray(res.data?.faqs) ? res.data.faqs : []
    faqs.value = arr.map(normalizeFaq)
  } catch (e) {
    error.value = e?.response?.data?.message || e.message
  } finally {
    loading.value = false
  }
}

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

const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({
  id: null,
  question: '',
  answer: '',
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
          'bulletedList', 'numberedList', '|',
          'blockQuote', 'removeFormat'
        ],
        shouldNotGroupWhenFull: true
      },
      placeholder: 'Tulis jawaban di sini...',
    }
    editorInstance = await ClassicEditor.create(editorElement.value, config)
    editorInstance.setData(data)
    editorInstance.model.document.on('change:data', () => {
      form.answer = editorInstance.getData()
    })
  } catch (err) {
    console.error('CKEditor init failed:', err)
    showAlert('Editor initialization failed: ' + (err.message || err), 'Error')
  }
}

const openCreate = async () => {
  modalMode.value = 'create'
  form.id = null
  form.question = ''
  form.answer = ''
  showModal.value = true
  setTimeout(() => initEditor(''), 50)
}

const openEdit = async row => {
  modalMode.value = 'edit'
  form.id = row.id
  form.question = row.question || ''
  form.answer = row.answer || ''
  showModal.value = true
  setTimeout(() => initEditor(form.answer), 50)
}

const submit = async () => {
  if (editorInstance) {
    form.answer = editorInstance.getData()
  }

  if (!form.question || !form.answer) {
    showAlert('Question dan Answer wajib diisi', 'Validasi')
    return
  }

  try {
    if (modalMode.value === 'create') {
      await api.post('/faqs', {
        question: form.question,
        answer: form.answer,
      })
      showModal.value = false
      await fetchFaqs()
      showAlert('FAQ berhasil ditambahkan!', 'Berhasil')
    } else {
      const payload = {}
      if (form.question) payload.question = form.question
      if (form.answer) payload.answer = form.answer
      await api.put(`/faqs/${form.id}`, payload)
      showModal.value = false
      await fetchFaqs()
      showAlert('FAQ berhasil diperbarui!', 'Berhasil')
    }
  } catch (e) {
    showAlert(e?.response?.data?.message || e.message, 'Error')
  }
}

const deleteFaq = async id => {
  const confirmed = await showConfirm('Yakin hapus FAQ ini?', 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/faqs/${id}`)
    await fetchFaqs()
    showAlert('FAQ berhasil dihapus!', 'Berhasil')
  } catch (e) {
    showAlert(e?.response?.data?.message || e.message, 'Error')
  }
}

onMounted(fetchFaqs)

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
      <h3>FAQ</h3>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah FAQ</button>
        <button class="btn btn-secondary" @click="fetchFaqs">Refresh</button>
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
        :data="faqs"
        :columns="faqColumns"
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
            if (delBtn) delBtn.addEventListener('click', () => deleteFaq(data.id))
          },
        }"
      />
    </div>

    <!-- Modal -->
    <div
      v-if="showModal"
      class="modal-backdrop"
      style="position: fixed; inset: 0; display: flex; align-items: center; justify-content: center; z-index: 1050"
    >
      <div class="card p-3" style="width: 720px; max-width: 95vw">
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h5 class="m-0">{{ modalMode === 'create' ? 'Tambah FAQ' : 'Edit FAQ' }}</h5>
          <button class="btn btn-sm btn-outline-secondary" @click="showModal = false">×</button>
        </div>
        <div class="row g-2">
          <div class="col-12">
            <label class="form-label">Question</label>
            <input class="form-control" v-model="form.question" />
          </div>
          <div class="col-12">
            <label class="form-label">Answer</label>
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
.modal-backdrop {
  background: rgba(0, 0, 0, 0.35);
}
.editor-placeholder {
  border: 1px solid #e5e7eb;
  min-height: 220px;
  border-radius: 8px;
  padding: 12px;
  background: white;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.03);
}

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
