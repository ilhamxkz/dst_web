<template>
  <div class="our-team-page">
    <div class="hero-section">
      <div class="container">
        <h1>TEAM</h1>
        <p>Temui para profesional berbakat di balik kesuksesan DAYA SINERGI</p>
      </div>
    </div>

    <div class="content-section">
      <div class="container">
        <div v-if="loading" class="loading">Loading team...</div>
        <div v-else-if="error" class="error">{{ error }}</div>
        <div v-else class="team-grid">
          <div v-for="member in teamMembers" :key="member.id" class="team-member">
            <div class="member-avatar" :class="{ 'has-image': !!member.photo }">
              <img v-if="member.photo" :src="getPhotoUrl(member.photo)" alt="avatar" />
              <i v-else class="ri-user-line"></i>
            </div>
            <h3>{{ member.name }}</h3>
            <p v-if="member.position">{{ member.position }}</p>
            <p v-if="member.bio">{{ member.bio }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'OurTeam',
  data() {
    return {
      teamMembers: [],
      loading: true,
      error: '',
    };
  },
  methods: {
    getApiBaseUrl() {
      return import.meta?.env?.VITE_API_BASE_URL || 'http://localhost:8080';
    },
    getPhotoUrl(photoPath) {
      if (!photoPath) return '';
      const base = this.getApiBaseUrl();
      if (photoPath.startsWith('http://') || photoPath.startsWith('https://')) return photoPath;
      return `${base}/${photoPath}`.replace(/\\/g, '/');
    },
    async fetchTeamMembers() {
      this.loading = true;
      this.error = '';
      try {
        const base = this.getApiBaseUrl();
        const res = await axios.get(`${base}/api/team_members`);
        const teams = res?.data?.teams || [];
        this.teamMembers = teams.map((t) => ({
          id: t.id,
          name: t.name,
          position: t.position || '',
          bio: t.bio || '',
          photo: t.photo || '',
        }));
      } catch (e) {
        this.error = 'Gagal memuat data tim.';
      } finally {
        this.loading = false;
      }
    },
  },
  mounted() {
    this.fetchTeamMembers();
  },
};
</script>

<style scoped>
.our-team-page {
  min-height: 100vh;
  background: #000;
  color: #fff;
}

.hero-section {
  padding: 120px 0 80px;
  text-align: center;
  background: linear-gradient(135deg, #0f3a8c 0%, #3b82f6 100%);
}

.hero-section h1 {
  font-size: 3rem;
  margin-bottom: 20px;
  font-weight: 700;
}

.hero-section p {
  font-size: 1.2rem;
  opacity: 0.9;
  max-width: 600px;
  margin: 0 auto;
}

.content-section {
  padding: 80px 0;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.team-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 40px;
  margin-top: 40px;
}

.team-member {
  background: #111;
  border-radius: 12px;
  padding: 30px;
  text-align: center;
  transition: transform 0.3s ease;
}

.team-member:hover {
  transform: translateY(-5px);
}

.member-avatar {
  width: 80px;
  height: 80px;
  background: #3b82f6;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
  overflow: hidden;
}

.member-avatar i {
  font-size: 2rem;
  color: #fff;
}

.member-avatar.has-image {
  background: transparent;
}

.member-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.team-member h3 {
  font-size: 1.5rem;
  margin-bottom: 5px;
  color: #fff;
}

.team-member p:first-of-type {
  color: #3b82f6;
  font-weight: 600;
  margin-bottom: 15px;
}

.team-member p:last-of-type {
  color: #999;
  line-height: 1.6;
}

@media (max-width: 768px) {
  .hero-section h1 {
    font-size: 2rem;
  }

  .team-grid {
    grid-template-columns: 1fr;
  }
}
</style>
