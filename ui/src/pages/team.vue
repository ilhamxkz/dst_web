<template>
  <section class="container py-5">
    <div class="d-flex justify-content-between mb-3">
      <h3>Data Team</h3>
      <div>
        <button class="btn btn-primary me-2" @click="openCreate">+ Tambah Team</button>
        <button class="btn btn-secondary" @click="fetchTeams">Refresh</button>
      </div>
    </div>

    <div v-if="loading" class="text-center py-4">
      <div class="spinner-border text-primary" role="status"></div>
      <p class="mt-2">Memuat data...</p>
    </div>

    <div v-else-if="error" class="alert alert-danger text-center">{{ error }}</div>

    <div v-else class="d-flex overflow-auto gap-3 pb-3">
      <div v-for="team in teams" :key="team.id"
        class="card text-center shadow-sm team-card flex-shrink-0"
        style="width: 280px; cursor: pointer"
        @click="openDetail(team)">
        <div class="card-body d-flex flex-column">
          <div class="mb-3">
            <img v-if="team.photo" 
              :src="getImageUrl(team.photo)" 
              :alt="team.name" 
              class="rounded-circle img-fluid team-photo mb-2" 
              style="width: 80px; height: 80px; object-fit: cover"
              @error="handleImageError($event)"
            />
            <div v-else class="rounded-circle bg-secondary d-flex align-items-center justify-content-center mx-auto mb-2" 
              style="width: 80px; height: 80px;">
              <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg>
            </div>
          </div>
          <h5 class="card-title">{{ team.name }}</h5>
          <p class="card-text text-primary fw-semibold">{{ team.position }}</p>
          <p class="card-text text-muted flex-grow-1 small">{{ truncateDescription(team.bio, 80) }}</p>
          
          <!-- Social Links -->
          <div class="d-flex justify-content-center gap-2 mb-3">
            <a v-if="team.linkedin" :href="team.linkedin" target="_blank" class="text-decoration-none">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="#0077B5">
                <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
              </svg>
            </a>
            <a v-if="team.github" :href="team.github" target="_blank" class="text-decoration-none">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="#333">
                <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
              </svg>
            </a>
          </div>
          
          <div class="mt-auto">
            <button class="btn btn-warning btn-sm me-2" @click.stop="startEdit(team)">Edit</button>
            <button class="btn btn-danger btn-sm" @click.stop="deleteTeam(team)">Hapus</button>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- Create/Edit Modal -->
  <div v-if="showModal" class="modal-backdrop" style="position:fixed;inset:0;display:flex;align-items:center;justify-content:center;z-index:1050">
    <div class="card p-3" style="width:500px;max-width:95vw">
      <div class="d-flex justify-content-between align-items-center mb-2">
        <h5 class="m-0">{{ modalMode==='create'?'Tambah Team':'Edit Team' }}</h5>
        <button class="btn btn-sm btn-outline-secondary" @click="showModal=false">×</button>
      </div>
      <div class="row g-2">
        <div class="col-12">
          <label class="form-label">Nama</label>
          <input class="form-control" v-model="form.name" required/>
        </div>
        <div class="col-12">
          <label class="form-label">Posisi</label>
          <input class="form-control" v-model="form.position"/>
        </div>
        <div class="col-12">
          <label class="form-label">Bio</label>
          <textarea class="form-control" rows="3" v-model="form.bio"></textarea>
        </div>
        <div class="col-12">
          <label class="form-label">Foto</label>
          <input type="file" class="form-control" accept="image/*" @change="e => form.photo = e.target.files[0]"/>
        </div>
        <div class="col-12">
          <label class="form-label">LinkedIn URL</label>
          <input type="url" class="form-control" v-model="form.linkedin" placeholder="https://linkedin.com/in/username"/>
        </div>
        <div class="col-12">
          <label class="form-label">GitHub URL</label>
          <input type="url" class="form-control" v-model="form.github" placeholder="https://github.com/username"/>
        </div>
      </div>
      <div class="d-flex justify-content-end mt-3">
        <button class="btn btn-secondary me-2" @click="showModal=false">Batal</button>
        <button class="btn btn-primary" @click="submit">{{ modalMode==='create'?'Simpan':'Update' }}</button>
      </div>
    </div>
  </div>

  <!-- Detail Modal -->
  <div v-if="showDetail && selectedTeam" class="modal-backdrop" style="position:fixed;inset:0;display:flex;align-items:center;justify-content:center;z-index:1050">
    <div class="card p-4" style="width:500px;max-width:95vw;max-height:90vh;overflow-y:auto">
      <div class="d-flex justify-content-between align-items-center mb-3">
        <h5 class="m-0">Detail Team Member</h5>
        <button class="btn btn-sm btn-outline-secondary" @click="showDetail=false">×</button>
      </div>
      <div class="text-center">
        <div class="mb-3">
          <img v-if="selectedTeam.photo" 
            :src="getImageUrl(selectedTeam.photo)" 
            :alt="selectedTeam.name"
            class="rounded-circle img-fluid"
            style="width: 120px; height: 120px; object-fit: cover"
            @error="handleImageError($event)"
          />
          <div v-else class="rounded-circle bg-secondary d-flex align-items-center justify-content-center mx-auto" 
            style="width: 120px; height: 120px;">
            <svg width="60" height="60" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle cx="12" cy="7" r="4"></circle>
            </svg>
          </div>
        </div>
        <h4 class="mb-2">{{ selectedTeam.name }}</h4>
        <p class="text-primary fw-semibold mb-3">{{ selectedTeam.position }}</p>
        
        <div v-if="selectedTeam.bio" class="mb-4 text-start">
          <h6 class="text-muted mb-2">Bio:</h6>
          <p class="text-muted" style="line-height: 1.6; white-space: pre-line;">{{ selectedTeam.bio }}</p>
        </div>
        
        <!-- Social Links in Detail -->
        <div v-if="selectedTeam.linkedin || selectedTeam.github" class="d-flex justify-content-center gap-3 mb-3">
          <a v-if="selectedTeam.linkedin" :href="selectedTeam.linkedin" target="_blank" 
            class="btn btn-outline-primary btn-sm text-decoration-none">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor" class="me-1">
              <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
            </svg>
            LinkedIn
          </a>
          <a v-if="selectedTeam.github" :href="selectedTeam.github" target="_blank" 
            class="btn btn-outline-dark btn-sm text-decoration-none">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor" class="me-1">
              <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
            </svg>
            GitHub
          </a>
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
import 'bootstrap/dist/css/bootstrap.min.css'

const teams = ref([])
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

const normalizeTeam = t => ({
  id: t.id ?? null,
  name: t.name ?? '',
  position: t.position ?? '',
  bio: t.bio ?? '',
  photo: t.photo ?? '',
  linkedin: t.linkedin ?? '',
  github: t.github ?? ''
})

const fetchTeams = async () => {
  loading.value = true
  error.value = null
  try {
    const res = await api.get('/team_members')
    const arr = Array.isArray(res.data.teams) ? res.data.teams : Array.isArray(res.data) ? res.data : []
    teams.value = arr.map(normalizeTeam)
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

const truncateDescription = (text, maxLength) => {
  if (!text) return 'Tidak ada bio'
  if (text.length <= maxLength) return text
  return text.substring(0, maxLength) + '...'
}

const handleImageError = (event) => {
  event.target.style.display = 'none'
}

const showModal = ref(false)
const modalMode = ref('create')
const form = reactive({ 
  id: null, 
  name: '', 
  position: '', 
  bio: '', 
  photo: null, 
  linkedin: '', 
  github: '' 
})

const showDetail = ref(false)
const selectedTeam = ref(null)

const openDetail = team => {
  selectedTeam.value = team
  showDetail.value = true
}

const openCreate = () => {
  modalMode.value = 'create'
  Object.assign(form, {
    id: null,
    name: '',
    position: '',
    bio: '',
    photo: null,
    linkedin: '',
    github: ''
  })
  showModal.value = true
}

const startEdit = team => {
  modalMode.value = 'edit'
  Object.assign(form, {
    id: team.id,
    name: team.name,
    position: team.position,
    bio: team.bio,
    photo: null, // Reset file input
    linkedin: team.linkedin,
    github: team.github
  })
  showModal.value = true
}

const submit = async () => {
  if (!form.name) {
    showAlert('Nama team member wajib diisi', 'Validasi')
    return
  }

  const fd = new FormData()
  fd.append('name', form.name)
  fd.append('position', form.position)
  fd.append('bio', form.bio)
  fd.append('linkedin', form.linkedin)
  fd.append('github', form.github)
  if (form.photo) fd.append('photo', form.photo)

  try {
    if (modalMode.value === 'create') {
      await api.post('/team_members', fd, { 
        headers: { 'Content-Type': 'multipart/form-data' } 
      })
      showModal.value = false
      await fetchTeams()
      showAlert('Team member berhasil ditambahkan!', 'Berhasil')
    } else if (modalMode.value === 'edit' && form.id) {
      await api.put(`/team_members/${form.id}`, fd, { 
        headers: { 'Content-Type': 'multipart/form-data' } 
      })
      showModal.value = false
      await fetchTeams()
      showAlert('Team member berhasil diperbarui!', 'Berhasil')
    }
  } catch (err) {
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

const deleteTeam = async team => {
  const confirmed = await showConfirm(`Hapus ${team.name}?`, 'Konfirmasi Hapus')
  if (!confirmed) return
  try {
    await api.delete(`/team_members/${team.id}`)
    await fetchTeams()
    showAlert('Team member berhasil dihapus!', 'Berhasil')
  } catch (err) {
    showAlert(err.response?.data?.error || err.message, 'Error')
  }
}

onMounted(fetchTeams)
</script>

<style scoped>
.team-card { 
  border-radius: 12px; 
  transition: transform 0.2s ease; 
}
.team-card:hover { 
  transform: translateY(-5px); 
}
.modal-backdrop { 
  background: rgba(0,0,0,0.35); 
}
.team-photo {
  border: 3px solid #f8f9fa;
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
