<template>
  <div class="careers-page">
    <div class="hero-section">
      <div class="container">
        <h1>Join Our Team</h1>
        <p>Build your career with CoreTech and be part of the IT transformation revolution</p>
      </div>
    </div>
    
    <div class="content-section">
      <div class="container">
        <div class="careers-intro">
          <h2>Why Work With Us?</h2>
          <p>At CoreTech, we believe in empowering our team members to achieve their full potential while driving innovation in the IT industry.</p>
        </div>

        <div class="benefits-grid">
          <div class="benefit-card">
            <div class="benefit-icon">
              <i class="ri-money-dollar-circle-line"></i>
            </div>
            <h3>Competitive Salary</h3>
            <p>Industry-leading compensation packages with performance bonuses</p>
          </div>
          
          <div class="benefit-card">
            <div class="benefit-icon">
              <i class="ri-heart-pulse-line"></i>
            </div>
            <h3>Health Benefits</h3>
            <p>Comprehensive health insurance and wellness programs</p>
          </div>
          
          <div class="benefit-card">
            <div class="benefit-icon">
              <i class="ri-time-line"></i>
            </div>
            <h3>Flexible Hours</h3>
            <p>Work-life balance with flexible working arrangements</p>
          </div>
          
          <div class="benefit-card">
            <div class="benefit-icon">
              <i class="ri-graduation-cap-line"></i>
            </div>
            <h3>Learning & Development</h3>
            <p>Continuous learning opportunities and professional development</p>
          </div>
        </div>

        <div class="job-openings">
          <h2>Current Openings</h2>
          <div class="job-list">
            <div class="job-card">
              <div class="job-header">
                <h3>Senior Full Stack Developer</h3>
                <span class="job-type">Full-time</span>
              </div>
              <p class="job-description">
                Join our development team to build cutting-edge web applications and contribute to our digital transformation projects.
              </p>
              <div class="job-details">
                <span class="job-location"><i class="ri-map-pin-line"></i> Remote/Hybrid</span>
                <span class="job-experience"><i class="ri-briefcase-line"></i> 5+ years</span>
              </div>
              <button class="apply-btn">Apply Now</button>
            </div>

            <div class="job-card">
              <div class="job-header">
                <h3>Cloud Solutions Architect</h3>
                <span class="job-type">Full-time</span>
              </div>
              <p class="job-description">
                Lead cloud migration projects and design scalable, secure cloud infrastructure for our enterprise clients.
              </p>
              <div class="job-details">
                <span class="job-location"><i class="ri-map-pin-line"></i> On-site</span>
                <span class="job-experience"><i class="ri-briefcase-line"></i> 7+ years</span>
              </div>
              <button class="apply-btn">Apply Now</button>
            </div>

            <div class="job-card">
              <div class="job-header">
                <h3>Cybersecurity Specialist</h3>
                <span class="job-type">Full-time</span>
              </div>
              <p class="job-description">
                Protect our clients' digital assets by implementing and maintaining comprehensive cybersecurity measures.
              </p>
              <div class="job-details">
                <span class="job-location"><i class="ri-map-pin-line"></i> Hybrid</span>
                <span class="job-experience"><i class="ri-briefcase-line"></i> 4+ years</span>
              </div>
              <button class="apply-btn">Apply Now</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Careers'
}
</script>

<style scoped>
.careers-page {
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

.careers-intro {
  text-align: center;
  margin-bottom: 60px;
}

.careers-intro h2 {
  font-size: 2.5rem;
  margin-bottom: 20px;
  color: #fff;
}

.careers-intro p {
  font-size: 1.1rem;
  color: #ccc;
  max-width: 700px;
  margin: 0 auto;
  line-height: 1.6;
}

.benefits-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 30px;
  margin-bottom: 80px;
}

.benefit-card {
  background: #111;
  border-radius: 12px;
  padding: 30px;
  text-align: center;
  transition: transform 0.3s ease;
  border: 1px solid #333;
}

.benefit-card:hover {
  transform: translateY(-5px);
  border-color: #3b82f6;
}

.benefit-icon {
  width: 60px;
  height: 60px;
  background: #3b82f6;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
}

.benefit-icon i {
  font-size: 1.5rem;
  color: #fff;
}

.benefit-card h3 {
  font-size: 1.3rem;
  margin-bottom: 15px;
  color: #fff;
}

.benefit-card p {
  color: #ccc;
  line-height: 1.6;
}

.job-openings {
  text-align: center;
}

.job-openings h2 {
  font-size: 2.5rem;
  margin-bottom: 40px;
  color: #fff;
}

.job-list {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 30px;
  text-align: left;
}

.job-card {
  background: #111;
  border-radius: 12px;
  padding: 30px;
  border: 1px solid #333;
  transition: transform 0.3s ease;
}

.job-card:hover {
  transform: translateY(-5px);
  border-color: #3b82f6;
}

.job-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 15px;
}

.job-header h3 {
  font-size: 1.4rem;
  color: #fff;
  margin: 0;
}

.job-type {
  background: #3b82f6;
  color: #fff;
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
}

.job-description {
  color: #ccc;
  line-height: 1.6;
  margin-bottom: 20px;
}

.job-details {
  display: flex;
  gap: 20px;
  margin-bottom: 25px;
  font-size: 0.9rem;
  color: #999;
}

.job-details span {
  display: flex;
  align-items: center;
  gap: 5px;
}

.apply-btn {
  background: linear-gradient(45deg, #0f3a8c 0%, #3b82f6 100%);
  color: #fff;
  border: none;
  padding: 12px 24px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 600;
  transition: filter 0.3s ease;
  width: 100%;
}

.apply-btn:hover {
  filter: brightness(1.1);
}

@media (max-width: 768px) {
  .hero-section h1 {
    font-size: 2rem;
  }
  
  .careers-intro h2,
  .job-openings h2 {
    font-size: 2rem;
  }
  
  .benefits-grid,
  .job-list {
    grid-template-columns: 1fr;
  }
  
  .job-header {
    flex-direction: column;
    gap: 10px;
  }
  
  .job-details {
    flex-direction: column;
    gap: 10px;
  }
}
</style>