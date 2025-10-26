<template>
  <div class="home">
    <!-- Hero Section -->
    <!-- Dynamic background cycling with fade transitions -->
    <section class="hero">
      <!-- Background images with fade transitions -->
      <div v-for="(bg, index) in heroBackgrounds" :key="bg.id" class="hero-bg-image" :class="{ active: currentBgIndex === index }" :style="{ backgroundImage: `url(${getImageUrl(bg.image)})` }"></div>

      <div class="container hero-container">
        <div class="hero-content">
          <p class="hero-subtitle">IT TRANSFORMATION & MODERNIZATION</p>
          <h1 class="hero-title">Transform Your Business With Our IT Services</h1>
          <div class="hero-buttons">
            <button class="btn btn-primary">Learn more <i class="ri-arrow-right-s-line"></i></button>
            <button class="btn btn-secondary">View All Services <i class="ri-arrow-right-s-line"></i></button>
          </div>
        </div>
      </div>
    </section>

    <!-- Client Logos (auto sliding left 1 item every 1s) -->
    <section class="clients">
      <div class="container">
        <div class="client-logos-wrap">
          <div class="client-logos-track" ref="track">
            <div class="client-logo" v-for="(logo, idx) in duplicatedLogos" :key="idx + '-' + logo">
              {{ logo }}
            </div>
          </div>
        </div>
      </div>
    </section>

    <section id="layanan" ref="layananSection" class="what-we-do">
  <div class="container">
    <div class="section-header">
      <span class="section-badge">WHAT WE DO</span>
      <h2 class="section-title light">Mendorong Pertumbuhan Bisnis Anda dengan Keahlian dan Solusi Inovatif Kami</h2>
    </div>

    <div v-if="servicesLoading" class="loading">Loading services...</div>
    <div v-else-if="servicesError" class="error">{{ servicesError }}</div>

    <!-- Slider layanan -->
    <div class="service-slider-wrap">
      <div class="service-slider-track" ref="serviceTrack">
        <div
          v-for="(service, idx) in duplicatedServices"
          :key="idx + '-' + service.id"
          class="service-card"
          @click="goToService(service.id)"
        >
          <div class="service-icon green">
            <img
              v-if="service.icon"
              :src="getPhotoUrl(service.icon)"
              alt="icon"
              style="width: 48px; height: 48px; object-fit: contain"
            />
            <i v-else class="ri-flashlight-line"></i>
          </div>
          <h3>{{ service.title }}</h3>
          <p>{{ truncateText(service.description, 40) || 'Deskripsi layanan belum tersedia.' }}</p>
        </div>
      </div>
    </div>
  </div>
</section>


    <!-- Capabilities Section -->
    <section class="capabilities">
      <div class="container">
        <div class="capabilities-grid">
          <div class="capabilities-content">
            <span class="section-badge light">OUR CAPABILITIES</span>
            <h2 class="capabilities-title"><span class="text-blue">Fuel your business</span> with the technology you need to succeed</h2>
          </div>
          <div class="capabilities-list">
            <div class="capability-item">
              <div class="capability-icon red"><i class="ri-save-line"></i></div>
              <div class="capability-content">
                <h3>Backup & Recovery</h3>
                <p>Lorem ipsum dolor sit amet consectetur adipiscing elit.</p>
                <a href="#" class="learn-more">Learn More <i class="ri-arrow-right-s-line"></i></a>
              </div>
            </div>
            <div class="capability-item">
              <div class="capability-icon blue"><i class="ri-brain-line"></i></div>
              <div class="capability-content">
                <h3>AI & Machine Learning</h3>
                <p>Lorem ipsum dolor sit amet consectetur adipiscing elit.</p>
                <a href="#" class="learn-more">Learn More <i class="ri-arrow-right-s-line"></i></a>
              </div>
            </div>
            <div class="capability-item">
              <div class="capability-icon green"><i class="ri-shield-check-line"></i></div>
              <div class="capability-content">
                <h3>Cybersecurity</h3>
                <p>Lorem ipsum dolor sit amet consectetur adipiscing elit.</p>
                <a href="#" class="learn-more">Learn More <i class="ri-arrow-right-s-line"></i></a>
              </div>
            </div>
            <div class="capability-item">
              <div class="capability-icon light-blue"><i class="ri-cloud-line"></i></div>
              <div class="capability-content">
                <h3>Cloud & DevOps</h3>
                <p>Lorem ipsum dolor sit amet consectetur adipiscing elit.</p>
                <a href="#" class="learn-more light">Learn More <i class="ri-arrow-right-s-line"></i></a>
              </div>
            </div>
            <div class="capability-item">
              <div class="capability-icon purple"><i class="ri-smartphone-line"></i></div>
              <div class="capability-content">
                <h3>Product Design</h3>
                <p>Lorem ipsum dolor sit amet consectetur adipiscing elit.</p>
                <a href="#" class="learn-more">Learn More <i class="ri-arrow-right-s-line"></i></a>
              </div>
            </div>
            <div class="capability-item">
              <div class="capability-icon info-blue"><i class="ri-briefcase-line"></i></div>
              <div class="capability-content">
                <h3>Consulting Services</h3>
                <p>Lorem ipsum dolor sit amet consectetur adipiscing elit.</p>
                <a href="#" class="learn-more">Learn More <i class="ri-arrow-right-s-line"></i></a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Stats Section -->
    <section class="stats">
      <div class="container">
        <h2 class="stats-title">Kenapa Memilih DAYA SINERGI?</h2>
        <div class="stats-grid">
          <div class="stat-card">
            <h4>Proyek Berhasil</h4>
            <p class="stat-number">54<span class="blue">K</span><span class="plus">+</span></p>
          </div>
          <div class="stat-card">
            <h4>Klien Puas</h4>
            <p class="stat-number">99<span class="blue">%</span></p>
          </div>
          <div class="stat-card">
            <h4>Tahun Pengalaman</h4>
            <p class="stat-number">20<span class="plus">+</span></p>
          </div>
          <div class="stat-card">
            <h4>Ahli Teknologi</h4>
            <p class="stat-number">88<span class="plus">+</span></p>
          </div>
        </div>
      </div>
    </section>

    <!-- About Section -->
    <section class="about">
      <div class="container about-container">
        <div class="about-visual"></div>
        <div class="about-content">
          <span class="section-badge">ABOUT US</span>
          <h2>Build your business on a strong IT Infrastructure</h2>
          <button class="btn btn-primary">Read More</button>
        </div>
      </div>
    </section>

    <!-- Our Team Section -->
    <section class="our-team" id="team">
      <div class="container">
        <div class="section-header">
          <h2 class="section-title light">TEAM</h2>
          <p class="team-subtitle">Temui para profesional berbakat di balik kesuksesan DAYA SINERGI</p>
        </div>
        <div v-if="teamLoading" class="loading">Loading team...</div>
        <div v-else-if="teamError" class="error">{{ teamError }}</div>
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
    </section>

  <!-- === Industries Section === -->
  <section id="industries" class="industries text-center">
    <div class="container">
      <h2 class="section-title light mb-5">
        Temukan keahlian mendalam kami di industri Anda
      </h2>

      <div class="year-filter">
  <button
    v-for="(year, index) in availableYears"
    :key="year"
    :class="['year-btn', { active: selectedYear === year }]"
    @click="selectYear(year)"
  >
    {{ year }}
    <span v-if="index < availableYears.length - 1" class="separator">|</span>
  </button>
</div>

      <!-- === Grid Industries === -->
      <div v-if="expError" class="text-danger mb-4">{{ expError }}</div>

      <div class="testimonials-grid">
        <div
          v-for="exp in visibleExperiences"
          :key="exp.id"
          class="testimonial-card"
          @click="goToIndustry(exp.id)"
        >
          <div class="testimonial-avatar">
            <img
              :src="getImageUrl(exp.logo || 'uploads/icons/dst.png')"
              alt="logo"
              class="avatar-image"
            />
          </div>
          <h4>{{ exp.title || "Untitled Experience" }}</h4>
          <div class="testimonial-role">{{ exp.year || "Tahun tidak tersedia" }}</div>
        </div>
      </div>

      <!-- Tombol Lihat Selengkapnya -->
      <div class="see-more-wrapper">
        <button v-if="!showAll" @click="showAll = true" class="see-more-btn">
          Lihat Selengkapnya
        </button>
        <button v-else @click="showAll = false" class="see-more-btn">
          Tutup
        </button>
      </div>
    </div>
  </section>

    <!-- Testimonials Section -->
    <section class="testimonials" id="testimonials">
      <div class="container">
        <span class="section-badge">Testimonials</span>
        <h2 class="section-title">Cepat, handal, dan layanan pelanggan yang dapat Anda andalkan.</h2>
        <div v-if="testimonialsLoading" class="loading">Loading testimonials...</div>
        <div v-else-if="testimonialsError" class="error">{{ testimonialsError }}</div>
        <div v-else class="testimonials-grid">
          <div v-for="testimonial in testimonials" :key="testimonial.id" class="testimonial-card">
            <div class="testimonial-avatar">
              <img v-if="testimonial.url_foto_profil && testimonial.url_foto_profil !== ''" :src="getPhotoUrl(testimonial.url_foto_profil)" :alt="testimonial.nama" @error="handleImageError" style="display: block" />
              <i class="ri-user-line" :style="{ display: !testimonial.url_foto_profil || testimonial.url_foto_profil === '' ? 'flex' : 'none' }"></i>
            </div>
            <h4>{{ testimonial.nama }}</h4>
            <p class="testimonial-role">{{ testimonial.perusahaan || 'Client' }}</p>
            <p class="testimonial-text">{{ testimonial.pesan }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- FAQ Section -->
    <section class="faq" id="faq">
      <div class="container faq-container">
        <div class="faq-sidebar">
          <span class="section-badge">FAQ</span>
          <h2>Questions</h2>
          <p>Butuh jawaban? Kami siap membantu!</p>
        </div>
        <div class="faq-list">
          <div v-if="faqLoading" class="text-center py-3">
            <div class="spinner-border text-primary" role="status"></div>
            <p class="mt-2">Memuat FAQ...</p>
          </div>
          <div v-else-if="faqError" class="alert alert-danger">{{ faqError }}</div>

          <div class="faq-item" v-for="(faq, index) in faqs" :key="faq.id || index">
            <div class="faq-question" @click="toggleFaq(index)">
              <div class="faq-question-text">{{ faq.question }}</div>
              <div class="faq-toggle">{{ faq.open ? '▲' : '▼' }}</div>
            </div>
            <div class="faq-answer" v-if="faq.open" v-html="faq.answer"></div>
          </div>
        </div>
      </div>
    </section>

    <!-- Blog Section -->
    <section class="pelatihan" id="pelatihan">
      <div class="container">
        <span class="section-badge">Pelatihan</span>
        <h2 class="section-title">Program Pelatihan Kami</h2>

        <div v-if="pelatihanLoading" class="loading">Memuat data pelatihan...</div>
        <div v-else-if="pelatihanError" class="error">{{ pelatihanError }}</div>

         <!-- Slider untuk menampilkan data pelatihan -->
         <div class="pelatihan-slider-wrap">
           <div class="pelatihan-slider-track" ref="pelatihanTrack">
             <div v-for="(p, idx) in duplicatedPelatihan" :key="idx + '-' + p.id" class="pelatihan-card" @click="goToPelatihan(p.id)" style="cursor: pointer" role="button" :aria-label="`Buka detail pelatihan ${p.category || ''}`">
               <div class="pelatihan-image">
                 <div class="pelatihan-fallback" :class="{ show: !p.image }"></div>
                 <img v-if="p.image" :src="getPhotoUrl(p.image)" alt="Pelatihan" class="pelatihan-img" @error="handlePelatihanImageError" />
               </div>
               <span class="pelatihan-category">{{ p.category }}</span>
               <span class="pelatihan-jenis">{{ p.jenis || 'Pelatihan' }}</span>
             </div>
           </div>
         </div>
      </div>
    </section>

    <!-- Location Section -->
    <section class="location py-5" id="maps">
      <div class="container">
        <h2 class="text-center mb-4">You can find us here</h2>

        <div class="location-grid d-flex flex-wrap justify-content-between align-items-start">
          <!-- Daftar lokasi -->
          <div class="location-list mb-4">
            <div class="location-item d-flex align-items-center mb-3">
              <i class="ri-map-pin-line me-2"></i>
              <span>Indonesia</span>
              <i class="ri-arrow-right-s-line ms-2"></i>
            </div>
          </div>

          <!-- Google Maps -->
          <div class="world-map flex-grow-1" style="min-height: 400px; border-radius: 10px; overflow: hidden">
            <iframe
              width="100%"
              height="400"
              style="border: 0"
              loading="lazy"
              allowfullscreen
              referrerpolicy="no-referrer-when-downgrade"
              src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d126743.26925911294!2d106.68942908765759!3d-6.229386739737676!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e69f42b1f2b4d2b%3A0x301576d14febd20!2sJakarta!5e0!3m2!1sen!2sid!4v1694769382644!5m2!1sen!2sid"
            >
            </iframe>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA Section -->
    <section class="cta py-5">
      <div class="container cta-container d-flex flex-wrap justify-content-between align-items-center">
        <div class="cta-content text-center text-md-start">
          <h2>Looking for support?</h2>
          <p>Are you tired of slow responses and a lack of transparency from your current IT provider?</p>
          <div class="cta-buttons mt-3">
            <!-- Tombol Get Started -->
            <a href="Contact" class="btn btn-primary me-2"> Kontak kami <i class="ri-arrow-right-s-line"></i> </a>
          </div>
        </div>
        <div class="cta-image"></div>
      </div>
    </section>
  </div>
</template>

<script>
import axios from 'axios';
import dstImage from '@/image/dst.png';

const API_BASE = 'http://localhost:8080/api';

export default {
  name: 'Home',
  data() {
    return {
      // Hero background cycling
      heroBackgrounds: [],
      currentBgIndex: 0,
      bgCycleInterval: null,
      heroBgLoading: false,
      heroBgError: null,

      // FAQ data
      faqs: [],
      faqLoading: false,
      faqError: null,
      pageLimit: 10,

      // Client logos
      logos: ['Blossom', 'Umbrella', 'Vision', 'Network', 'Flash'],
      currentIndex: 0,
      intervalId: null,

      // Canvas animation state
      _particleAnimId: null,
      _particles: null,
      _heroCanvasCleanup: null,
      experiences: [],
      expError: null,
      searchQuery: "",      // 🆕 pencarian
      selectedYear: new Date().getFullYear(), // otomatis 2025 (tahun sekarang)
      visibleCount: 6,
      showAll: false,

      // Team data
      teamMembers: [],
      teamLoading: false,
      teamError: '',

      // Services data
      services: [],
      servicesLoading: false,
      servicesError: '',
      serviceCurrentIndex: 0,
      serviceIntervalId: null,

      // Testimonials data
      testimonials: [],
      testimonialsLoading: false,
      testimonialsError: '',

      pelatihan: [],
      pelatihanLoading: false,
      pelatihanError: '',
      
      // Pelatihan slider state
      pelatihanCurrentIndex: 0,
      pelatihanIntervalId: null,
    };
  },
  computed: {
    // === Filter berdasarkan tahun yang dipilih ===
   filteredExperiences() {
  return this.experiences.filter((exp) => {
    const expYear = Number(exp.year);
    const selYear = Number(this.selectedYear);
    return expYear === selYear;
  });
},
    // 🆕 Data yang ditampilkan (6 pertama saja)
    visibleExperiences() {
      return this.filteredExperiences.slice(0, this.visibleCount);
    },
    // ✅ Menentukan berapa card yang tampil
     visibleExperiences() {
    return this.showAll ? this.filteredExperiences : this.filteredExperiences.slice(0, 6);
  },
    // 🆕 Daftar tahun 2015–tahun sekarang
    availableYears() {
      const start = 2015;
      const end = new Date().getFullYear();
      const years = [];
      for (let y = start; y <= end; y++) years.push(y);
      return years.reverse();
    },
    // duplicate logos so the track can loop seamlessly
    duplicatedLogos() {
      return [...this.logos, ...this.logos];
    },
    // duplicate pelatihan data so the track can loop seamlessly
    duplicatedPelatihan() {
      if (this.pelatihan.length === 0) return [];
      return [...this.pelatihan, ...this.pelatihan];
    },

    duplicatedServices() {
    if (this.services.length === 0) return []
    return [...this.services, ...this.services]
    },

  },
  mounted() {
    // start auto sliding every 3 seconds
    this.startAutoSlide();
    window.addEventListener('resize', this.resetPosition);

    // initialize hero canvas animation
    // this.initHeroCanvas();

    // fetch faqs from API
    this.fetchFaqs();

    // fetch team members for Our Team section
    this.fetchTeamMembers();

    // fetch services from API
    this.fetchServices().then(() => {
    // mulai slider setelah data ter-load
    this.$nextTick(() => {
      setTimeout(() => {
        this.startServiceAutoSlide()
      }, 500)
    })
  })

    // fetch testimonials from API
    this.fetchTestimonials();

    // 🔹 fetch experiences dari API
    this.fetchExperiences();

    this.fetchPelatihan();

    // fetch hero backgrounds and start cycling
    this.fetchHeroBackgrounds();
    
    // start pelatihan auto sliding
    this.startPelatihanAutoSlide();
  },
  beforeUnmount() {
    this.stopAutoSlide();
    window.removeEventListener('resize', this.resetPosition);

    // stop background cycling
    this.stopBackgroundCycling();
    
    // stop pelatihan auto sliding
    this.stopPelatihanAutoSlide();

    // stop hero canvas animation
    // this.stopHeroCanvas();

    this.stopServiceAutoSlide()

  },
  methods: {
    getApiBase() {
      return import.meta?.env?.VITE_API_BASE_URL || 'http://localhost:8080';
    },

    getPhotoUrl(photoPath) {
      if (!photoPath) return '';
      const base = this.getApiBase();
      if (photoPath.startsWith('http://') || photoPath.startsWith('https://')) return photoPath;
      const url = `${base}/${photoPath}`.replace(/\\/g, '/');
      console.log('getPhotoUrl:', photoPath, '->', url);
      return url;
    },

    async fetchPelatihan() {
      this.pelatihanLoading = true;
      this.pelatihanError = '';
      try {
        const base = this.getApiBase();
        const res = await axios.get(`${base}/api/pelatihan`);

        // lihat dulu bentuk respons untuk debugging
        console.log('[fetchPelatihan] raw response:', res && res.data ? res.data : res);

        // Ambil kandidat data dari beberapa kemungkinan struktur respons
        let raw = undefined;
        if (res?.data?.pelatihan !== undefined) {
          raw = res.data.pelatihan; // { pelatihan: [...] }  (yang kita harapkan)
        } else if (res?.data?.data !== undefined) {
          raw = res.data.data; // kemungkinan struktur pagination { data: [...] }
        } else if (res?.data !== undefined) {
          raw = res.data; // langsung data atau objek tunggal
        } else {
          raw = []; // fallback
        }

        // Jika backend mengirimkan objek single item, bungkus jadi array
        let data = [];
        if (Array.isArray(raw)) {
          data = raw;
        } else if (raw && typeof raw === 'object') {
          // Jika raw punya property yang merupakan array (mis. {items: [...]})
          const arrKey = Object.keys(raw).find((k) => Array.isArray(raw[k]));
          if (arrKey) {
            data = raw[arrKey];
          } else {
            // anggap raw adalah single object -> jadikan array 1 elemen
            data = [raw];
          }
        } else {
          // bukan object nor array (mis. string/error) -> kosongkan
          data = [];
        }

        // Jika masih bukan array, aman-kan lagi
        if (!Array.isArray(data)) data = [];

         // mapping ke struktur yang dipakai komponen dan batasi hanya 3 data
         this.pelatihan = data.slice(0, 3).map((p, i) => ({
           id: p.id ?? i,
           image: p.image ?? '',
           category: p.category ?? '',
           jenis: p.jenis ?? '',
           description: p.deskripsi ?? p.description ?? '',
         }));

        // bila tidak ada data, kamu bisa tunjukkan pesan khusus
        if (this.pelatihan.length === 0) {
          // jika backend mengembalikan error dalam body, tampilkan
          if (res?.data?.error) {
            this.pelatihanError = res.data.error;
          } else {
            this.pelatihanError = 'Tidak ada data pelatihan.';
          }
        } else {
          // Start pelatihan slider after data is loaded
          this.$nextTick(() => {
            this.startPelatihanAutoSlide();
          });
        }
      } catch (e) {
        // jika ada respons dari server bisa juga tampilkan untuk debugging
        console.error('fetchPelatihan error:', e);
        if (e?.response) {
          console.error('response data:', e.response.data);
          console.error('response status:', e.response.status);
        }
        this.pelatihanError = 'Gagal memuat data pelatihan.';
        this.pelatihan = [];
      } finally {
        this.pelatihanLoading = false;
      }
    },

    async fetchTeamMembers() {
      this.teamLoading = true;
      this.teamError = '';
      try {
        const base = this.getApiBase();
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
        this.teamError = 'Gagal memuat data tim.';
        console.error('fetchTeamMembers error:', e);
      } finally {
        this.teamLoading = false;
      }
    },

   async fetchServices() {
    this.servicesLoading = true
    this.servicesError = ''
    try {
      const base = this.getApiBase()
      const res = await axios.get(`${base}/api/services?page=1&limit=6`)
      const servicesData = res?.data?.services || []
      this.services = servicesData.map((s) => ({
        id: s.id,
        title: s.title || '',
        description: s.description || '',
        icon: s.icon || '',
        images: s.images || '',
      }))
    } catch (e) {
      this.servicesError = 'Gagal memuat data layanan.'
      console.error('fetchServices error:', e)
    } finally {
      this.servicesLoading = false
    }
  },

  startServiceAutoSlide() {
    this.stopServiceAutoSlide()
    if (this.services.length === 0) return
    this.serviceIntervalId = setInterval(() => {
      this.serviceCurrentIndex += 1
      this.updateServiceTrack()
    }, 4000)
  },

  stopServiceAutoSlide() {
    if (this.serviceIntervalId) {
      clearInterval(this.serviceIntervalId)
      this.serviceIntervalId = null
    }
  },

  updateServiceTrack() {
    const track = this.$refs.serviceTrack
    if (!track) return

    const cards = track.querySelectorAll('.service-card')
    if (cards.length === 0) return

    const itemWidth = cards[0].offsetWidth + 30 // jarak antar item
    const totalWidth = itemWidth * this.services.length

    if (this.serviceCurrentIndex >= this.services.length) {
      // reset
      this.serviceCurrentIndex = 0
      track.style.transition = 'none'
      track.style.transform = 'translateX(0px)'
      void track.offsetWidth // trigger reflow
      setTimeout(() => {
        track.style.transition = 'transform 0.8s ease'
      }, 50)
      return
    }

    const translateX = -(this.serviceCurrentIndex * itemWidth)
    track.style.transition = 'transform 0.8s ease'
    track.style.transform = `translateX(${translateX}px)`
  },

    async fetchFaqs() {
      this.faqLoading = true;
      this.faqError = null;
      try {
        const token = localStorage.getItem('token');
        const headers = token ? { Authorization: `Bearer ${token}` } : {};

        const res = await axios.get(`${API_BASE}/faqs`, {
          headers,
          params: { limit: this.pageLimit },
        });

        const arr = Array.isArray(res.data) ? res.data : Array.isArray(res.data?.faqs) ? res.data.faqs : [];
        this.faqs = arr.map((f) => ({
          id: f.id ?? f.ID ?? null,
          question: f.question ?? f.question_text ?? '',
          answer: f.answer ?? f.body ?? '',
          open: false,
        }));
      } catch (e) {
        this.faqError = e?.response?.data?.message || e.message || 'Gagal mengambil FAQ';
        console.error('fetchFaqs error:', e);
      } finally {
        this.faqLoading = false;
      }
    },

    toggleFaq(idx) {
      this.faqs = this.faqs.map((f, i) => ({ ...f, open: i === idx ? !f.open : false }));
    },

    startAutoSlide() {
      this.stopAutoSlide();
      const step = 1;
      this.intervalId = setInterval(() => {
        this.currentIndex += step;
        this.updateTrack();
      }, 3000);
    },

    stopAutoSlide() {
      if (this.intervalId) {
        clearInterval(this.intervalId);
        this.intervalId = null;
      }
    },

    resetPosition() {
      this.currentIndex = 0;
      this.updateTrack(true);
    },

    updateTrack(noTransition = false) {
      this.$nextTick(() => {
        const track = this.$refs.track;
        if (!track) return;
        const item = track.querySelector('.client-logo');
        if (!item) return;
        const itemWidth = item.offsetWidth + parseInt(getComputedStyle(item).marginRight || '0', 10);
        const translateX = -(this.currentIndex * itemWidth);
        if (noTransition) {
          track.style.transition = 'none';
        } else {
          track.style.transition = 'transform 0.6s ease';
        }
        track.style.transform = `translateX(${translateX}px)`;

        if (this.currentIndex >= this.logos.length) {
          setTimeout(() => {
            track.style.transition = 'none';
            track.style.transform = `translateX(0px)`;
            this.currentIndex = 0;
          }, 610);
        }
      });
    },

    initHeroCanvas() {
      if (typeof window === 'undefined' || !this.$refs.heroCanvas) return;
      const raw = this.$refs.heroCanvas;
      const canvas = Array.isArray(raw) ? raw[0] : raw;
      if (!canvas || !canvas.getContext) return;
      const ctx = canvas.getContext('2d');
      if (!ctx) return;
      const DPR = window.devicePixelRatio || 1;

      const resize = () => {
        const rect = canvas.getBoundingClientRect();
        canvas.width = Math.max(1, Math.floor(rect.width * DPR));
        canvas.height = Math.max(1, Math.floor(rect.height * DPR));
        if (ctx && ctx.setTransform) ctx.setTransform(DPR, 0, 0, DPR, 0, 0);
      };

      const createParticles = (count) => {
        const particles = [];
        for (let i = 0; i < count; i++) {
          particles.push({
            x: Math.random() * (canvas.width / DPR),
            y: Math.random() * (canvas.height / DPR),
            vx: (Math.random() - 0.5) * 0.5,
            vy: (Math.random() - 0.5) * 0.5,
            r: 1 + Math.random() * 3,
            alpha: 0.25 + Math.random() * 0.6,
            hue: 200 + Math.random() * 80,
          });
        }
        return particles;
      };

      const onResize = () => {
        resize();
        const count = Math.max(30, Math.floor((canvas.width / DPR) * 0.06));
        this._particles = createParticles(count);
      };

      window.addEventListener('resize', onResize);
      onResize();

      const draw = () => {
        if (!ctx) return;
        const w = canvas.width / DPR;
        const h = canvas.height / DPR;
        ctx.clearRect(0, 0, w, h);

        const grad = ctx.createLinearGradient(0, 0, 0, h);
        grad.addColorStop(0, 'rgba(12,18,36,0.12)');
        grad.addColorStop(1, 'rgba(0,0,0,0.28)');
        ctx.fillStyle = grad;
        ctx.fillRect(0, 0, w, h);

        const particles = this._particles || [];
        const t = Date.now();

        for (let i = 0; i < particles.length; i++) {
          const p = particles[i];
          p.x += p.vx + Math.sin((t + i * 100) * 0.0008) * 0.3;
          p.y += p.vy + Math.cos((t + i * 120) * 0.0006) * 0.3;

          if (p.x < -10) p.x = w + 10;
          if (p.x > w + 10) p.x = -10;
          if (p.y < -10) p.y = h + 10;
          if (p.y > h + 10) p.y = -10;

          const rg = ctx.createRadialGradient(p.x, p.y, 0, p.x, p.y, p.r * 6);
          rg.addColorStop(0, `hsla(${p.hue},85%,70%,${p.alpha})`);
          rg.addColorStop(0.3, `hsla(${p.hue},75%,60%,${p.alpha * 0.6})`);
          rg.addColorStop(1, `hsla(${p.hue},60%,40%,0)`);
          ctx.fillStyle = rg;
          ctx.beginPath();
          ctx.arc(p.x, p.y, p.r * 6, 0, Math.PI * 2);
          ctx.fill();

          ctx.beginPath();
          ctx.fillStyle = `rgba(255,255,255,${Math.min(1, p.alpha)})`;
          ctx.arc(p.x, p.y, p.r, 0, Math.PI * 2);
          ctx.fill();
        }

        for (let a = 0; a < particles.length; a++) {
          for (let b = a + 1; b < particles.length; b++) {
            const pa = particles[a];
            const pb = particles[b];
            const dx = pa.x - pb.x;
            const dy = pa.y - pb.y;
            const dist = Math.sqrt(dx * dx + dy * dy);
            if (dist < 90) {
              ctx.strokeStyle = `rgba(255,255,255,${0.04 * (1 - dist / 90)})`;
              ctx.lineWidth = 1;
              ctx.beginPath();
              ctx.moveTo(pa.x, pa.y);
              ctx.lineTo(pb.x, pb.y);
              ctx.stroke();
            }
          }
        }

        this._particleAnimId = requestAnimationFrame(draw);
      };

      draw();

      this._heroCanvasCleanup = () => {
        window.removeEventListener('resize', onResize);
        if (this._particleAnimId) cancelAnimationFrame(this._particleAnimId);
      };
    },

    stopHeroCanvas() {
      if (this._heroCanvasCleanup) {
        try {
          this._heroCanvasCleanup();
        } catch (e) {
          console.error('heroCanvasCleanup error:', e);
        }
        this._heroCanvasCleanup = null;
      }
      this._particles = null;
      if (this._particleAnimId) {
        try {
          cancelAnimationFrame(this._particleAnimId);
        } catch (e) {
          console.error('cancelAnimationFrame error:', e);
        }
        this._particleAnimId = null;
      }
    },

    goToService(serviceId) {
      // Navigate to service detail page
      this.$router.push(`/services/${serviceId}`);
    },

    async fetchExperiences() {
      try {
        const res = await axios.get(`${API_BASE}/experiences`);
        this.experiences = Array.isArray(res.data) ? res.data : [];
      } catch (err) {
        this.expError = 'Gagal memuat data experiences';
        console.error(err);
      }
    },
    // 🆕 Tombol untuk menambah jumlah card
    showMore() {
      this.visibleCount += 6;
    },
     selectYear(year) {
      this.selectedYear = year;
      this.showAll = false; // reset ke tampilan awal setiap ganti tahun
    },
    goToIndustry(id) {
      this.$router.push(`/industry/${id}`);
    },

    formatYear(dateStr) {
  if (!dateStr) return 'N/A';
  const date = new Date(dateStr);
  if (isNaN(date.getTime())) return 'N/A';
  return date.getFullYear();
},


    goToPelatihan(id) {
      if (id === undefined || id === null) return;
      this.$router.push(`/pelatihan/${id}`);
    },

    async fetchTestimonials() {
      this.testimonialsLoading = true;
      this.testimonialsError = '';
      try {
        const base = this.getApiBase();
        const res = await axios.get(`${base}/api/testimonis`);
        console.log('Testimonials API response:', res.data);
        const testimonialsData = res?.data?.testimonials || [];
        console.log('Testimonials data:', testimonialsData);
        this.testimonials = testimonialsData.map((t) => ({
          id: t.id,
          nama: t.nama || '',
          url_foto_profil: t.url_foto_profil || '',
          perusahaan: t.perusahaan || '',
          pesan: t.pesan || '',
        }));
        console.log('Processed testimonials:', this.testimonials);
      } catch (e) {
        this.testimonialsError = 'Gagal memuat data testimoni.';
        console.error('Error fetching testimonials:', e);
      } finally {
        this.testimonialsLoading = false;
      }
    },

    handleImageError(event) {
      // safer handling with null checks
      const img = event && (event.target || event.currentTarget);
      if (!img) return;

      try {
        img.style.display = 'none';
      } catch (e) {
        // ignore style errors
      }

      const parent = img.parentElement;
      if (!parent) return;

      // prefer an element with class ".fallback-icon"
      const icon = parent.querySelector('i.fallback-icon') || parent.querySelector('i');
      if (icon) {
        try {
          icon.style.display = 'flex';
          icon.setAttribute('role', 'img');
          icon.setAttribute('aria-label', 'Gambar tidak tersedia');
        } catch (e) {
          // ignore
        }
      }
    },

    handlePelatihanImageError(event) {
      const img = event && (event.target || event.currentTarget);
      if (!img) return;
      try {
        img.style.display = 'none';
      } catch (e) {
        // ignore
      }
      const wrapper = img && img.parentElement;
      if (!wrapper) return;
      const fallback = wrapper.querySelector('.pelatihan-fallback');
      if (fallback) {
        try {
          fallback.style.display = 'block';
        } catch (e) {
          // ignore
        }
      }
    },

    truncateText(text, maxWords = 50) {
      if (!text) return '';
      const words = String(text).split(/\s+/).filter(Boolean);
      if (words.length <= maxWords) return text;
      return words.slice(0, maxWords).join(' ') + '...';
    },

    // Hero Background Methods
    async fetchHeroBackgrounds() {
      this.heroBgLoading = true;
      this.heroBgError = null;
      try {
        const response = await axios.get(`${API_BASE}/hero_backgrounds`);
        const backgrounds = response.data.hero_backgrounds || response.data || [];

        // Get only the 3 most recent backgrounds
        this.heroBackgrounds = backgrounds.sort((a, b) => new Date(b.created_at) - new Date(a.created_at)).slice(0, 3);

        // Start cycling if we have backgrounds
        if (this.heroBackgrounds.length > 0) {
          this.startBackgroundCycling();
        }
      } catch (error) {
        console.error('Error fetching hero backgrounds:', error);
        this.heroBgError = error.message;
      } finally {
        this.heroBgLoading = false;
      }
    },

    getImageUrl(imagePath) {
      if (!imagePath) return '';
      // Normalize Windows-style paths to URL paths and ensure leading slash
      let normalized = String(imagePath).replace(/\\/g, '/');
      if (!normalized.startsWith('/')) normalized = `/${normalized}`;
      return `http://localhost:8080${normalized}`;
    },

    startBackgroundCycling() {
      if (this.heroBackgrounds.length <= 1) return;

      // Clear any existing interval
      this.stopBackgroundCycling();

      // Cycle every 5 seconds
      this.bgCycleInterval = setInterval(() => {
        this.currentBgIndex = (this.currentBgIndex + 1) % this.heroBackgrounds.length;
      }, 5000);
    },

    stopBackgroundCycling() {
      if (this.bgCycleInterval) {
        clearInterval(this.bgCycleInterval);
        this.bgCycleInterval = null;
      }
    },

    // Pelatihan Slider Methods
    startPelatihanAutoSlide() {
      this.stopPelatihanAutoSlide();
      if (this.pelatihan.length === 0) return;
      
      this.pelatihanIntervalId = setInterval(() => {
        this.pelatihanCurrentIndex += 1;
        this.updatePelatihanTrack();
      }, 4000); // Slide every 4 seconds
    },

    stopPelatihanAutoSlide() {
      if (this.pelatihanIntervalId) {
        clearInterval(this.pelatihanIntervalId);
        this.pelatihanIntervalId = null;
      }
    },

    updatePelatihanTrack() {
      this.$nextTick(() => {
        const track = this.$refs.pelatihanTrack;
        if (!track) return;
        
        const item = track.querySelector('.pelatihan-card');
        if (!item) return;
        
        const itemWidth = item.offsetWidth + parseInt(getComputedStyle(item).marginRight || '0', 10);
        const translateX = -(this.pelatihanCurrentIndex * itemWidth);
        
        track.style.transition = 'transform 0.8s ease';
        track.style.transform = `translateX(${translateX}px)`;

        // Reset position when we reach the end of the first set
        if (this.pelatihanCurrentIndex >= this.pelatihan.length) {
          setTimeout(() => {
            track.style.transition = 'none';
            track.style.transform = `translateX(0px)`;
            this.pelatihanCurrentIndex = 0;
          }, 810);
        }
      });
    },
  },
};
</script>

<style>
.home {
  color: #333;
  background-color: #000;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

/* Hero Section */
.hero {
  min-height: 600px;
  display: flex;
  align-items: center;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  color: #fff;
  position: relative;
}

/* Hero background cycling styles */
.hero-bg-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  opacity: 0;
  transition: opacity 2s ease-in-out;
  z-index: 1;
}

.hero-bg-image.active {
  opacity: 1;
}

.hero-container {
  position: relative;
  z-index: 2;
}

.hero-content {
  max-width: 900px;
  text-align: left;
}

.hero-canvas {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  z-index: 1;
  pointer-events: none;
}

.hero .hero-content {
  position: relative;
  z-index: 2;
}

.hero > .container {
  margin: 0;
  padding-left: 20px;
  padding-right: 20px;
  max-width: 1200px;
}

.hero-subtitle {
  color: #64b5f6;
  font-size: 12px;
  letter-spacing: 2px;
  margin-bottom: 20px;
}

.hero-title {
  font-size: 72px;
  font-weight: 700;
  line-height: 1.02;
  margin: 10px 0 30px;
  color: #fff;
  letter-spacing: -1px;
  text-wrap: balance;
}

.hero-buttons {
  display: flex;
  gap: 15px;
}

.btn {
  padding: 14px 28px;
  border-radius: 4px;
  border: none;
  cursor: pointer;
  font-size: 14px;
  font-weight: 600;
  transition: all 0.3s;
}

.btn-primary {
  background: linear-gradient(90deg, #1e3a8a 0%, #3b82f6 60%);
  color: #fff;
  border-radius: 12px;
  padding: 14px 26px;
  box-shadow: 0 6px 20px rgba(59, 91, 219, 0.25);
}

.btn-primary:hover {
  transform: translateY(-3px);
}

.btn-secondary {
  background: rgba(0, 0, 0, 0.5);
  color: #fff;
  border: 2px solid rgba(255, 255, 255, 0.4);
  border-radius: 12px;
  padding: 12px 24px;
}

.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.06);
}

.btn-white {
  background: #fff;
  color: #000;
}

.btn-lime {
  margin-top: 7px;
  background: #84cc16;
  color: #fff;
  text-decoration: none;
}

.btn-lime:hover {
  background: #65a30d;
  color: #fff;
  text-decoration: none;
}

/* Clients */
.clients {
  background: #000;
  padding: 40px 0;
}

.client-logos-wrap {
  overflow: hidden;
  width: 100%;
}

.client-logos-track {
  display: flex;
  gap: 30px;
  align-items: center;
  will-change: transform;
}

.client-logo {
  flex: 0 0 auto;
  min-width: 160px;
  padding: 18px 24px;
  background: rgba(255, 255, 255, 0.03);
  color: #666;
  border-radius: 8px;
  text-align: center;
  font-size: 20px;
  opacity: 0.5;
}

/* What We Do */
.what-we-do {
  padding: 100px 0;
  background: #000;
}

.section-header {
  text-align: center;
  margin-bottom: 60px;
}

.section-badge {
  display: inline-block;
  background: #e6eefb;
  color: #000;
  padding: 8px 20px;
  border-radius: 20px;
  font-size: 12px;
  letter-spacing: 1px;
  margin-bottom: 20px;
}

.section-badge.light {
  background: rgba(255, 255, 255, 0.1);
}

.section-title-service {
  font-size: 40px;
  max-width: 800px;
  margin: 0 auto;
  line-height: 1.4;
  color: #000000;
}

.section-title {
  font-size: 40px;
  max-width: 800px;
  margin: 0 auto;
  line-height: 1.4;
  color: #000;
}

.section-title.light {
  color: #fff;
}

.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 30px;
}

/* SERVICE SLIDER STYLING */
.service-slider-wrap {
  overflow: hidden;
  width: 100%;
  margin-top: 40px;
}

.service-slider-track {
  display: flex;
  gap: 30px;
  align-items: stretch;
  will-change: transform;
}

.service-card {
  background: rgba(255, 255, 255, 0.05);
  padding: 40px;
  border-radius: 8px;
  text-align: center;
  transition: transform 0.3s;
  color: #fff;
  flex: 0 0 auto;
  min-width: 300px;
  max-width: 350px;
  width: calc((100% - 60px) / 3); /* 3 item per slide */
}

.service-card:hover {
  transform: translateY(-10px);
}

@media (max-width: 768px) {
  .service-card {
    width: calc((100% - 20px) / 1);
  }
  .service-slider-track {
    gap: 20px;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  .service-card {
    width: calc((100% - 30px) / 2);
  }
}


.service-card:hover {
  transform: translateY(-10px);
}

.service-icon {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32px;
  margin: 0 auto 20px;
  border: 2px solid;
}

.service-icon i,
.capability-icon i {
  font-size: 32px;
  color: inherit;
}

.learn-more i,
.btn i {
  font-size: 16px;
  vertical-align: middle;
  margin-left: 6px;
}

.service-icon.green {
  border-color: #10b981;
  color: #10b981;
}
.service-icon.blue {
  border-color: #3b82f6;
  color: #3b82f6;
}
.service-icon.purple {
  border-color: #8b5cf6;
  color: #8b5cf6;
}
.service-icon.orange {
  border-color: #f59e0b;
  color: #f59e0b;
}

.service-card h3 {
  font-size: 20px;
  margin-bottom: 15px;
}

/* Capabilities */
.capabilities {
  padding: 100px 0;
  background: #f8f9fa;
}

.capabilities-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: start;
}

.capabilities-title {
  font-size: 42px;
  line-height: 1.3;
  margin-top: 20px;
}

.text-blue {
  color: #3b82f6;
}

.capabilities-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.capability-item {
  display: flex;
  gap: 20px;
  padding: 30px;
  background: #fff;
  border-radius: 8px;
  transition: all 0.3s;
}

.capability-item:hover {
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.15);
  background: #000;
  color: #fff;
  transform: translateY(-6px);
}

.capability-item:hover .learn-more {
  color: #fff;
}

.capability-icon {
  width: 60px;
  height: 60px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 28px;
  flex-shrink: 0;
}

.capability-icon.red {
  background: #fee;
  color: #ef4444;
}
.capability-icon.blue {
  background: #dbeafe;
  color: #3b82f6;
}
.capability-icon.green {
  background: #d1fae5;
  color: #10b981;
}
.capability-icon.light-blue {
  background: #3b82f6;
  color: #fff;
}
.capability-icon.purple {
  background: #ede9fe;
  color: #8b5cf6;
}
.capability-icon.info-blue {
  background: #dbeafe;
  color: #3b82f6;
}

.capability-content h3 {
  font-size: 20px;
  margin-bottom: 10px;
}

.learn-more {
  color: #3b82f6;
  text-decoration: none;
  font-size: 14px;
  font-weight: 600;
  display: inline-block;
  margin-top: 10px;
}

.learn-more.light {
  color: #fff;
}

/* Stats */
.stats {
  padding: 100px 0;
  background: #fff;
}

.stats-title {
  text-align: center;
  font-size: 40px;
  margin-bottom: 30px;
}

.tabs {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin-bottom: 60px;
}

.tab {
  background: transparent;
  border: none;
  color: #666;
  padding: 10px 20px;
  cursor: pointer;
  font-size: 16px;
  position: relative;
}

.tab.active {
  color: #3b82f6;
}

.tab.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: #3b82f6;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 40px;
}

.stat-card {
  text-align: center;
  padding: 40px;
  background: #f8f9fa;
  border-radius: 8px;
}

.stat-card h4 {
  font-size: 16px;
  color: #666;
  margin-bottom: 20px;
}

.stat-number {
  font-size: 72px;
  font-weight: 700;
}

.stat-number .blue {
  color: #3b82f6;
}

.stat-number .plus {
  color: #000;
  font-size: 48px;
}

/* About */
.about {
  padding: 100px 0;
  background: #1e3a8a;
  color: #fff;
}

.about-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: center;
}

.about-visual {
  height: 400px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 8px;
  position: relative;
}

.about-content h2 {
  font-size: 42px;
  margin: 20px 0 30px;
}

/* === Industries Section (styled like Testimonials) === */
.industries {
  padding: 100px 0;
  background: #fff;
  text-align: center;
}

.industries .section-title {
  font-size: 32px;
  font-weight: 700;
  color: #222;
  margin-bottom: 20px;
  position: relative;
  display: inline-block;
  text-transform: capitalize;
  letter-spacing: 0.5px;
}

.section-title::after {
  content: "";
  display: block;
  width: 80px;
  height: 3px;
  margin: 12px auto 0;
  border-radius: 2px;
}

/* === Filter Tahun === */
.year-filter {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 5px;
  margin-bottom: 40px;
}

.year-btn {
  background: none;
  border: none;
  color: #8b5e00; /* cokelat keemasan */
  font-weight: 600;
  font-size: 16px;
  cursor: pointer;
  transition: color 0.3s ease;
  display: flex;
  align-items: center;
}

.year-btn:hover {
  color: #007bff; /* hijau saat hover */
}

.year-btn.active {
  color: #007bff; /* hijau untuk tahun aktif */
  text-decoration: underline;
}

.separator {
  margin: 0 8px;
  color: #228b22;
  font-weight: 500;
}

/* === Grid Card === */
.testimonials-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 30px;
  margin-top: 30px;
}

/* === Card === */
.testimonial-card {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 25px;
  text-align: left;
  transition: all 0.3s ease;
  cursor: pointer;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.1);
}

.testimonial-card:hover {
  transform: translateY(-6px);
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.15);
}

.testimonial-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background: #e2e8f0;
  margin-bottom: 14px;
  overflow: hidden;
}

.testimonial-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.testimonial-card h4 {
  font-size: 16px;
  font-weight: 600;
  color: #222;
  margin-bottom: 5px;
}

.testimonial-role {
  color: #666;
  font-size: 14px;
}

/* === Tombol Lihat Selengkapnya === */
.see-more-wrapper {
  margin-top: 40px;
  text-align: center;
}

.see-more-btn {
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 6px;
  padding: 10px 20px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.3s ease;
}

.see-more-btn:hover {
  background-color: #0056b3;
}

/* Testimonials */
.testimonials {
  padding: 100px 0;
  background: #fff;
  text-align: center;
}

.testimonials-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 30px;
  margin-top: 60px;
}

/* Pelatihan Section */
.pelatihan {
  background-color: #fff;
  padding: 80px 0;
  text-align: center;
}

.pelatihan-slider-wrap {
  overflow: hidden;
  width: 100%;
  margin-top: 40px;
}

.pelatihan-slider-track {
  display: flex;
  gap: 30px;
  align-items: stretch;
  will-change: transform;
}

.pelatihan-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 30px;
  margin-top: 40px;
}

.pelatihan-card {
  background: #f8f9fa;
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  padding: 0 0 20px 0;
  text-align: left;
  flex: 0 0 auto;
  min-width: 300px;
  max-width: 350px;
  width: calc((100% - 60px) / 3); /* 3 items with 30px gap each */
}

.pelatihan-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
}

.pelatihan-image {
  position: relative;
  width: 100%;
  height: 220px;
  overflow: hidden;
}

.pelatihan-img {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.pelatihan-fallback {
  position: absolute;
  inset: 0;
  display: none;
  background: linear-gradient(135deg, rgb(9, 125, 226) 50%, #3da5e0 100%);
}

.pelatihan-fallback.show {
  display: block;
}

.pelatihan-category {
  display: inline-block;
  margin: 16px 16px 8px 16px;
  padding: 6px 12px;
  border-radius: 999px;
  font-size: 12px;
  letter-spacing: 0.6px;
  color: #000;
  background: #e6eefb;
}

.pelatihan-jenis {
  display: block;
  margin: 0 16px 18px 16px;
  color: #000;
  font-weight: 600;
  font-size: 14px;
  line-height: 1.4;
}

.testimonial-card {
  box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.1);
  padding: 40px;
  background: #f8f9fa;
  border-radius: 8px;
  text-align: left;
  transition: all 0.3s ease; /* supaya animasinya halus */
}

.testimonial-card:hover {
  transform: translateY(-8px); /* naik sedikit */
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15); /* bayangan makin tegas */
}

.testimonial-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background: #cbd5e1;
  margin-bottom: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  position: relative;
}

.testimonial-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 50%;
}

.testimonial-avatar i {
  font-size: 28px;
  color: #475569;
  display: flex;
  align-items: center;
  justify-content: center;
}

.testimonial-card h4 {
  font-size: 20px;
  margin-bottom: 5px;
}

.testimonial-role {
  color: #666;
  font-size: 14px;
  margin-bottom: 15px;
}

/* FAQ */
.faq {
  padding: 100px 0;
  background: #f8f9fa;
}

.faq-container {
  display: grid;
  grid-template-columns: 1fr 2fr;
  gap: 60px;
}

.faq-sidebar h2 {
  font-size: 48px;
  margin: 20px 0 15px;
  color: #000;
}

.faq-sidebar p {
  color: #000;
}

.faq-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.faq-item {
  background: #fff;
  padding: 25px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
}

.faq-item:hover {
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

.faq-question {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 16px;
  font-weight: 600;
}

.faq-question-text {
  font-weight: 600;
  color: #000;
}

.faq-answer {
  margin-top: 10px;
  color: #3b3b3b;
  line-height: 1.6;
}

.faq-toggle {
  color: #666;
}

/* Blog */
.blog {
  padding: 100px 0;
  background: #fff;
}

.blog-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 30px;
  margin-top: 60px;
}

.blog-card {
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s;
  padding: 0;
}

.blog-card:hover {
  transform: translateY(-5px);
}

.blog-image {
  height: 220px;
  background: #cbd5e1;
}

.blog-image.ai {
  background: linear-gradient(135deg, #0ea5e9 0%, #0284c7 100%);
}
.blog-image.aws {
  background: linear-gradient(135deg, #1e3a8a 0%, #1e40af 100%);
}
.blog-image.healthcare {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
}

.blog-card > * {
  padding: 0 20px;
}

.blog-category {
  display: inline-block;
  color: #3b82f6;
  font-size: 12px;
  font-weight: 600;
  margin: 20px 0 10px;
  padding: 0 20px;
}

.blog-card h3 {
  font-size: 20px;
  margin-bottom: 10px;
  padding: 0 20px;
}

.blog-meta {
  color: #666;
  font-size: 13px;
  margin-bottom: 20px;
  padding: 0 20px 20px;
}

/* Location */
.location {
  padding: 100px 0;
  background: #f8f9fa;
}

.location h2 {
  text-align: center;
  font-size: 40px;
  margin-bottom: 15px;
}

.location-subtitle {
  text-align: center;
  max-width: 600px;
  margin: 0 auto 60px;
  color: #666;
}

.location-grid {
  display: grid;
  grid-template-columns: 1fr 2fr;
  gap: 60px;
}

.location-item {
  padding: 20px;
  background: #fff;
  border-radius: 8px;
  margin-bottom: 15px;
  cursor: pointer;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.location-item:hover {
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  transform: translateX(5px);
}

.world-map {
  height: 400px;
  background: #e2e8f0;
  border-radius: 8px;
}

/* CTA */
.cta {
  padding: 100px 0;
  background: linear-gradient(135deg, #7dd3fc 0%, #60a5fa 100%);
}

.cta-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: center;
}

.cta-content h2 {
  font-size: 42px;
  margin-bottom: 15px;
}

.cta-buttons {
  display: flex;
  gap: 15px;
  margin-top: 30px;
}
.cta-buttons a {
  text-decoration: none;
}

.cta-image {
  height: 300px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 8px;
}

/* Our Team */
.our-team {
  padding: 100px 0;
  background: linear-gradient(180deg, #1e3a8a 0%, #60a5fa 100%);
  color: #fff;
}

.team-subtitle {
  font-size: 16px;
  opacity: 0.85;
  max-width: 700px;
  margin: 10px auto 0;
  color: #fff;
  text-align: center;
}

.team-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 40px;
  margin-top: 40px;
}

.team-member {
  background: #ffffff;
  border-radius: 12px;
  padding: 30px;
  text-align: center;
  transition: transform 0.3s ease;
}

.team-member:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
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
  color: #000000;
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

.loading,
.error {
  text-align: center;
  margin-top: 20px;
  color: #bbb;
  font-size: 1.1rem;
}

.error {
  color: #f87171;
}

@media (max-width: 768px) {
  .hero-title {
    font-size: 36px;
  }

  .capabilities-grid,
  .about-container,
  .cta-container,
  .location-grid,
  .faq-container {
    grid-template-columns: 1fr;
  }
  
  /* Pelatihan slider responsive */
  .pelatihan-card {
    min-width: 280px;
    width: calc((100% - 20px) / 1); /* 1 item on mobile */
  }
  
  .pelatihan-slider-track {
    gap: 20px;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  /* Tablet: show 2 items */
  .pelatihan-card {
    width: calc((100% - 30px) / 2); /* 2 items on tablet */
  }
}
</style>
