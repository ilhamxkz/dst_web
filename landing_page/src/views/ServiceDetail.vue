<template>
  <div class="service-detail-page">
    <div class="container">
      <button @click="goBack" class="back-button">
        <i class="ri-arrow-left-line"></i> Kembali
      </button>

      <div v-if="loading" class="loading">Memuat data layanan...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else-if="service" class="experience-card">
        <div class="card-image" v-if="service.images || service.icon">
          <img :src="getImageUrl(service.images || service.icon)" :alt="service.title || 'Layanan'" />
        </div>

        <div class="card-content">
          <h2>{{ service.title || 'Layanan' }}</h2>

          <div
            class="description"
            v-html="service.description || '<p>Informasi detail layanan ini belum tersedia.</p>'"
          ></div>

          <ul class="info-list">
            <li v-if="service.category"><strong>Kategori:</strong> {{ service.category }}</li>
            <li v-if="service.created_at"><strong>Dibuat pada:</strong> {{ formatDateTime(service.created_at) }}</li>
          </ul>
        </div>
      </div>
      <div v-else class="error">Data layanan tidak ditemukan.</div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ServiceDetail',
  data() {
    return {
      service: null,
      loading: true,
      error: '',
    }
  },
  methods: {
    getApiBase() {
      return import.meta?.env?.VITE_API_BASE_URL || 'http://localhost:8080'
    },

    async fetchService() {
      this.loading = true
      this.error = ''
      try {
        const id = this.$route.params.id
        const base = this.getApiBase()
        const res = await axios.get(`${base}/api/services/${id}`)
        const raw = res?.data?.service || res?.data?.data || res?.data
        this.service = raw || null

        if (!this.service) {
          this.error = 'Data layanan tidak ditemukan.'
        }
      } catch (err) {
        console.error('Error fetching service:', err)
        this.error = err.response?.data?.message || 'Gagal memuat data layanan.'
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
    this.fetchService()
  },
  watch: {
    '$route.params.id'() {
      this.fetchService()
    }
  }
}
</script>

<style scoped>
.service-detail-page {
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
  height: 200px;
  object-fit: contain;
  object-position: center;
  background-color: #ffffff;
  border-radius: 10px;
  display: block;
  padding: 10px;
  box-sizing: border-box;
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
