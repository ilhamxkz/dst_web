<template>
  <div class="case-study-page">
    <div class="hero-section">
      <div class="container">
        <h1>Case Studies</h1>
        <p>Discover how we've helped businesses transform their IT infrastructure</p>
      </div>
    </div>
    
    <div class="content-section">
      <div class="container">
        <div class="case-study-grid">
          <div class="case-study-card">
            <div class="case-study-header">
              <div class="company-logo">
                <i class="ri-building-line"></i>
              </div>
              <div class="case-study-meta">
                <h3>Manufacturing Company X</h3>
                <span class="industry-tag">Manufacturing</span>
              </div>
            </div>
            <div class="case-study-content">
              <h4>Cloud Migration & Modernization</h4>
              <p>Helped a traditional manufacturing company migrate their entire infrastructure to the cloud, resulting in 40% cost reduction and improved scalability.</p>
              <div class="case-study-stats">
                <div class="stat">
                  <span class="stat-value">40%</span>
                  <span class="stat-label">Cost Reduction</span>
                </div>
                <div class="stat">
                  <span class="stat-value">99.9%</span>
                  <span class="stat-label">Uptime</span>
                </div>
              </div>
            </div>
          </div>

          <div class="case-study-card">
            <div class="case-study-header">
              <div class="company-logo">
                <i class="ri-hospital-line"></i>
              </div>
              <div class="case-study-meta">
                <h3>Healthcare Provider Y</h3>
                <span class="industry-tag">Healthcare</span>
              </div>
            </div>
            <div class="case-study-content">
              <h4>Cybersecurity Implementation</h4>
              <p>Implemented comprehensive cybersecurity measures for a healthcare provider, ensuring HIPAA compliance and protecting sensitive patient data.</p>
              <div class="case-study-stats">
                <div class="stat">
                  <span class="stat-value">100%</span>
                  <span class="stat-label">HIPAA Compliant</span>
                </div>
                <div class="stat">
                  <span class="stat-value">0</span>
                  <span class="stat-label">Security Breaches</span>
                </div>
              </div>
            </div>
          </div>

          <div class="case-study-card">
            <div class="case-study-header">
              <div class="company-logo">
                <i class="ri-car-line"></i>
              </div>
              <div class="case-study-meta">
                <h3>Automotive Corp Z</h3>
                <span class="industry-tag">Automotive</span>
              </div>
            </div>
            <div class="case-study-content">
              <h4>Digital Transformation</h4>
              <p>Led complete digital transformation initiative, modernizing legacy systems and implementing IoT solutions for smart manufacturing.</p>
              <div class="case-study-stats">
                <div class="stat">
                  <span class="stat-value">60%</span>
                  <span class="stat-label">Efficiency Gain</span>
                </div>
                <div class="stat">
                  <span class="stat-value">30%</span>
                  <span class="stat-label">Faster Production</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CaseStudy'
}
</script>

<style scoped>
.case-study-page {
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

.case-study-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 30px;
  margin-top: 40px;
}

.case-study-card {
  background: #111;
  border-radius: 12px;
  padding: 30px;
  transition: transform 0.3s ease;
  border: 1px solid #333;
}

.case-study-card:hover {
  transform: translateY(-5px);
  border-color: #3b82f6;
}

.case-study-header {
  display: flex;
  align-items: center;
  gap: 15px;
  margin-bottom: 20px;
}

.company-logo {
  width: 50px;
  height: 50px;
  background: #3b82f6;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.company-logo i {
  font-size: 1.5rem;
  color: #fff;
}

.case-study-meta h3 {
  margin: 0 0 5px 0;
  font-size: 1.3rem;
  color: #fff;
}

.industry-tag {
  background: #333;
  color: #3b82f6;
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
}

.case-study-content h4 {
  font-size: 1.2rem;
  color: #3b82f6;
  margin-bottom: 15px;
}

.case-study-content p {
  color: #ccc;
  line-height: 1.6;
  margin-bottom: 25px;
}

.case-study-stats {
  display: flex;
  gap: 30px;
}

.stat {
  text-align: center;
}

.stat-value {
  display: block;
  font-size: 1.8rem;
  font-weight: 700;
  color: #3b82f6;
}

.stat-label {
  font-size: 0.9rem;
  color: #999;
}

@media (max-width: 768px) {
  .hero-section h1 {
    font-size: 2rem;
  }
  
  .case-study-grid {
    grid-template-columns: 1fr;
  }
  
  .case-study-stats {
    flex-direction: column;
    gap: 15px;
  }
}
</style>