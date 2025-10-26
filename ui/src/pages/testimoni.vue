<template>
  <section class="container py-5">
    <div class="d-flex justify-content-between mb-3">
      <h3>Data Testimoni</h3>
      <div>
        <button
          class="btn btn-primary me-2"
          @click="openCreate"
        >
          + Tambah Testimoni
        </button>
        <button
          class="btn btn-secondary"
          @click="fetchTestimonials"
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
        v-for="testimonial in testimonials"
        :key="testimonial.id"
        class="card text-center shadow-sm testimonial-card flex-shrink-0"
        style="width: 280px; cursor: pointer"
        @click="openDetail(testimonial)"
      >
        <div class="card-body d-flex flex-column">
          <div class="mb-3">
            <img
              v-if="testimonial.url_foto_profil"
              :src="getImageUrl(testimonial.url_foto_profil)"
              :alt="testimonial.nama"
              class="rounded-circle img-fluid testimonial-photo mb-2"
              style="width: 80px; height: 80px; object-fit: cover"
              @error="handleImageError($event)"
            />
            <div
              v-else
              class="rounded-circle bg-secondary d-flex align-items-center justify-content-center mx-auto mb-2"
              style="width: 80px; height: 80px"
            >
              <svg
                width="40"
                height="40"
                viewBox="0 0 24 24"
                fill="none"
                stroke="white"
                stroke-width="2"
              >
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle
                  cx="12"
                  cy="7"
                  r="4"
                ></circle>
              </svg>
            </div>
          </div>
          <h5 class="card-title">{{ testimonial.nama }}</h5>
          <p class="card-text text-primary fw-semibold">{{ testimonial.perusahaan || 'Tidak ada perusahaan' }}</p>
          <p class="card-text text-muted flex-grow-1 small">{{ truncateDescription(testimonial.pesan, 80) }}</p>

          <div class="mt-auto">
            <button
              class="btn btn-warning btn-sm me-2"
              @click.stop="startEdit(testimonial)"
            >
              Edit
            </button>
            <button
              class="btn btn-danger btn-sm"
              @click.stop="deleteTestimonial(testimonial)"
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
        <h5 class="m-0">{{ modalMode === 'create' ? 'Tambah Testimoni' : 'Edit Testimoni' }}</h5>
        <button
          class="btn btn-sm btn-outline-secondary"
          @click="showModal = false"
        >
          ×
        </button>
      </div>
      <div class="row g-2">
        <div class="col-12">
          <label class="form-label">Nama *</label>
          <input
            class="form-control"
            v-model="form.nama"
            required
          />
        </div>
        <div class="col-12">
          <label class="form-label">Perusahaan</label>
          <input
            class="form-control"
            v-model="form.perusahaan"
          />
        </div>
        <div class="col-12">
          <label class="form-label">Pesan *</label>
          <textarea
            class="form-control"
            rows="3"
            v-model="form.pesan"
            required
          ></textarea>
        </div>
        <div class="col-12">
          <label class="form-label">Foto Profil</label>
          <input
            type="file"
            class="form-control"
            accept="image/*"
            @change="e => (form.url_foto_profil = e.target.files[0])"
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
    v-if="showDetail && selectedTestimonial"
    class="modal-backdrop"
    style="position: fixed; inset: 0; display: flex; align-items: center; justify-content: center; z-index: 1050"
  >
    <div
      class="card p-4"
      style="width: 500px; max-width: 95vw; max-height: 90vh; overflow-y: auto"
    >
      <div class="d-flex justify-content-between align-items-center mb-3">
        <h5 class="m-0">Detail Testimoni</h5>
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
            v-if="selectedTestimonial.url_foto_profil"
            :src="getImageUrl(selectedTestimonial.url_foto_profil)"
            :alt="selectedTestimonial.nama"
            class="rounded-circle img-fluid"
            style="width: 120px; height: 120px; object-fit: cover"
            @error="handleImageError($event)"
          />
          <div
            v-else
            class="rounded-circle bg-secondary d-flex align-items-center justify-content-center mx-auto"
            style="width: 120px; height: 120px"
          >
            <svg
              width="60"
              height="60"
              viewBox="0 0 24 24"
              fill="none"
              stroke="white"
              stroke-width="2"
            >
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle
                cx="12"
                cy="7"
                r="4"
              ></circle>
            </svg>
          </div>
        </div>
        <h4 class="mb-2">{{ selectedTestimonial.nama }}</h4>
        <p class="text-primary fw-semibold mb-3">{{ selectedTestimonial.perusahaan || 'Tidak ada perusahaan' }}</p>

        <div
          v-if="selectedTestimonial.pesan"
          class="mb-4 text-start"
        >
          <h6 class="text-muted mb-2">Pesan:</h6>
          <p
            class="text-muted"
            style="line-height: 1.6; white-space: pre-line"
          >
            {{ selectedTestimonial.pesan }}
          </p>
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

const testimonials = ref([])
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

const normalizeTestimonial = t => ({
  id: t.id ?? null,
  nama: t.nama ?? '',
  perusahaan: t.perusahaan ?? '',
  pesan: t.pesan ?? '',
  url_foto_profil: t.url_foto_profil ?? '',
})

const fetchTestimonials = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/testimonis')
    const arr = Array.isArray(res.data.testimonials) ? res.data.testimonials : Array.isArray(res.data) ? res.data : []
    testimonials.value = arr.map(normalizeTestimonial)
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

const truncateDescription = (text, maxLength) => {
  if (!text) return 'Tidak ada pesan'
  if (text.length <= maxLength) return text
  return text.substring(0, maxLength) + '...'
}

const handleImageError = event => {
  event.target.style.display = 'none'
}

const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({
  id: null,
  nama: '',
  perusahaan: '',
  pesan: '',
  url_foto_profil: null,
})

const showDetail = ref(false)
const selectedTestimonial = ref(null)

const openDetail = testimonial => {
  selectedTestimonial.value = testimonial
  showDetail.value = true
}

const openCreate = () => {
  modalMode.value = 'create'
  Object.assign(form, {
    id: null,
    nama: '',
    perusahaan: '',
    pesan: '',
    url_foto_profil: null,
  })
  showModal.value = true
}

const startEdit = testimonial => {
  modalMode.value = 'edit'
  Object.assign(form, {
    id: testimonial.id,
    nama: testimonial.nama,
    perusahaan: testimonial.perusahaan,
    pesan: testimonial.pesan,
    url_foto_profil: null, // Reset file input
  })
  showModal.value = true
}

const submit = async () => {
  if (!form.nama) {
    showAlert('Nama wajib diisi', 'Validasi')
    return
  }
  if (!form.pesan) {
    showAlert('Pesan wajib diisi', 'Validasi')
    return
  }

  const fd = new FormData()
  fd.append('nama', form.nama)
  fd.append('perusahaan', form.perusahaan)
  fd.append('pesan', form.pesan)
  if (form.url_foto_profil) fd.append('url_foto_profil', form.url_foto_profil)

  try {
    if (modalMode.value === 'create') {
      await api.post('/testimonis', fd, {
        headers: { 'Content-Type': 'multipart/form-data' },
      })
      showModal.value = false
      await fetchTestimonials()
      showAlert('Testimoni berhasil ditambahkan!', 'Berhasil')
    } else if (modalMode.value === 'edit' && form.id) {
      await api.put(`/testimonis/${form.id}`, fd, {
        headers: { 'Content-Type': 'multipart/form-data' },
      })
      showModal.value = false
      await fetchTestimonials()
      showAlert('Testimoni berhasil diperbarui!', 'Berhasil')
    }
  } catch (err) {
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

const deleteTestimonial = async testimonial => {
  const confirmed = await showConfirm(`Hapus ${testimonial.nama}?`, 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/testimonis/${testimonial.id}`)
    await fetchTestimonials()
    showAlert('Testimoni berhasil dihapus!', 'Berhasil')
  } catch (err) {
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

onMounted(fetchTestimonials)
</script>

<style scoped>
.testimonial-card {
  border-radius: 12px;
  transition: transform 0.2s ease;
}
.testimonial-card:hover {
  transform: translateY(-5px);
}
.modal-backdrop {
  background: rgba(0, 0, 0, 0.35);
}
.testimonial-photo {
  border: 3px solid #f8f9fa;
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
