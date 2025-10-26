<script setup>
import { reactive, onMounted } from 'vue'
import axios from 'axios'
import { toast } from 'vue3-toastify'
import 'vue3-toastify/dist/index.css'

const api = axios.create({ baseURL: 'http://localhost:8080/api' })

const contactForm = reactive({
  name: '',
  phone: '',
  email: '',
  message: ''
})

// simpan token captcha
let captchaToken = null

// load script recaptcha
onMounted(() => {
  const script = document.createElement("script")
  script.src = "https://www.google.com/recaptcha/api.js"
  script.async = true
  script.defer = true
  document.head.appendChild(script)

  // assign callback agar bisa dipanggil dari window
  window.onCaptchaVerified = (token) => {
    captchaToken = token
  }
})

const submitContact = async () => {
  try {
    if (!captchaToken) {
      toast.error("⚠️ Harap selesaikan reCAPTCHA terlebih dahulu", { autoClose: 3000 })
      return
    }

    await api.post('/contact_messages', { ...contactForm, captcha: captchaToken })

    toast.success('✅ Pesan berhasil dikirim!', { autoClose: 3000 })
    Object.assign(contactForm, { name: '', phone: '', email: '', message: '' })
    captchaToken = null
    if (window.grecaptcha) {
      window.grecaptcha.reset()
    }
  } catch (err) {
    console.error('submitContact err ->', err)
    toast.error('❌ Gagal kirim pesan: ' + (err.response?.data?.message || err.message), { autoClose: 4000 })
  }
}
</script>


<template>
  <div class="mt-5">
    <h3 class="text-center mb-4 fw-bold">Hubungi Kami</h3>
    <div class="card shadow-lg p-4 rounded-4">
      <div class="row g-4">
        <!-- Form Contact -->
        <div class="col-md-6">
          <h5 class="fw-bold mb-3">Kirim Pesan</h5>
          <form @submit.prevent="submitContact">
            <div class="mb-3">
              <input v-model="contactForm.name" type="text" class="form-control" placeholder="Nama Lengkap" required>
            </div>
            <div class="mb-3">
              <input v-model="contactForm.phone" type="tel" class="form-control" placeholder="No HP" required>
            </div>
            <div class="mb-3">
              <input v-model="contactForm.email" type="email" class="form-control" placeholder="Email" required>
            </div>
            <div class="mb-3">
              <textarea v-model="contactForm.message" class="form-control" rows="4" placeholder="Pesan anda..." required></textarea>
            </div>

            <!-- Google reCAPTCHA -->
            <div class="text-center mb-3">
              <div class="g-recaptcha" 
                   data-sitekey="6LfZyskrAAAAAFq4Lg2wcU82zg3B8TIb69M8ivfp" 
                   data-callback="onCaptchaVerified">
              </div>
            </div>

            <button type="submit" class="btn btn-primary w-100 rounded-pill">
              <i class="bi bi-send-fill me-1"></i> Kirim
            </button>
          </form>
        </div>

        <!-- Maps + Alamat -->
        <div class="col-md-6">
          <h5 class="fw-bold mb-3">Lokasi Kami</h5>
          <div style="width: 100%; height: 260px; border-radius: 12px; overflow: hidden;">
            <iframe
              src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3959.340246635408!2d107.63625317499858!3d-6.903729467527378!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e68e65fd7401fbf%3A0x2d761c686baabca1!2sPT%20Waditra%20Reka%20Cipta!5e0!3m2!1sid!2sid!4v1693647420000!5m2!1sid!2sid"
              width="100%"
              height="100%"
              style="border:0;"
              allowfullscreen
              loading="lazy"
              referrerpolicy="no-referrer-when-downgrade">
            </iframe>
          </div>
          <div class="mt-3">
            <p class="mb-2">
              <i class="bi bi-geo-alt-fill text-danger me-2"></i>
              PT DAYA SINERGI TEKNOMANDIRI<br>
              Satrio Tower Level 16<br>
              Jl. Prof. Dr. Satrio Kav C4<br>
              Jakarta Selatan - 12950
            </p>
            <p class="mb-2">
              <i class="bi bi-telephone-fill text-primary me-2"></i>
              081311530954
            </p>
            <p>
              <i class="bi bi-envelope-fill text-success me-2"></i>
              dayasinergi@.com
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
