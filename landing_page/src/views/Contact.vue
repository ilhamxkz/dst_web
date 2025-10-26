<template>
  <div class="contact-page">
    <div class="container">
      <section class="hero-section">
        <h1>Kontak Kami</h1>
      </section>

      <section class="contact-content">
        <div class="contact-grid">
          <!-- Contact Info -->
          <div class="contact-info">
            <h2>Let's Start a Conversation</h2>
            <p>
              Ready to transform your business with cutting-edge IT solutions?
              Our team of experts is here to help.
            </p>

            <div class="contact-methods">
              <div class="contact-method">
                <div class="icon blue">
                  <i class="ri-phone-line"></i>
                </div>
                <div class="details">
                  <h4>Phone</h4>
                  <p>(888) 807-5000</p>
                </div>
              </div>

              <div class="contact-method">
                <div class="icon green">
                  <i class="ri-mail-line"></i>
                </div>
                <div class="details">
                  <h4>Email</h4>
                  <p>info@coretech.com</p>
                </div>
              </div>

              <div class="contact-method">
                <div class="icon orange">
                  <i class="ri-map-pin-line"></i>
                </div>
                <div class="details">
                  <h4>Office</h4>
                  <p>123 Technology Street<br />Silicon Valley, CA 94000</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Contact Form -->
          <div class="contact-form">
            <form @submit.prevent="submitForm">
              <div class="form-group">
                <label for="name">Name *</label>
                <input
                  type="text"
                  id="name"
                  v-model="form.name"
                  required
                  class="form-control"
                />
              </div>

              <div class="form-group">
                <label for="email">Email *</label>
                <input
                  type="email"
                  id="email"
                  v-model="form.email"
                  required
                  class="form-control"
                />
              </div>

              <div class="form-group">
                <label for="phone">Phone *</label>
                <input
                  type="text"
                  id="phone"
                  v-model="form.phone"
                  required
                  class="form-control"
                />
              </div>

              <div class="form-group">
                <label for="message">Message *</label>
                <textarea
                  id="message"
                  v-model="form.message"
                  required
                  class="form-control"
                  rows="5"
                ></textarea>
              </div>

              <button
                type="submit"
                class="btn btn-primary"
                :disabled="submitting"
              >
                {{ submitting ? "Sending..." : "Send Message" }}
                <i class="ri-send-plane-line"></i>
              </button>
            </form>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script>
import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080/api", // ubah sesuai API backend kamu
});

export default {
  name: "Contact",
  data() {
    return {
      form: {
        name: "",
        email: "",
        phone: "",
        message: "",
      },
      submitting: false,
    };
  },
  mounted() {
    document.title = "Contact Us - CoreTech";
  },
  methods: {
    async submitForm() {
      this.submitting = true;
      try {
        // 🧩 pastikan nama endpoint sesuai dengan backend kamu
        const res = await api.post("/contact_messages", this.form);

        if (res.status === 200 || res.status === 201) {
          alert("✅ Thank you for your message! We will get back to you soon.");
          this.form = { name: "", email: "", phone: "", message: "" };
        } else {
          alert("⚠️ Failed to send message. Please try again later.");
        }
      } catch (err) {
        console.error("API Error:", err);
        if (err.response && err.response.status === 404) {
          alert("❌ API 'contact_messages' not found. Check your backend route!");
        } else {
          alert("❌ Error sending message. Please check your connection or API.");
        }
      } finally {
        this.submitting = false;
      }
    },
  },
};
</script>

<style scoped>
.contact-page {
  color: #333;
  background: #fff;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.hero-section {
  text-align: center;
  padding: 20px 0; /* sebelumnya 80px 0 */
  background: linear-gradient(135deg, #004d97 0%, #0079c5 100%);
  color: white;
  margin-bottom: 60px; /* bisa juga dikurangi */
}


.hero-section h1 {
  font-size: 48px;
  margin-bottom: 20px;
}

.hero-section .lead {
  font-size: 20px;
  opacity: 0.9;
}

.contact-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 80px;
  align-items: start;
}

.contact-form {
  background: #f8f9fa;
  padding: 40px;
  border-radius: 8px;
}

.form-group {
  margin-bottom: 20px;
}

.form-control {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid #e2e8f0;
  border-radius: 8px;
}

.btn-primary {
  background: linear-gradient(90deg, #1e3a8a 0%, #3b82f6 60%);
  color: #fff;
}
</style>
