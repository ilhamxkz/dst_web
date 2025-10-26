<template>
  <div class="pelatihan-detail-page">
    <div class="container">
      <button @click="goBack" class="back-button">
        <i class="ri-arrow-left-line"></i> Kembali
      </button>

      <div v-if="loading" class="loading">Memuat data pelatihan...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else-if="item" class="experience-card">
        <div class="card-image" v-if="item.image">
          <img :src="getImageUrl(item.image)" :alt="item.category || 'Pelatihan'" />
        </div>

        <div class="card-content">
          <h2>{{ item.category || 'Pelatihan' }}</h2>

          <!-- Deskripsi pelatihan -->
          <div
            class="description"
            v-html="item.deskripsi || item.description || '<p>Informasi detail pelatihan ini belum tersedia.</p>'"
          ></div>

          <ul class="info-list">
            <li v-if="item.level"><strong>Level:</strong> {{ item.level }}</li>
            <li v-if="item.duration"><strong>Durasi:</strong> {{ item.duration }}</li>
            <li v-if="item.location"><strong>Lokasi:</strong> {{ item.location }}</li>
            <li v-if="item.price"><strong>Harga:</strong> {{ formatPrice(item.price) }}</li>
          </ul>

          <p class="created-at">
            <small>Dibuat pada: {{ formatDateTime(item.created_at) }}</small>
          </p>
        </div>
      </div>
      <div v-else class="error">Data pelatihan tidak ditemukan.</div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'PelatihanDetail',
  data() {
    return {
      item: null,
      loading: true,
      error: '',
    }
  },
  methods: {
    getApiBase() {
      return import.meta?.env?.VITE_API_BASE_URL || 'http://localhost:8080';
    },

    async fetchPelatihan() {
      this.loading = true
      this.error = ''
      try {
        const id = this.$route.params.id
        const base = this.getApiBase()
        const res = await axios.get(`${base}/api/pelatihan/${id}`)

        // Coba beberapa kemungkinan struktur payload
        const raw = res?.data?.pelatihan || res?.data?.data || res?.data
        this.item = raw || null

        if (!this.item) {
          this.error = 'Data pelatihan tidak ditemukan.'
        }
      } catch (err) {
        console.error('Error fetching pelatihan:', err)
        this.error = err.response?.data?.message || 'Gagal memuat data pelatihan.'
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

    formatPrice(value) {
      if (value === undefined || value === null || value === '') return '-'
      const num = Number(value)
      if (Number.isNaN(num)) return String(value)
      return num.toLocaleString('id-ID', { style: 'currency', currency: 'IDR' })
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
    this.fetchPelatihan()
  },
  watch: {
    '$route.params.id'() {
      this.fetchPelatihan()
    }
  }
}
</script>

<style scoped>
.pelatihan-detail-page {
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


