<template>
  <div class="layout">
    <!-- Navigation (unchanged) -->
    <nav
      class="navbar"
      :style="{
        backgroundColor: `rgba(0,0,0,${navOpacity})`,
        boxShadow: navOpacity > 0.08 ? '0 2px 10px rgba(0,0,0,0.2)' : 'none',
        transform: navHidden ? 'translateY(-120%)' : 'translateY(0)',
      }"
    >
      <div class="container nav-container">
        <div class="logo">
          <img :src="logo" alt="Daya Sinergy" class="logo-img" />
        </div>

        <!-- Desktop Menu -->
        <ul class="nav-menu desktop-menu">
          <li>
            <router-link to="/" :class="{ 'active-link': isActive('/') }" @click="scrollToTop"> Beranda </router-link>
          </li>
          <li>
            <a href="#layanan" :class="{ 'active-link': activeSection === 'layanan' }" @click.prevent="scrollToSection('layanan')"> Layanan </a>
          </li>
          <li>
            <a href="#pelatihan" :class="{ 'active-link': activeSection === 'pelatihan' }" @click.prevent="scrollToSection('pelatihan')"> Pelatihan </a>
          </li>
             <li>
            <a 
              href="#industries" 
              :class="{ 'active-link': activeSection === 'industries' }"
              @click.prevent="scrollToSection('industries')"
            >
              Pengalaman
            </a>
             </li>
          <li class="dropdown" @mouseenter="showDropdown" @mouseleave="hideDropdown">
            <a href="#" class="dropdown-toggle"> Tentang kami <i class="ri-arrow-down-s-line"></i> </a>
            <ul class="dropdown-menu" v-show="showPageDropdown">
              <li>
                <a href="#team" :class="{ 'active-link': activeSection === 'team' }" @click.prevent="scrollToSection('team')"> Team </a>
              </li>
              <li>
                <router-link to="/case-study" :class="{ 'active-link': isActive('/case-study') }"> Studi kasus </router-link>
              </li>
              <li>
                <router-link to="/careers" :class="{ 'active-link': isActive('/careers') }"> Karir </router-link>
              </li>
              <li>
                <a href="#faq" :class="{ 'active-link': activeSection === 'faq' }" @click.prevent="scrollToSection('faq')"> FAQ </a>
              </li>
              <li>
                <a href="#testimonials" :class="{ 'active-link': activeSection === 'testimonials' }" @click.prevent="scrollToSection('testimonials')"> Testimonials </a>
              </li>
            </ul>
          </li>
            
          <li class="contactLast">
            <router-link 
              to="#maps"
              :class="{ 'active-link': isActive('/contact') }"
            >
              Kontak kami
            </router-link>
          </li>
        </ul>

        <!-- Mobile Menu -->
        <button class="mobile-menu-toggle" @click="toggleMobileMenu" :class="{ active: showMobileMenu }">
          <span></span>
          <span></span>
          <span></span>
        </button>
      </div>
    </nav>

    <!-- Main -->
    <main class="main-content">
      <router-view />
    </main>

    <!-- Footer (UPDATED: grouped & clickable like navbar) -->
    <footer class="footer">
      <div class="container">
        <div class="footer-brand">
          <div class="logo">
            <img :src="logo" alt="Daya Sinergy" class="logo-img" />
          </div>
          <p class="footer-tagline">Setting New Standards in IT Transformation</p>
        </div>

        <!-- Render groups from data.footerGroups -->
        <div class="footer-grid">
          <div class="footer-column" v-for="(group, gi) in footerGroups" :key="gi">
            <h4>{{ group.title }}</h4>
            <ul>
              <li v-for="(item, ii) in group.items" :key="ii">
                <!-- Route links -->
                <router-link v-if="item.type === 'route'" :to="item.to" :class="{ 'active-link': isActive(item.to) }" @click.native="onFooterRouteClick(item)">
                  {{ item.label }}
                </router-link>

                <!-- Anchor / scroll links -->
                <a v-else-if="item.type === 'anchor'" href="#" :class="{ 'active-link': activeSection === item.to }" @click.prevent="scrollToSection(item.to)">
                  {{ item.label }}
                </a>

                <!-- External links -->
                <a v-else-if="item.type === 'external'" :href="item.to" target="_blank" rel="noopener noreferrer">
                  {{ item.label }}
                </a>

                <!-- If item.type missing or unknown, render plain text (kept minimal per request) -->
                <span v-else>{{ item.label }}</span>
              </li>
            </ul>
          </div>
        </div>

        <div class="footer-bottom">
          <p>Disclaimer | Privacy Service | Corporation</p>
          <p>Copyright © 2025 Daya Sinergy Teknomandiri</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
import logo from '../image/dst.png';

export default {
  name: 'Layout',
  data() {
    return {
      logo,
      navOpacity: 1,
      navHidden: false,
      prevScroll: 0,
      showPageDropdown: false,
      showMobileMenu: false,
      activeSection: null,

      // New structured footer groups. Only include items that exist (route or anchor).
      footerGroups: [
        {
          title: 'Home',
          items: [
            { label: 'Beranda', type: 'route', to: '/' },
            { label: 'Layanan', type: 'anchor', to: 'layanan' },
            { label: 'Pengalaman', type: 'anchor', to: 'industries' },
            { label: 'Pelatihan', type: 'anchor', to: 'pelatihan' },
          ],
        },
        {
          title: 'Tentang kami',
          items: [
            { label: 'Team', type: 'anchor', to: 'team' },
            { label: 'Studi kasus', type: 'route', to: '/case-study' },
            { label: 'Karir', type: 'route', to: '/careers' },
            { label: 'FAQ', type: 'anchor', to: 'faq' },
            { label: 'Testimonials', type: 'anchor', to: 'testimonials' },
          ],
        },
        {
          title: 'Resources',
          items: [
            { label: 'Blog', type: 'route', to: '/blog' },
            { label: 'Terms of Service', type: 'external', to: '/terms' },
            { label: 'Privacy Policy', type: 'external', to: '/privacy' },
          ],
        },
        {
          title: 'Contact',
          items: [
            { label: 'Kontak kami', type: 'route', to: '/contact' },
            { label: 'Facebook', type: 'external', to: 'https://facebook.com' },
            { label: 'Instagram', type: 'external', to: 'https://instagram.com' },
          ],
        },
      ],
    };
  },
  mounted() {
    window.addEventListener('scroll', this.onScroll);
    window.addEventListener('scroll', this.handleScrollSpy);
    this.handleScrollSpy(); // Initial check
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.onScroll);
    window.removeEventListener('scroll', this.handleScrollSpy);
  },
  methods: {
    onScroll() {
      const y = window.scrollY;
      this.navOpacity = y > 100 ? 0.85 : 1;
      this.navHidden = y > this.prevScroll && y > 50;
      this.prevScroll = y;
    },
    handleScrollSpy() {
      // Only run scroll spy on home page
      if (this.$route.path !== '/') {
        this.activeSection = null;
        return;
      }

      const sections = ['layanan', 'team', 'faq', 'testimonials', 'industries'];
      const scrollPosition = window.scrollY + 150; // Offset for better detection

      // Check if we're at the very top (before any section)
      if (window.scrollY < 100) {
        this.activeSection = null;
        return;
      }

      // Find which section is currently visible
      for (const sectionId of sections) {
        const element = document.getElementById(sectionId);
        if (element) {
          const offsetTop = element.offsetTop;
          const offsetBottom = offsetTop + element.offsetHeight;

          if (scrollPosition >= offsetTop && scrollPosition < offsetBottom) {
            this.activeSection = sectionId;
            return;
          }
        }
      }

      // If no section is active, clear it
      this.activeSection = null;
    },
    scrollToSection(sectionId) {
      // If not on home page, navigate to home first
      if (this.$route.path !== '/') {
        this.$router.push('/').then(() => {
          setTimeout(() => {
            this.smoothScrollTo(sectionId);
          }, 120);
        });
      } else {
        this.smoothScrollTo(sectionId);
      }
      this.showMobileMenu = false;
    },
    smoothScrollTo(sectionId) {
      const element = document.getElementById(sectionId);
      if (element) {
        const offsetTop = element.offsetTop - 80; // Offset for navbar
        window.scrollTo({
          top: offsetTop,
          behavior: 'smooth',
        });
      }
    },
    scrollToTop() {
      window.scrollTo({
        top: 0,
        behavior: 'smooth',
      });
    },
    isActive(path) {
      // For exact match on home page
      if (path === '/') {
        return this.$route.path === '/' && !this.activeSection;
      }
      // For other pages
      return this.$route.path === path;
    },
    toggleMobileMenu() {
      this.showMobileMenu = !this.showMobileMenu;
    },
    showDropdown() {
      this.showPageDropdown = true;
    },
    hideDropdown() {
      this.showPageDropdown = false;
    },

    // called when footer route links clicked to optionally scroll or close mobile menu
    onFooterRouteClick(item) {
      // Close mobile menu if open
      this.showMobileMenu = false;
      // If the user clicked a route that is the home page but with an anchor in the item
      // we keep behavior simple: just navigate. Anchors are handled by type === 'anchor'.
    },
  },
};
</script>

<style scoped>

/* underline untuk semua anchor dan item navigasi/footer */
a,
.nav-menu a,
.dropdown-toggle,
.mobile-nav-menu a,
.footer-column ul li a {
  text-decoration: none;
}

/* khusus tombol Kontak (desktop & mobile) */
.contact-btn,
.mobile-contact-btn {
  text-decoration: none;
}

/* jika style gradient/button menimpa underline, pakai !important */
.contact-btn,
.mobile-contact-btn {
  text-decoration: none !important;
}

/* Active link styling - BIRU untuk yang active */
.nav-menu a.active-link,
.dropdown-menu a.active-link,
.footer-column ul li a.active-link {
  color: #3b82f6 !important;
  
}

/* Default link color - PUTIH untuk yang tidak active */
.nav-menu a,
.dropdown-menu a,
.footer-column ul li a {
  color: #fff;
}

/* Hover effect */
.nav-menu a:hover,
.dropdown-menu a:hover,
.footer-column ul li a:hover {
  color: #3b82f6;
}

/* Footer specific tweaks */
.footer-column ul li a {
  color: #999;
  text-decoration: none;
}

.footer-column ul li a:hover {
  color: #fff;
}
</style>

<style>
/* existing global styles kept unchanged (omitted in this snippet) */
</style>

<style scoped>
/* Active link styling - BIRU untuk yang active */
.nav-menu a.active-link,
.dropdown-menu a.active-link {
  color: #3b82f6 !important;
}

/* Default link color - PUTIH untuk yang tidak active */
.nav-menu a,
.dropdown-menu a {
  color: #fff;
}

/* Hover effect */
.nav-menu a:hover,
.dropdown-menu a:hover {
  color: #3b82f6;
}
</style>

<style>
.layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
  color: #333;
  background-color: #000;
  overflow-x: hidden;

}

.main-content {
  flex: 1;
  /* reserve space for fixed navbar */
  padding-top: 72px;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

/* Navigation */
.navbar {
  background: #000;
  padding: 17px 0;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  width: 100%;
  z-index: 9999;
  transition: background-color 0.25s ease, box-shadow 0.25s ease, transform 0.35s cubic-bezier(0.2, 0.9, 0.2, 1);
  backdrop-filter: blur(6px);
}

.nav-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  display: flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
  position: relative;
  z-index: 10;
}

.logo-img {
  filter: brightness(1.1);
  display: block;
  height: 60px;
  max-height: 60px;
  width: auto;
  object-fit: contain;
  transition: transform 0.3s ease, filter 0.3s ease;
}

.logo:hover .logo-img {
  transform: scale(1.05);
  filter: brightness(1.2);
}

.nav-menu {
  display: flex;
  gap: 30px;
  list-style: none;
}

.nav-menu a {
  color: #fff;
  text-decoration: none;
  font-size: 14px;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.55);
  transition: color 0.3s ease, text-shadow 0.18s ease;
  cursor: pointer;
}

/* Dropdown Styles */
.dropdown {
  position: relative;
}

.dropdown-toggle {
  color: #fff;
  text-decoration: none;
  font-size: 14px;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.55);
  transition: text-shadow 0.18s ease;
  display: flex;
  align-items: center;
  gap: 5px;
  cursor: pointer;
}

.dropdown-toggle:hover {
  color: #3b82f6;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  background: rgba(0, 0, 0, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 8px;
  min-width: 160px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  list-style: none;
  padding: 10px 0;
  margin: 0;
  margin-top: 5px;
  z-index: 10000;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

/* Add pseudo-element to bridge the gap */
.dropdown-menu::before {
  content: '';
  position: absolute;
  top: -5px;
  left: 0;
  right: 0;
  height: 5px;
  background: transparent;
}

.dropdown-menu li {
  margin: 0;
}

.dropdown-menu a {
  display: block;
  padding: 10px 20px;
  color: #fff;
  text-decoration: none;
  font-size: 13px;
  transition: all 0.2s ease;
  text-shadow: none;
  cursor: pointer;
}

.dropdown-menu a:hover {
  background: rgba(59, 130, 246, 0.2);
  color: #3b82f6;
}

.contactLast {
  margin-right: 9rem;
}
/* Mobile Menu Styles */
.mobile-menu-toggle {
  display: none;
  flex-direction: column;
  background: none;
  border: none;
  cursor: pointer;
  padding: 8px;
  width: 30px;
  height: 30px;
  justify-content: space-between;
}

.mobile-menu-toggle span {
  width: 100%;
  height: 3px;
  background: #fff;
  border-radius: 2px;
  transition: all 0.3s ease;
  transform-origin: center;
}

.mobile-menu-toggle.active span:nth-child(1) {
  transform: rotate(45deg) translate(6px, 6px);
}

.mobile-menu-toggle.active span:nth-child(2) {
  opacity: 0;
}

.mobile-menu-toggle.active span:nth-child(3) {
  transform: rotate(-45deg) translate(6px, -6px);
}

.mobile-menu {
  display: none;
  position: fixed;
  top: 100%;
  left: 0;
  right: 0;
  background: rgba(0, 0, 0, 0.98);
  backdrop-filter: blur(10px);
  transform: translateY(-100%);
  transition: transform 0.3s cubic-bezier(0.2, 0.9, 0.2, 1);
  z-index: 9998;
  max-height: calc(100vh - 90px);
  overflow-y: auto;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.mobile-menu.active {
  transform: translateY(0);
}

.mobile-nav-menu {
  list-style: none;
  margin: 0;
  padding: 20px;
}

.mobile-nav-menu > li {
  margin-bottom: 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.mobile-nav-menu > li:last-child {
  border-bottom: none;
}

.mobile-nav-menu a {
  display: block;
  color: #fff;
  text-decoration: none;
  padding: 15px 0;
  font-size: 16px;
  font-weight: 500;
  transition: color 0.3s ease;
}

.mobile-nav-menu a:hover,
.mobile-nav-menu a.active-link {
  color: #3b82f6;
}

.mobile-dropdown-toggle {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.mobile-dropdown-toggle i {
  transition: transform 0.3s ease;
}

.mobile-dropdown-toggle i.rotated {
  transform: rotate(180deg);
}

.mobile-dropdown-menu {
  list-style: none;
  margin: 0;
  padding: 0;
  padding-left: 20px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 8px;
  margin-top: 10px;
}

.mobile-dropdown-menu li {
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.mobile-dropdown-menu li:last-child {
  border-bottom: none;
}

.mobile-dropdown-menu a {
  padding: 12px 0;
  font-size: 14px;
  color: #ccc;
}

.mobile-dropdown-menu a:hover,
.mobile-dropdown-menu a.active-link {
  color: #3b82f6;
}

.mobile-contact-btn {
  background: linear-gradient(45deg, #0f3a8c 0%, #3b82f6 100%);
  color: #fff;
  border: none;
  padding: 15px 20px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 16px;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  margin: 20px;
  transition: filter 0.3s ease;
}

.mobile-contact-btn:hover {
  filter: brightness(1.1);
}

.mobile-contact-btn i {
  font-size: 18px;
}

/* Footer */
.footer {
  background: #000;
  color: #fff;
  padding: 80px 0 30px;
  margin-top: auto;
}

.footer-brand {
  margin-bottom: 60px;
}

.footer-tagline {
  color: #999;
  margin-top: 15px;
}

.footer-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 40px;
  margin-bottom: 60px;
}

.footer-column h4 {
  font-size: 18px;
  margin-bottom: 20px;
}

.footer-column ul {
  list-style: none;
}

.footer-column ul li {
  margin-bottom: 12px;
}

.footer-column ul li a {
  color: #999;
  text-decoration: none;
  font-size: 14px;
  transition: color 0.3s;
}

.footer-column ul li a:hover {
  color: #fff;
}

.footer-column ul li a.active-link {
  color: #3b82f6;
}

.footer-bottom {
  border-top: 1px solid #333;
  padding-top: 30px;
  display: flex;
  justify-content: space-between;
  color: #666;
  font-size: 13px;
}

@media (max-width: 768px) {
  /* Hide desktop menu and show mobile elements */
  .desktop-menu {
    display: none !important;
  }

  .desktop-contact,
  .contact-btn:not(.mobile-contact-btn) {
    display: none !important;
  }

  .mobile-menu-toggle {
    display: flex !important;
    position: absolute;
    right: 20px;
    top: 50%;
    transform: translateY(-50%);
    z-index: 10001;
  }

  .mobile-menu {
    display: block !important;
  }

  .logo-img {
    height: 44px;
    max-height: 44px;
  }

  /* Navbar adjustments for mobile */
  .navbar {
    padding: 12px 0;
  }

  .nav-container {
    position: relative;
    justify-content: flex-start;
  }

  /* Ensure hamburger doesn't conflict with other elements */
  .nav-container > *:not(.logo):not(.mobile-menu-toggle):not(.mobile-menu) {
    display: none;
  }

  .footer-bottom {
    flex-direction: column;
    gap: 10px;
  }
}
</style>

<!-- Global styles -->
<style>
html,
body,
#app {
  height: 100%;
  margin: 0;
  padding: 0;
}
</style>
