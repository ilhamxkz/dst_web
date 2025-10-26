<template>
  <div class="industry-detail-page">
    <div class="container">
      <button @click="goBack" class="back-button">
        <i class="ri-arrow-left-line"></i> Kembali
      </button>

      <div v-if="loading" class="loading">Memuat data project...</div>
      <div v-else-if="error" class="error">{{ error }}</div>

      <div v-else-if="experience" class="experience-card">
        <div class="card-image" v-if="experience.cover_image">
          <img :src="getImageUrl(experience.cover_image)" :alt="experience.title" />
        </div>

        <div class="card-content">
          <h2>{{ experience.title }}</h2>

          <!-- ✅ perbaikan di sini -->
          <div
            class="description"
            v-html="experience.description || '<p>Informasi detail project ini belum tersedia.</p>'"
          ></div>

          <ul class="info-list">
            <li v-if="experience.client_name"><strong>Client:</strong> {{ experience.client_name }}</li>
            <li><strong>Mulai:</strong> {{ formatDate(experience.start_date) }}</li>
            <li><strong>Selesai:</strong> {{ formatDate(experience.end_date) }}</li>
            <li><strong>Status:</strong> {{ experience.status }}</li>
          </ul>

          <p class="created-at">
            <small>Dibuat pada: {{ formatDateTime(experience.created_at) }}</small>
          </p>
        </div>
      </div>

      <div v-else class="error">Data project tidak ditemukan.</div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'IndustryDetail',
  data() {
    return {
      experience: null,
      loading: true,
      error: '',
    }
  },
  methods: {
    getApiBase() {
      return import.meta?.env?.VITE_API_BASE_URL || 'http://localhost:8080';
    },

    async fetchExperience() {
      this.loading = true
      this.error = ''
      try {
        const id = this.$route.params.id
        const base = this.getApiBase()
        const res = await axios.get(`${base}/api/experiences/${id}`)
        this.experience = res.data.experience || res.data

        if (!this.experience) {
          this.error = 'Data project tidak ditemukan.'
        }
      } catch (err) {
        console.error('Error fetching experience:', err)
        this.error = err.response?.data?.message || 'Gagal memuat data project.'
      } finally {
        this.loading = false
      }
    },

    getImageUrl(filePath) {
      if (!filePath) return ''
      const base = this.getApiBase()
      if (filePath.startsWith('http://') || filePath.startsWith('https://')) {
        return filePath
      }
      return `${base}/${filePath}`.replace(/\\/g, '/')
    },

    formatDate(date) {
      if (!date) return '-'
      const d = new Date(date)
      return d.toLocaleDateString('id-ID', { year: 'numeric', month: 'long', day: 'numeric' })
    },

    formatDateTime(datetime) {
      if (!datetime) return '-'
      const d = new Date(datetime)
      return d.toLocaleString('id-ID')
    },

    goBack() {
      if (window.history.length > 1) {
        this.$router.go(-1)
      } else {
        this.$router.push('/')
      }
    },
  },
  mounted() {
    this.fetchExperience()
  },
  watch: {
    '$route.params.id'() {
      this.fetchExperience()
    }
  }
}
</script>

<style scoped>
.industry-detail-page {
  background-color: #f9fafb;
  min-height: 100vh;
  padding: 60px 0;
  color: #333;
}

.container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 0 20px;
}

.back-button {
  background: none;
  border: none;
  color: #2563eb;
  font-weight: 600;
  font-size: 16px;
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 30px;
  cursor: pointer;
  transition: color 0.3s;
}
.back-button:hover {
  color: #1d4ed8;
}

.experience-card {
  display: flex;
  flex-wrap: wrap;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.08);
  overflow: hidden;
  transition: transform 0.3s ease;
  padding: 30px;
}
.experience-card:hover {
  transform: translateY(-4px);
}

.card-image {
  flex: 1 1 250px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 20px;
}
.card-image img {
   width: 100%;
  height: 200px; /* tinggi tetap agar sejajar */
  object-fit: contain; /* gambar menyesuaikan tanpa terpotong */
  object-position: center; /* posisi gambar di tengah */
  background-color: #ffffff; /* latar belakang putih biar netral */
  border-radius: 10px; /* opsional, biar sudut gambar halus */
  display: block;
  padding: 10px; /* beri jarak kecil agar gambar tidak terlalu mepet */
  box-sizing: border-box;;
}

.card-content {
  flex: 2 1 500px;
  padding-left: 30px;
}

.card-content h2 {
  color: #1e3a8a;
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 15px;
}

.description {
  color: #4b5563;
  font-size: 1.1rem;
  line-height: 1.7;
  margin-bottom: 20px;
}

.info-list {
  list-style-type: disc;
  margin-left: 20px;
  color: #1e40af;
  font-size: 1rem;
  margin-bottom: 25px;
}
.info-list li {
  margin-bottom: 8px;
  color: #374151;
}

.created-at {
  font-size: 0.9rem;
  color: #6b7280;
}

.loading,
.error {
  text-align: center;
  font-size: 1.2rem;
  padding: 60px 20px;
}

.error {
  color: #dc2626;
}

@media (max-width: 768px) {
  .experience-card {
    flex-direction: column;
    align-items: center;
    text-align: center;
  }
  .card-content {
    padding-left: 0;
  }
  .card-image img {
    width: 180px;
    height: 180px;
  }
}
</style>
