<template>
  <div class="container menu-page mt-4">
    <div class="d-flex justify-content-between align-items-center mb-3">
      <h3>Menu sidebar</h3>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah Menu</button>
        <button class="btn btn-secondary" @click="fetchMenus">Refresh</button>
      </div>
    </div>

    <!-- inline status (opsional) -->
    <div v-if="status.message" :class="['status', status.type]" class="mb-3">
      {{ status.message }}
    </div>

    <!-- Tabel daftar menu (DataTables) -->
    <div v-if="loading" class="text-center py-4">
      <div class="spinner-border text-primary" role="status"></div>
      <p class="mt-2">Memuat data...</p>
    </div>

    <div v-else-if="error" class="alert alert-danger text-center mb-3">
      {{ error }}
    </div>

    <div v-else>
      <DataTable
        class="table table-striped table-hover table-bordered align-middle text-center w-100"
        :data="menus"
        :columns="menuColumns"
        :options="dataTableOptions"
      />
    </div>

    <!-- Form Modal (appear in front of screen like posts code) -->
    <div
      v-if="showModal"
      class="modal-backdrop"
      style="position: fixed; inset: 0; display:flex; align-items:center; justify-content:center; z-index:1050;"
    >
      <div class="card p-3" style="width:420px; max-width:95vw;">
        <div class="d-flex justify-content-between align-items-center mb-2">
          <h5 class="m-0">{{ editing ? 'Edit Menu' : 'Tambah Menu' }}</h5>
          <button class="btn btn-sm btn-outline-secondary" @click="closeModal">×</button>
        </div>

        <form @submit.prevent="submitForm">
          <div class="row g-2">
            <div class="col-12">
              <label class="form-label">Nama Menu</label>
              <input ref="namaInput" class="form-control" v-model="form.nama_menu" required />
            </div>
            <div class="col-12">
              <label class="form-label">Routes</label>
              <input class="form-control" v-model="form.routes" required />
            </div>
          </div>

          <div class="d-flex justify-content-end gap-2 mt-3">
            <button type="button" class="btn btn-secondary" v-if="editing" @click="cancelEdit">Batal</button>
            <button type="submit" class="btn btn-primary">
              {{ editing ? 'Simpan Perubahan' : 'Tambah' }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Custom Alert/Confirm Modal (consistent with other pages) -->
    <div
      v-if="alertModal.show"
      class="modal-backdrop custom-alert-modal"
      style="position: fixed; inset: 0; display:flex; align-items:center; justify-content:center; z-index:1060;"
    >
      <div class="alert-card">
        <div class="alert-header"><h5 class="alert-title">{{ alertModal.title }}</h5></div>
        <div class="alert-body"><p class="alert-message">{{ alertModal.message }}</p></div>
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

<script setup>
import { ref, reactive, onMounted, nextTick } from 'vue'
import axios from 'axios'

// styles
import 'bootstrap/dist/css/bootstrap.min.css'
import 'datatables.net-bs5/css/dataTables.bootstrap5.min.css'
import 'datatables.net-buttons-bs5/css/buttons.bootstrap5.min.css'

// DataTables
import DataTable from 'datatables.net-vue3'
import DataTablesCore from 'datatables.net-bs5'
import Buttons from 'datatables.net-buttons-bs5'
DataTable.use(DataTablesCore)
DataTable.use(Buttons)

// axios client
const api = axios.create({ baseURL: 'http://localhost:8080/api' })
api.interceptors.request.use(cfg => {
  const token = localStorage.getItem('token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

// state
const menus = ref([])
const loading = ref(true)
const error = ref(null)
const status = reactive({ message: '', type: '' })

const form = reactive({ id: null, nama_menu: '', routes: '' })
const editing = ref(false)
const showModal = ref(false)
const namaInput = ref(null)

// Custom Alert/Confirm State & Helpers
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
  alertModal.onConfirm = () => {
    alertModal.show = false
  }
  status.message = message
  status.type = title === 'Error' ? 'error' : 'success'
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

// fetch menus
const fetchMenus = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/menus')
    menus.value = Array.isArray(res.data?.menus) ? res.data.menus : (Array.isArray(res.data) ? res.data : [])
  } catch (err) {
    console.error(err)
    error.value = err?.response?.data?.message || 'Gagal mengambil data menu.'
    showAlert(error.value, 'Error')
  } finally {
    loading.value = false
  }
}

// open create modal
const openCreate = async () => {
  resetForm()
  editing.value = false
  showModal.value = true
  await nextTick()
  namaInput.value?.focus()
}

// open edit modal
const startEdit = (m) => {
  editing.value = true
  form.id = m.id
  form.nama_menu = m.nama_menu
  form.routes = m.routes
  showModal.value = true
  nextTick(() => namaInput.value?.focus())
}

const closeModal = () => {
  showModal.value = false
}

// submit create/update
const submitForm = async () => {
  try {
    if (!form.nama_menu || !form.routes) {
      showAlert('Nama Menu dan Routes wajib diisi', 'Validasi')
      return
    }

    if (editing.value) {
      await api.put(`/menus/${form.id}`, { nama_menu: form.nama_menu, routes: form.routes })
      showAlert('Menu berhasil diperbarui.', 'Berhasil')
    } else {
      const res = await api.post('/menus', { nama_menu: form.nama_menu, routes: form.routes })
      if (res.data?.menu) menus.value.push(res.data.menu)
      showAlert('Menu berhasil ditambahkan.', 'Berhasil')
    }

    closeModal()
    await fetchMenus()
    resetForm()
  } catch (err) {
    console.error(err)
    const msg = err.response?.data?.message || 'Gagal menyimpan menu.'
    showAlert(msg, 'Error')
  }
}

const cancelEdit = () => {
  resetForm()
  closeModal()
}

const resetForm = () => {
  editing.value = false
  form.id = null
  form.nama_menu = ''
  form.routes = ''
}

// delete with confirm
const removeMenu = async (id) => {
  const confirmed = await showConfirm('Yakin ingin menghapus menu ini?', 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/menus/${id}`)
    showAlert('Menu berhasil dihapus.', 'Berhasil')
    await fetchMenus()
  } catch (err) {
    console.error(err)
    const msg = err.response?.data?.message || 'Gagal menghapus menu.'
    showAlert(msg, 'Error')
  }
}

// DataTables columns & options
const menuColumns = [
  { title: 'ID', data: 'id' },
  { title: 'Nama Menu', data: 'nama_menu' },
  { title: 'Routes', data: 'routes' },
  {
    title: 'Aksi',
    data: null,
    render: () => `
      <button class='btn btn-warning btn-sm me-1 edit-btn'>Edit</button>
      <button class='btn btn-danger btn-sm delete-btn'>Hapus</button>
    `
  }
]

const dataTableOptions = {
  dom: 'lBfrtip',
  buttons: ['copy', 'csv', 'excel', 'pdf', 'print'],
  pageLength: 10,
  lengthMenu: [[10, 25, 50], [10, 25, 50]],
  responsive: true,
  createdRow: function (row, data) {
    const editBtn = row.querySelector('.edit-btn')
    const delBtn = row.querySelector('.delete-btn')
    if (editBtn) editBtn.addEventListener('click', () => startEdit(data))
    if (delBtn) delBtn.addEventListener('click', () => removeMenu(data.id))
  }
}

onMounted(fetchMenus)
</script>

<style scoped>
.menu-page { max-width: 1100px; margin: 12px auto; font-family: Inter, system-ui, -apple-system, "Segoe UI", Roboto, "Helvetica Neue", Arial; }

.status { padding: 10px 12px; border-radius: 8px; margin-bottom: 12px; font-weight: 500; }
.status.success { background: #e8f5e9; color: #2e7d32; }
.status.error { background: #ffebee; color: #c62828; }

/* DataTable/header styling */
table.dataTable th { background: #f5f5f5; font-weight: 600; text-align: center; }
table.dataTable td { vertical-align: middle; text-align: center; }
table.dataTable tbody tr:hover { background: #f1f1f1; transition: background-color 0.15s ease; }

/* Modal form/backdrop */
.modal-backdrop { background: rgba(0,0,0,0.35); }

/* Alert modal styles (same as other pages) */
.custom-alert-modal { background: rgba(0, 0, 0, 0.5); }
.alert-card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 20px 25px -5px rgba(0,0,0,.1), 0 10px 10px -5px rgba(0,0,0,.04);
  width: 400px; max-width: 90vw; overflow: hidden;
}
.alert-header { padding: 20px 24px 0; border-bottom: none; }
.alert-title { margin: 0; font-size: 1.05rem; font-weight: 600; color: #111827; }
.alert-body { padding: 12px 24px; }
.alert-message { margin: 0; color: #6b7280; line-height: 1.5; word-break: break-word; }
.alert-footer { padding: 0 24px 16px; display: flex; justify-content: flex-end; gap: 8px; }
.alert-footer .btn { padding: 6px 14px; border-radius: 6px; font-size: .9rem; font-weight: 500; }
</style>
