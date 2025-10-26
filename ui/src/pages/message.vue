<script setup>
import { ref, onMounted, onBeforeUnmount, reactive } from 'vue'
import axios from 'axios'

// DataTables + Bootstrap
import DataTable from 'datatables.net-vue3'
import DataTablesCore from 'datatables.net-bs5'
import Buttons from 'datatables.net-buttons-bs5'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'

// aktifkan plugin
DataTable.use(DataTablesCore)
DataTable.use(Buttons)

// axios instance
const api = axios.create({ baseURL: 'http://localhost:8080' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

// state
const messages = ref([])
const loading = ref(true)
const error = ref(null)

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

const handleError = (err) => {
  console.error('Message error ->', err)
  const msg = err?.response?.data?.message || err.message || 'Gagal memuat data'
  error.value = msg
  showAlert(msg, 'Error')
}

// Kolom tabel
const messageColumns = [
  { title: 'ID', data: 'id' },
  { title: 'Nama', data: 'name' },
  { title: 'Email', data: 'email' },
  { title: 'Telepon', data: 'phone' },
  { title: 'Pesan', data: 'message' },
  { title: 'Dikirim Pada', data: 'created_at' },
  {
    title: 'Aksi',
    data: null,
    orderable: false,
    render: (data, type, row) => {
      return `<button class="btn btn-sm btn-danger delete-btn" data-id="${row.id}">Hapus</button>`
    },
  },
]

// Fetch data
const fetchMessages = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/contact_messages')
    let arr = []
    if (Array.isArray(res.data)) arr = res.data
    else if (Array.isArray(res.data?.data)) arr = res.data.data
    messages.value = arr
  } catch (err) {
    handleError(err)
  } finally {
    loading.value = false
  }
}

// Delete message
const deleteMessage = async (id) => {
  const confirmed = await showConfirm('Yakin ingin menghapus pesan ini?', 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/contact_messages/${id}`)
    messages.value = messages.value.filter(m => m.id !== id)
    showAlert('Pesan berhasil dihapus!', 'Berhasil')
  } catch (err) {
    handleError(err)
  }
}

// event listener wrapper so we can remove it later
const clickHandler = (e) => {
  const target = e.target
  if (target && target.classList && target.classList.contains('delete-btn')) {
    const id = target.getAttribute('data-id')
    if (id) deleteMessage(id)
  }
}

onMounted(() => {
  fetchMessages()
  document.addEventListener('click', clickHandler)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', clickHandler)
})
</script>

<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between mb-3">
      <h3>Daftar Pesan Masuk</h3>
      <button class="btn btn-secondary" @click="fetchMessages">Refresh</button>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="text-center py-4">
      <div class="spinner-border text-primary" role="status"></div>
      <p class="mt-2">Sedang memuat pesan...</p>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="alert alert-danger text-center">
      {{ error }}
    </div>

    <!-- Tabel -->
    <div v-else>
      <DataTable
        class="table table-striped table-hover table-bordered align-middle text-center w-100"
        :data="messages"
        :columns="messageColumns"
        :options="{
          dom: 'lBfrtip',
          buttons: ['copy', 'csv', 'excel', 'pdf', 'print'],
          pageLength: 10,
          lengthMenu: [[10, 25, 50], [10, 25, 50]],
          responsive: true,
        }"
      />
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
