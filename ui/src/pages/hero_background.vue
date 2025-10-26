<template>
  <section class="container py-5">
    <div class="d-flex justify-content-between mb-3">
      <h3>Data Hero Background</h3>
      <div>
        <button
          class="btn btn-primary me-2"
          @click="openCreate"
        >
          + Tambah Hero Background
        </button>
        <button
          class="btn btn-secondary"
          @click="fetchHeroBackgrounds"
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

    <div
      v-else
      class="d-flex overflow-auto gap-3 pb-3"
    >
      <div
        v-for="heroBg in heroBackgrounds"
        :key="heroBg.id"
        class="card text-center shadow-sm hero-bg-card flex-shrink-0"
        style="width: 300px; cursor: pointer"
        @click="openDetail(heroBg)"
      >
        <div class="card-body d-flex flex-column">
          <div class="mb-3">
            <img
              v-if="heroBg.image"
              :src="getImageUrl(heroBg.image)"
              :alt="heroBg.alt_text || 'Hero Background'"
              class="img-fluid hero-bg-image mb-2"
              style="width: 100%; height: 150px; object-fit: cover; border-radius: 8px"
              @error="handleImageError($event)"
            />
            <div
              v-else
              class="bg-secondary d-flex align-items-center justify-content-center mx-auto mb-2"
              style="width: 100%; height: 150px; border-radius: 8px"
            >
              <svg
                width="60"
                height="60"
                viewBox="0 0 24 24"
                fill="none"
                stroke="white"
                stroke-width="2"
              >
                <rect
                  x="3"
                  y="3"
                  width="18"
                  height="18"
                  rx="2"
                  ry="2"
                ></rect>
                <circle
                  cx="8.5"
                  cy="8.5"
                  r="1.5"
                ></circle>
                <polyline points="21,15 16,10 5,21"></polyline>
              </svg>
            </div>
          </div>
          <h5 class="card-title">{{ heroBg.alt_text || 'Tidak ada alt text' }}</h5>
          <p class="card-text text-muted small">{{ formatDate(heroBg.created_at) }}</p>

          <div class="mt-auto">
            <button
              class="btn btn-warning btn-sm me-2"
              @click.stop="startEdit(heroBg)"
            >
              Edit
            </button>
            <button
              class="btn btn-danger btn-sm"
              @click.stop="deleteHeroBackground(heroBg)"
            >
              Hapus
            </button>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- Create/Edit Modal -->
  <div
    v-if="showModal"
    class="modal-backdrop"
    style="position: fixed; inset: 0; display: flex; align-items: center; justify-content: center; z-index: 1050"
  >
    <div
      class="card p-3"
      style="width: 500px; max-width: 95vw"
    >
      <div class="d-flex justify-content-between align-items-center mb-2">
        <h5 class="m-0">{{ modalMode === 'create' ? 'Tambah Hero Background' : 'Edit Hero Background' }}</h5>
        <button
          class="btn btn-sm btn-outline-secondary"
          @click="showModal = false"
        >
          ×
        </button>
      </div>
      <div class="row g-2">
        <div class="col-12">
          <label class="form-label">Gambar *</label>
          <input
            type="file"
            class="form-control"
            accept="image/*"
            @change="e => (form.image = e.target.files[0])"
            required
          />
          <small class="text-muted">Format yang didukung: JPG, PNG, GIF</small>
        </div>
        <div class="col-12">
          <label class="form-label">Alt Text</label>
          <input
            class="form-control"
            v-model="form.alt_text"
            placeholder="Deskripsi gambar untuk aksesibilitas"
          />
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

  <!-- Detail Modal -->
  <div
    v-if="showDetail && selectedHeroBg"
    class="modal-backdrop"
    style="position: fixed; inset: 0; display: flex; align-items: center; justify-content: center; z-index: 1050"
  >
    <div
      class="card p-4"
      style="width: 600px; max-width: 95vw; max-height: 90vh; overflow-y: auto"
    >
      <div class="d-flex justify-content-between align-items-center mb-3">
        <h5 class="m-0">Detail Hero Background</h5>
        <button
          class="btn btn-sm btn-outline-secondary"
          @click="showDetail = false"
        >
          ×
        </button>
      </div>
      <div class="text-center">
        <div class="mb-3">
          <img
            v-if="selectedHeroBg.image"
            :src="getImageUrl(selectedHeroBg.image)"
            :alt="selectedHeroBg.alt_text || 'Hero Background'"
            class="img-fluid"
            style="width: 100%; max-height: 300px; object-fit: cover; border-radius: 8px"
            @error="handleImageError($event)"
          />
          <div
            v-else
            class="bg-secondary d-flex align-items-center justify-content-center mx-auto"
            style="width: 100%; height: 200px; border-radius: 8px"
          >
            <svg
              width="80"
              height="80"
              viewBox="0 0 24 24"
              fill="none"
              stroke="white"
              stroke-width="2"
            >
              <rect
                x="3"
                y="3"
                width="18"
                height="18"
                rx="2"
                ry="2"
              ></rect>
              <circle
                cx="8.5"
                cy="8.5"
                r="1.5"
              ></circle>
              <polyline points="21,15 16,10 5,21"></polyline>
            </svg>
          </div>
        </div>
        <h4 class="mb-2">{{ selectedHeroBg.alt_text || 'Tidak ada alt text' }}</h4>
        <p class="text-muted mb-3">Dibuat: {{ formatDate(selectedHeroBg.created_at) }}</p>
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
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import axios from 'axios'
import 'bootstrap/dist/css/bootstrap.min.css'

const heroBackgrounds = ref([])
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
  onCancel: null,
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

const normalizeHeroBackground = h => ({
  id: h.id ?? null,
  image: h.image ?? '',
  alt_text: h.alt_text ?? '',
  created_at: h.created_at ?? '',
})

const fetchHeroBackgrounds = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/hero_backgrounds')
    const arr = Array.isArray(res.data.hero_backgrounds)
      ? res.data.hero_backgrounds
      : Array.isArray(res.data)
      ? res.data
      : []
    heroBackgrounds.value = arr.map(normalizeHeroBackground)
  } catch (err) {
    error.value = err.response?.data?.error || err.message
  } finally {
    loading.value = false
  }
}

const getImageUrl = imagePath => {
  if (!imagePath) return ''
  const cleanPath = imagePath.startsWith('/') ? imagePath.substring(1) : imagePath
  return `http://localhost:8080/${cleanPath}`
}

const formatDate = dateString => {
  if (!dateString) return 'Tidak ada tanggal'
  const date = new Date(dateString)
  return date.toLocaleDateString('id-ID', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}

const handleImageError = event => {
  event.target.style.display = 'none'
}

const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({
  id: null,
  image: null,
  alt_text: '',
})

const showDetail = ref(false)
const selectedHeroBg = ref(null)

const openDetail = heroBg => {
  selectedHeroBg.value = heroBg
  showDetail.value = true
}

const openCreate = () => {
  modalMode.value = 'create'
  Object.assign(form, {
    id: null,
    image: null,
    alt_text: '',
  })
  showModal.value = true
}

const startEdit = heroBg => {
  modalMode.value = 'edit'
  Object.assign(form, {
    id: heroBg.id,
    image: null, // Reset file input
    alt_text: heroBg.alt_text,
  })
  showModal.value = true
}

const submit = async () => {
  if (!form.image && modalMode.value === 'create') {
    showAlert('Gambar wajib diisi', 'Validasi')
    return
  }

  const fd = new FormData()
  fd.append('alt_text', form.alt_text)
  if (form.image) fd.append('image', form.image)

  try {
    if (modalMode.value === 'create') {
      await api.post('/hero_backgrounds', fd, {
        headers: { 'Content-Type': 'multipart/form-data' },
      })
      showModal.value = false
      await fetchHeroBackgrounds()
      showAlert('Hero background berhasil ditambahkan!', 'Berhasil')
    } else if (modalMode.value === 'edit' && form.id) {
      await api.put(`/hero_backgrounds/${form.id}`, fd, {
        headers: { 'Content-Type': 'multipart/form-data' },
      })
      showModal.value = false
      await fetchHeroBackgrounds()
      showAlert('Hero background berhasil diperbarui!', 'Berhasil')
    }
  } catch (err) {
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

const deleteHeroBackground = async heroBg => {
  const confirmed = await showConfirm(`Hapus hero background ini?`, 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/hero_backgrounds/${heroBg.id}`)
    await fetchHeroBackgrounds()
    showAlert('Hero background berhasil dihapus!', 'Berhasil')
  } catch (err) {
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

onMounted(fetchHeroBackgrounds)
</script>

<style scoped>
.hero-bg-card {
  border-radius: 12px;
  transition: transform 0.2s ease;
}
.hero-bg-card:hover {
  transform: translateY(-5px);
}
.modal-backdrop {
  background: rgba(0, 0, 0, 0.35);
}
.hero-bg-image {
  border: 2px solid #f8f9fa;
}

/* Custom alert modal styles */
.custom-alert-modal {
  background: rgba(0, 0, 0, 0.5);
}
.alert-card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
  width: 380px;
  max-width: 90vw;
  overflow: hidden;
  animation: slideIn 0.2s ease-out;
}
.alert-header {
  padding: 16px 20px 0;
}
.alert-title {
  margin: 0;
  font-size: 1.1rem;
  font-weight: 600;
  color: #1f2937;
}
.alert-body {
  padding: 12px 20px;
}
.alert-message {
  margin: 0;
  color: #6b7280;
  line-height: 1.5;
}
.alert-footer {
  padding: 0 20px 16px;
  display: flex;
  justify-content: flex-end;
  gap: 8px;
}
.alert-footer .btn {
  padding: 6px 14px;
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
