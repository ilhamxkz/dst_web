<template>
  <div class="pelatihan-container">
    <h1 class="title">Manajemen Pelatihan</h1>

    <!-- Form Tambah / Edit -->
    <form @submit.prevent="savePelatihan" enctype="multipart/form-data" class="form-box">
      <div class="form-left">
        <label>Upload Gambar</label>
        <input type="file" @change="onFileChange" accept="image/*" />
        <div v-if="preview" class="preview">
          <img :src="preview" alt="Preview" />
        </div>
      </div>

      <div class="form-right">
        <label>Kategori</label>
        <input v-model="form.category" type="text" placeholder="Masukkan kategori" required />

        <label>Jenis</label>
        <input v-model="form.jenis" type="text" placeholder="Masukkan jenis pelatihan" required />

        <label>Deskripsi</label>
        <textarea v-model="form.deskripsi" rows="3" placeholder="Masukkan deskripsi" required></textarea>

        <div class="button-group">
          <button type="submit" class="btn-primary">
            {{ editMode ? 'Perbarui Data' : 'Tambah Data' }}
          </button>
          <button v-if="editMode" type="button" @click="cancelEdit" class="btn-secondary">
            Batal
          </button>
        </div>
      </div>
    </form>

    <!-- Pesan Status -->
    <p v-if="loading" class="info-text">Memuat data pelatihan...</p>
    <p v-if="error" class="error-text">{{ error }}</p>

    <!-- Daftar Pelatihan -->
    <div class="grid-box">
      <div v-for="item in pelatihan" :key="item.id" class="card">
        <img :src="getImageUrl(item.image)" alt="Pelatihan" class="card-image" />
        <div class="card-body">
          <h3>{{ item.category }}</h3>
          <p v-if="item.jenis" class="jenis-badge">{{ item.jenis }}</p>
          <p>{{ item.deskripsi }}</p>
          <p v-if="item.created_at" class="created-at">
            <small>Dibuat: {{ formatDateTime(item.created_at) }}</small>
          </p>
          <div class="card-buttons">
            <button @click="editPelatihan(item)" class="btn-warning">Edit</button>
            <button @click="deletePelatihan(item.id)" class="btn-danger">Hapus</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'PelatihanPage',
  data() {
    return {
      pelatihan: [],
      loading: false,
      error: '',
      editMode: false,
      preview: null,
      selectedFile: null,
      form: {
        id: null,
        image: '',
        category: '',
        jenis: '',
        deskripsi: ''
      }
    }
  },
  mounted() {
    this.fetchPelatihan()
  },
  methods: {
    getApiBase() {
      return import.meta.env.VITE_API_BASE || 'http://localhost:8080'
    },

    async fetchPelatihan() {
      this.loading = true
      this.error = ''
      try {
        const base = this.getApiBase()
        const res = await axios.get(`${base}/api/pelatihan`)
        this.pelatihan = res.data.pelatihan || res.data
      } catch {
        this.error = 'Gagal memuat data pelatihan.'
      } finally {
        this.loading = false
      }
    },

    onFileChange(e) {
      const file = e.target.files[0]
      if (file) {
        this.selectedFile = file
        this.preview = URL.createObjectURL(file)
      }
    },

    async savePelatihan() {
      const base = this.getApiBase()
      const formData = new FormData()
      formData.append('category', this.form.category)
      formData.append('jenis', this.form.jenis)
      formData.append('deskripsi', this.form.deskripsi)
      if (this.selectedFile) formData.append('image', this.selectedFile)

      try {
        if (this.editMode) {
          await axios.put(`${base}/api/pelatihan/${this.form.id}`, formData)
        } else {
          await axios.post(`${base}/api/pelatihan`, formData)
        }
        this.resetForm()
        this.fetchPelatihan()
      } catch {
        this.error = 'Gagal menyimpan data pelatihan.'
      }
    },

    editPelatihan(item) {
      this.form.id = item.id
      this.form.category = item.category
      this.form.jenis = item.jenis || ''
      this.form.deskripsi = item.deskripsi
      this.form.image = item.image
      this.preview = this.getImageUrl(item.image)
      this.editMode = true
    },

    cancelEdit() {
      this.resetForm()
    },

    async deletePelatihan(id) {
      if (!confirm('Yakin ingin menghapus data ini?')) return
      const base = this.getApiBase()
      try {
        await axios.delete(`${base}/api/pelatihan/${id}`)
        this.fetchPelatihan()
      } catch {
        this.error = 'Gagal menghapus data pelatihan.'
      }
    },

    resetForm() {
      this.form = { id: null, image: '', category: '', jenis: '', deskripsi: '' }
      this.preview = null
      this.selectedFile = null
      this.editMode = false
    },

    getImageUrl(path) {
      if (!path) return ''
      const base = this.getApiBase()
      return `${base}/${path}`
    },

    formatDateTime(datetime) {
      if (!datetime) return '-'
      const d = new Date(datetime)
      return d.toLocaleString('id-ID')
    }
  }
}
</script>

<style scoped>
.pelatihan-container {
  padding: 24px;
  background: #f8f9fb;
  min-height: 100vh;
}

.title {
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 20px;
}

.form-box {
  display: flex;
  flex-wrap: wrap;
  gap: 24px;
  background: #fff;
  padding: 16px;
  border-radius: 12px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
  margin-bottom: 32px;
}

.form-left,
.form-right {
  flex: 1;
  min-width: 280px;
}

label {
  display: block;
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 6px;
  color: #333;
}

input[type="text"],
textarea,
input[type="file"] {
  width: 100%;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 8px;
  font-size: 14px;
  background-color: #fff;
}

.preview {
  margin-top: 10px;
}

.preview img {
  width: 120px;
  height: 120px;
  object-fit: cover;
  border-radius: 8px;
  border: 1px solid #ddd;
}

.button-group {
  display: flex;
  gap: 10px;
  margin-top: 12px;
}

.btn-primary {
  background: #2563eb;
  color: #fff;
  padding: 8px 14px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  border: none;
}

.btn-primary:hover {
  background: #1e40af;
}

.btn-secondary {
  background: #6b7280;
  color: #fff;
  padding: 8px 14px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  border: none;
}

.btn-secondary:hover {
  background: #4b5563;
}

.info-text {
  color: #555;
  font-size: 14px;
  margin-bottom: 10px;
}

.error-text {
  color: #e11d48;
  font-weight: 500;
  margin-bottom: 10px;
}

.grid-box {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 16px;
}

.card {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
  overflow: hidden;
  display: flex;
  flex-direction: column;
  transition: transform 0.2s;
}

.card:hover {
  transform: translateY(-3px);
}

.card-image {
  width: 100%;
  height: 180px;
  object-fit: cover;
}

.card-body {
  padding: 12px;
}

.card-body h3 {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 8px;
}

.card-body p {
  font-size: 14px;
  color: #555;
  margin-bottom: 12px;
}

.jenis-badge {
  display: inline-block;
  background: #e6eefb;
  color: #1e3a8a;
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
  margin-bottom: 8px;
}

.created-at {
  font-size: 12px !important;
  color: #6b7280 !important;
  margin-bottom: 8px !important;
}

.card-buttons {
  display: flex;
  gap: 8px;
}

.btn-warning {
  background: #f59e0b;
  color: #fff;
  border: none;
  padding: 6px 10px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 13px;
}

.btn-warning:hover {
  background: #d97706;
}

.btn-danger {
  background: #dc2626;
  color: #fff;
  border: none;
  padding: 6px 10px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 13px;
}

.btn-danger:hover {
  background: #b91c1c;
}

</style>
