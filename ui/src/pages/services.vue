<template>
  <section class="hero-section d-flex align-items-center text-white"
    :style="{background: `url(${bgHero}) no-repeat center center`, backgroundSize: 'cover', height: '70vh'}">
    <div class="container text-center">
      <h1 class="display-4 fw-bold">Services that scale with your digital strategies</h1>
      <p class="lead mt-3">Kami menyediakan layanan untuk mendukung transformasi digital bisnis Anda</p>
    </div>
  </section>

  <section class="container py-5">
    <div class="d-flex justify-content-between mb-3">
      <h3>Data Service</h3>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah Service</button>
        <button class="btn btn-secondary" @click="fetchServices">Refresh</button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-4">
      <div class="spinner-border text-primary" role="status"></div>
      <p class="mt-2">Memuat data...</p>
    </div>

    <div v-else-if="error" class="alert alert-danger text-center">{{ error }}</div>

    <div v-else class="d-flex overflow-auto gap-3 pb-3">
      <div v-for="srv in services" :key="srv.id"
        class="card text-center shadow-sm service-card flex-shrink-0"
        style="width: 260px; cursor: pointer"
        @click="openDetail(srv)">
        <div class="card-body d-flex flex-column">
          <img v-if="srv.icon" :src="getImageUrl(srv.icon)" alt="logo" class="img-fluid mb-2" style="max-height:60px"/>
          <h5 class="card-title">{{ srv.title }}</h5>
          <p class="card-text text-muted flex-grow-1">{{ truncateDescription(srv.description, 50) }}</p>
          <div class="mt-3">
            <button class="btn btn-warning btn-sm me-2" @click.stop="startEdit(srv)">Edit</button>
            <button class="btn btn-danger btn-sm" @click.stop="deleteService(srv)">Hapus</button>
          </div>
        </div>
      </div>
    </div>
  </section>

  <div v-if="showModal" class="modal-backdrop" style="position:fixed;inset:0;display:flex;align-items:center;justify-content:center;z-index:1050">
    <div class="card p-3" style="width:500px;max-width:95vw">
      <div class="d-flex justify-content-between align-items-center mb-2">
        <h5 class="m-0">{{ modalMode==='create'?'Tambah Service':'Edit Service' }}</h5>
        <button class="btn btn-sm btn-outline-secondary" @click="showModal=false">×</button>
      </div>
      <div class="row g-2">
        <div class="col-12">
          <label class="form-label">Judul Service</label>
          <input class="form-control" v-model="form.title"/>
        </div>
        <div class="col-12">
          <label class="form-label">Deskripsi</label>
          <textarea class="form-control" rows="3" v-model="form.description"></textarea>
        </div>
        <div class="col-12">
          <label class="form-label">Logo</label>
          <input type="file" class="form-control" @change="e => form.icon = e.target.files[0]"/>
        </div>
        <div class="col-12">
          <label class="form-label">Image</label>
          <input type="file" class="form-control" @change="e => form.image = e.target.files[0]"/>
        </div>
      </div>
      <div class="d-flex justify-content-end mt-3">
        <button class="btn btn-secondary me-2" @click="showModal=false">Batal</button>
        <button class="btn btn-primary" @click="submit">{{ modalMode==='create'?'Simpan':'Update' }}</button>
      </div>
    </div>
  </div>

  <div v-if="showDetail && selectedService" class="modal-backdrop" style="position:fixed;inset:0;display:flex;align-items:center;justify-content:center;z-index:1050">
    <div class="card p-3" style="width:500px;max-width:95vw;max-height:90vh;overflow-y:auto">
      <div class="d-flex justify-content-between align-items-center mb-3">
        <h5 class="m-0">Detail Service</h5>
        <button class="btn btn-sm btn-outline-secondary" @click="showDetail=false">×</button>
      </div>
      <div class="text-center">
        <div v-if="selectedService.icon" class="mb-3">
          <img 
            :src="getImageUrl(selectedService.icon)" 
            :alt="selectedService.title + ' logo'"
            class="img-fluid"
            style="max-height:80px;object-fit:contain"
            @error="handleImageError($event, 'icon')"
          />
        </div>
        <h5 class="mb-3">{{ selectedService.title }}</h5>
        <div class="mb-4 text-start">
          <h6 class="text-muted mb-2">Deskripsi Lengkap:</h6>
          <p class="text-muted" style="line-height: 1.6; white-space: pre-line;">{{ selectedService.description || 'Tidak ada deskripsi' }}</p>
        </div>
        <div v-if="hasValidImage(selectedService.image)" class="mb-3">
          <img 
            :src="getImageUrl(selectedService.image)" 
            :alt="selectedService.title + ' image'"
            class="img-fluid rounded"
            style="max-height:300px;max-width:100%;object-fit:contain"
            @error="handleImageError($event, 'image')"
          />
        </div>
        <div v-else class="mb-3 p-4 bg-light rounded text-muted text-center">
          <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="mb-2">
            <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
            <circle cx="8.5" cy="8.5" r="1.5"/>
            <polyline points="21,15 16,10 5,21"/>
          </svg>
          <p class="mb-0">Tidak ada gambar untuk service ini</p>
        </div>
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
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import axios from 'axios'
import bgHero from '@/assets/images/bg_service.jpg'
import 'bootstrap/dist/css/bootstrap.min.css'

const services = ref([])
const loading = ref(false)
const error = ref(null)

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

const normalizeService = s => ({
  id: s.id ?? null,
  title: s.title ?? '',
  description: s.description ?? '',
  icon: s.icon ?? '',
  image: s.images ?? null
})

const fetchServices = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/services')
    const arr = Array.isArray(res.data.services) ? res.data.services : []
    services.value = arr.map(normalizeService)
  } catch (err) {
    error.value = err.response?.data?.error || err.message
  } finally {
    loading.value = false
  }
}

const getImageUrl = (imagePath) => {
  if (!imagePath) return ''
  const cleanPath = imagePath.startsWith('/') ? imagePath.substring(1) : imagePath
  return `http://localhost:8080/${cleanPath}`
}

const hasValidImage = (image) => {
  return image && typeof image === 'string' && image.trim() !== ''
}

const truncateDescription = (text, maxLength) => {
  if (!text) return 'Tidak ada deskripsi'
  if (text.length <= maxLength) return text
  return text.substring(0, maxLength) + '...'
}

const handleImageError = (event) => {
  event.target.style.display = 'none'
}

const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({ id: null, title: '', description: '', icon: null, image: null })

const showDetail = ref(false)
const selectedService = ref(null)

const openDetail = srv => {
  selectedService.value = srv
  showDetail.value = true
}

const openCreate = () => {
  modalMode.value = 'create'
  form.id = null
  form.title = ''
  form.description = ''
  form.icon = null
  form.image = null
  showModal.value = true
}

const startEdit = row => {
  modalMode.value = 'edit'
  form.id = row.id
  form.title = row.title
  form.description = row.description
  form.icon = null
  form.image = null
  showModal.value = true
}

const submit = async () => {
  if (!form.title) {
    showAlert('Judul service wajib diisi', 'Validasi')
    return
  }
  const fd = new FormData()
  fd.append('title', form.title)
  fd.append('description', form.description)
  if (form.icon) fd.append('icon', form.icon)
  if (form.image) fd.append('image', form.image)

  try {
    if (modalMode.value === 'create') {
      await api.post('/services', fd, { headers: { 'Content-Type': 'multipart/form-data' } })
      showModal.value = false
      await fetchServices()
      showAlert('Service berhasil ditambahkan!', 'Berhasil')
    } else if (modalMode.value === 'edit' && form.id) {
      await api.put(`/services/${form.id}`, fd, { headers: { 'Content-Type': 'multipart/form-data' } })
      showModal.value = false
      await fetchServices()
      showAlert('Service berhasil diperbarui!', 'Berhasil')
    }
  } catch (err) {
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

const deleteService = async row => {
  const confirmed = await showConfirm(`Hapus ${row.title}?`, 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/services/${row.id}`)
    await fetchServices()
    showAlert('Service berhasil dihapus!', 'Berhasil')
  } catch (err) {
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

onMounted(fetchServices)
</script>

<style scoped>
.hero-section { position: relative; }
.hero-section::after {
  content: "";
  position: absolute; inset: 0;
  background: rgba(0,0,0,0.5);
}
.hero-section .container { position: relative; z-index: 1; }
.service-card { border-radius: 12px; transition: transform 0.2s ease; }
.service-card:hover { transform: translateY(-5px); }
.modal-backdrop { background: rgba(0,0,0,0.35); }

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
