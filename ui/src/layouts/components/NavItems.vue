<script setup>
import { ref, onMounted, computed } from 'vue'
import axios from 'axios'

import VerticalNavSectionTitle from '@/@layouts/components/VerticalNavSectionTitle.vue'
import VerticalNavGroup from '@layouts/components/VerticalNavGroup.vue'
import VerticalNavLink from '@layouts/components/VerticalNavLink.vue'
import Editor from '@/pages/editor.vue'

const menus = ref([])
const loading = ref(true)
const error = ref(null)

const iconMap = {
  dashboard: 'ri-dashboard-line',
  users: 'ri-user-line',
  user: 'ri-user-line',
  roles: 'ri-id-card-line',
  role: 'ri-id-card-line',
  akses: 'ri-accessibility-line',
  login: 'ri-login-box-line',
  editor: 'ri-edit-box-fill',
  menu: 'ri-menu-line',
  logout: 'ri-logout-box-line',
  post: 'ri-article-line',
  coments: 'ri-message-3-line',
  coment: 'ri-message-3-line',
  blog_post: 'ri-article-line',
  blog_comments: 'ri-message-3-line',
  icons: 'ri-palette-line',
  cards: 'ri-layout-grid-fill',
  tables: 'ri-table-line',
  form_layouts: 'ri-layout-grid-fill',
  page: 'ri-file-text-line',
  pages: 'ri-file-text-line',
  contact: 'ri-message-line',
  client: 'ri-user-line',
  team: 'ri-team-line',
  service: 'ri-service-line',
  companyprofile: 'ri-building-line',
  project: 'ri-booklet-fill',
  projects: 'ri-booklet-fill',
  projectimage: 'ri-gallery-upload-fill',
  project_images: 'ri-gallery-upload-fill',
  experiences: 'ri-trophy-line', // 👈 tambahkan ini
  experience: 'ri-trophy-line',
}

function mapMenuItem(m) {
  const rawName = (m.nama_menu || '').toString()
  const key = rawName ? rawName.toLowerCase() : ''
  const title = rawName ? rawName.charAt(0).toUpperCase() + rawName.slice(1) : 'Untitled'
  const icon = iconMap[key] || 'ri-menu-line'
  return {
    id: m.id,
    key,
    title,
    to: m.routes || '#',
    icon,
  }
}

/**
 * Daftar kunci yang masuk ke group "akses control".
 * Jika mau menambahkan (mis. 'login') tambahkan di sini.
 */
const aksesKeys = new Set(['role', 'roles', 'user', 'users', 'menu', 'menus', 'akses'])

const aksesMenus = computed(() => menus.value.filter(m => aksesKeys.has(m.key)))
const pagesMenus = computed(() => menus.value.filter(m => !aksesKeys.has(m.key)))

onMounted(async () => {
  loading.value = true
  error.value = null

  try {
    const token = localStorage.getItem('token') || localStorage.getItem('authToken') || null

    const res = await axios.get('http://localhost:8080/api/menus', {
      headers: token ? { Authorization: `Bearer ${token}` } : {},
    })

    const dataMenus = res && res.data && Array.isArray(res.data.menus) ? res.data.menus : []
    // Map dan simpan
    menus.value = dataMenus.map(mapMenuItem)
  } catch (err) {
    console.error('Failed to load menus', err)
    if (err.response && err.response.status === 401) {
      error.value = 'Untuk mengakses silahkan login terlebih dahulu'
      menus.value = []
    } else {
      error.value = err.response?.data?.message || err.message || 'Gagal mengambil menu'
    }
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <VerticalNavSectionTitle :item="{ heading: 'Apps & Pages' }" />

  <div
    v-if="loading"
    class="message loading"
  >
    Memuat menu...
  </div>

  <div
    v-else-if="error"
    class="message error"
  >
    Error: {{ error }}
  </div>

  <template v-else>
    <VerticalNavGroup
      :item="{
        title: 'akses control',
        badgeContent: String(aksesMenus.length || 0),
        badgeClass: 'bg-error',
        icon: 'ri-home-smile-line',
      }"
    >
      <VerticalNavLink
        v-for="menu in aksesMenus"
        :key="menu.id"
        :item="{ title: menu.title, icon: menu.icon, to: menu.to }"
      />
    </VerticalNavGroup>

    <VerticalNavGroup
      :item="{
        title: 'pages',
        badgeContent: String(pagesMenus.length || 0),
        badgeClass: 'bg-error',
        icon: 'ri-file-text-line',
      }"
    >
      <VerticalNavLink
        v-for="menu in pagesMenus"
        :key="menu.id"
        :item="{ title: menu.title, icon: menu.icon, to: menu.to }"
      />
    </VerticalNavGroup>
  </template>
</template>


<style scoped>
a {
  text-decoration: none;
}
a:focus,
a:hover,
a:active {
  text-decoration: none;
}

</style>