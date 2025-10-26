<template>
  <div class="services-page">
    <div class="container">
      <section class="hero-section">
        <h1>Our Services</h1>
        <p class="lead">Comprehensive IT solutions for your business</p>
      </section>

      <section class="services-overview">
        <div class="services-grid">
          <div class="service-card">
            <div class="service-icon blue">
              <i class="ri-flashlight-line"></i>
            </div>
            <h3>Managed IT Services</h3>
            <p>Complete management of your IT infrastructure with 24/7 monitoring and support.</p>
            <ul>
              <li>Network monitoring</li>
              <li>Server maintenance</li>
              <li>Help desk support</li>
              <li>System updates</li>
            </ul>
          </div>
          
          <div class="service-card">
            <div class="service-icon green">
              <i class="ri-cloud-line"></i>
            </div>
            <h3>Cloud Services</h3>
            <p>Scalable cloud solutions to modernize your business infrastructure.</p>
            <ul>
              <li>Cloud migration</li>
              <li>Infrastructure as a Service</li>
              <li>Cloud security</li>
              <li>Backup & disaster recovery</li>
            </ul>
          </div>
          
          <div class="service-card">
            <div class="service-icon orange">
              <i class="ri-shield-line"></i>
            </div>
            <h3>Cybersecurity Solutions</h3>
            <p>Protect your business from cyber threats with advanced security measures.</p>
            <ul>
              <li>Security audits</li>
              <li>Threat detection</li>
              <li>Employee training</li>
              <li>Compliance management</li>
            </ul>
          </div>
          
          <div class="service-card">
            <div class="service-icon purple">
              <i class="ri-brain-line"></i>
            </div>
            <h3>AI & Machine Learning</h3>
            <p>Leverage artificial intelligence to gain insights and automate processes.</p>
            <ul>
              <li>Data analytics</li>
              <li>Process automation</li>
              <li>Predictive modeling</li>
              <li>AI integration</li>
            </ul>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Services',
  mounted() {
    document.title = 'Our Services - CoreTech';
  }
};
</script>

<style scoped>
.services-page {
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

.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 40px;
}

.service-card {
  background: #f8f9fa;
  padding: 40px 30px;
  border-radius: 8px;
  text-align: center;
  transition: transform 0.3s, box-shadow 0.3s;
}

.service-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.15);
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

.service-icon.blue {
  border-color: #3b82f6;
  color: #3b82f6;
  background: rgba(59, 130, 246, 0.1);
}

.service-icon.green {
  border-color: #10b981;
  color: #10b981;
  background: rgba(16, 185, 129, 0.1);
}

.service-icon.orange {
  border-color: #f59e0b;
  color: #f59e0b;
  background: rgba(245, 158, 11, 0.1);
}

.service-icon.purple {
  border-color: #8b5cf6;
  color: #8b5cf6;
  background: rgba(139, 92, 246, 0.1);
}

.service-card h3 {
  font-size: 24px;
  margin-bottom: 15px;
  color: #1e3a8a;
}

.service-card p {
  color: #666;
  line-height: 1.6;
  margin-bottom: 25px;
}

.service-card ul {
  list-style: none;
  padding: 0;
  text-align: left;
}

.service-card ul li {
  padding: 8px 0;
  color: #666;
  border-bottom: 1px solid #e2e8f0;
}

.service-card ul li:last-child {
  border-bottom: none;
}

.service-card ul li::before {
  content: '✓';
  color: #10b981;
  font-weight: bold;
  margin-right: 10px;
}

@media (max-width: 768px) {
  .hero-section h1 {
    font-size: 32px;
  }
  
  .services-grid {
    grid-template-columns: 1fr;
  }
}
</style>