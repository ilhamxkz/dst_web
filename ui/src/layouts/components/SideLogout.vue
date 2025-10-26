<script setup>
import VerticalNavLink from '@layouts/components/VerticalNavLink.vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { toast } from 'vue3-toastify'
import 'vue3-toastify/dist/index.css'

const router = useRouter()

// ✅ Helper notif
const notifySuccess = (msg) =>
  toast.success(msg, {
    autoClose: 2000,
    position: 'top-right',
    theme: 'colored',
  })

const notifyError = (msg) =>
  toast.error(msg, {
    autoClose: 2500,
    position: 'top-right',
    theme: 'colored',
  })

const handleLogout = async () => {
  try {
    const token = localStorage.getItem('token') || sessionStorage.getItem('token')

    await axios.post(
      'http://localhost:8080/api/logout',
      {},
      { headers: { Authorization: `Bearer ${token}` } }
    )

    // Hapus token di localStorage/sessionStorage
    localStorage.removeItem('token')
    sessionStorage.removeItem('token')

    notifySuccess('Logout berhasil, sampai jumpa!')
    setTimeout(() => router.push('/login'), 1500)
  } catch (error) {
    console.error('Logout gagal:', error)
    notifyError('Logout gagal, silakan coba lagi.')
  }
}
</script>

<template>
  <VerticalNavLink
    :item="{ title: 'Logout', icon: 'ri-logout-box-line', to: '/' }"
    @click="handleLogout"
  />
</template>
