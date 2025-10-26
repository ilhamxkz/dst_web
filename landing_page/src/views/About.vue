<template>
  <div class="about-page">
    <div class="container">
      <section class="hero-section">
        <h1>About CoreTech</h1>
        <p class="lead">Setting New Standards in IT Transformation</p>
      </section>

      <section class="company-overview">
        <div class="overview-grid">
          <div class="content">
            <h2>Our Story</h2>
            <p>
              CoreTech has been at the forefront of IT transformation for over two decades, 
              helping businesses leverage technology to achieve their goals and drive innovation.
            </p>
            <p>
              We specialize in providing comprehensive IT solutions that enable organizations 
              to modernize their infrastructure, enhance security, and accelerate digital transformation.
            </p>
          </div>
          <div class="image-placeholder"></div>
        </div>
      </section>

      <section class="values">
        <h2>Our Values</h2>
        <div class="values-grid">
          <div class="value-card">
            <div class="icon blue">
              <i class="ri-lightbulb-line"></i>
            </div>
            <h3>Innovation</h3>
            <p>We constantly push boundaries and embrace cutting-edge technologies to deliver exceptional solutions.</p>
          </div>
          <div class="value-card">
            <div class="icon green">
              <i class="ri-team-line"></i>
            </div>
            <h3>Collaboration</h3>
            <p>We work closely with our clients to understand their unique needs and deliver customized solutions.</p>
          </div>
          <div class="value-card">
            <div class="icon orange">
              <i class="ri-award-line"></i>
            </div>
            <h3>Excellence</h3>
            <p>We maintain the highest standards of quality and professionalism in everything we do.</p>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script>
export default {
  name: 'About',
  mounted() {
    document.title = 'About Us - CoreTech';
  }
};
</script>

<style scoped>
.about-page {
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
  padding: 80px 0;
  background: linear-gradient(135deg, #1e3a8a 0%, #3b82f6 100%);
  color: white;
  margin-bottom: 80px;
}

.hero-section h1 {
  font-size: 48px;
  margin-bottom: 20px;
}

.hero-section .lead {
  font-size: 20px;
  opacity: 0.9;
}

.company-overview {
  margin-bottom: 80px;
}

.overview-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: center;
}

.overview-grid h2 {
  font-size: 36px;
  margin-bottom: 30px;
  color: #1e3a8a;
}

.overview-grid p {
  font-size: 16px;
  line-height: 1.6;
  margin-bottom: 20px;
  color: #666;
}

.image-placeholder {
  height: 300px;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-radius: 8px;
}

.values h2 {
  text-align: center;
  font-size: 36px;
  margin-bottom: 60px;
  color: #1e3a8a;
}

.values-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 40px;
}

.value-card {
  text-align: center;
  padding: 40px 30px;
  background: #f8f9fa;
  border-radius: 8px;
  transition: transform 0.3s;
}

.value-card:hover {
  transform: translateY(-10px);
}

.icon {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
  font-size: 32px;
}

.icon.blue {
  background: #dbeafe;
  color: #3b82f6;
}

.icon.green {
  background: #d1fae5;
  color: #10b981;
}

.icon.orange {
  background: #fed7aa;
  color: #f59e0b;
}

.value-card h3 {
  font-size: 24px;
  margin-bottom: 15px;
  color: #1e3a8a;
}

.value-card p {
  color: #666;
  line-height: 1.6;
}

@media (max-width: 768px) {
  .hero-section h1 {
    font-size: 32px;
  }
  
  .overview-grid {
    grid-template-columns: 1fr;
  }
}
</style>