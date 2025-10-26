<template>
  <div class="blog-page">
    <div class="container">
      <section class="hero-section">
        <h1>Our Blog</h1>
        <p class="lead">Latest insights and expertise in technology</p>
      </section>

      <section class="blog-content">
        <div class="blog-grid">
          <article class="blog-card">
            <div class="blog-image ai"></div>
            <div class="blog-content">
              <span class="blog-category">IT Services</span>
              <h2>AI-powered project management: Benefits & future</h2>
              <p class="blog-excerpt">
                Discover how artificial intelligence is revolutionizing project management with automated scheduling, 
                predictive analytics, and intelligent resource allocation.
              </p>
              <div class="blog-meta">
                <span class="author">by CoreTech Team</span>
                <span class="date">January 28, 2025</span>
              </div>
              <a href="#" class="read-more">Read More <i class="ri-arrow-right-s-line"></i></a>
            </div>
          </article>
          
          <article class="blog-card">
            <div class="blog-image aws"></div>
            <div class="blog-content">
              <span class="blog-category">Cloud Services</span>
              <h2>Key takeaways from AWS re:Invent</h2>
              <p class="blog-excerpt">
                A comprehensive overview of the latest AWS announcements and how they impact enterprise cloud strategies 
                for the upcoming year.
              </p>
              <div class="blog-meta">
                <span class="author">by Cloud Team</span>
                <span class="date">January 25, 2025</span>
              </div>
              <a href="#" class="read-more">Read More <i class="ri-arrow-right-s-line"></i></a>
            </div>
          </article>
          
          <article class="blog-card">
            <div class="blog-image healthcare"></div>
            <div class="blog-content">
              <span class="blog-category">IT Consulting</span>
              <h2>Exploring the potential of AI in healthcare</h2>
              <p class="blog-excerpt">
                How artificial intelligence is transforming patient care, diagnostics, and operational efficiency 
                in modern healthcare systems.
              </p>
              <div class="blog-meta">
                <span class="author">by AI Specialists</span>
                <span class="date">January 22, 2025</span>
              </div>
              <a href="#" class="read-more">Read More <i class="ri-arrow-right-s-line"></i></a>
            </div>
          </article>
          
          <article class="blog-card">
            <div class="blog-image cybersecurity"></div>
            <div class="blog-content">
              <span class="blog-category">Cybersecurity</span>
              <h2>Zero Trust Security: Implementation Guide</h2>
              <p class="blog-excerpt">
                A practical guide to implementing Zero Trust security architecture in your organization 
                to protect against modern cyber threats.
              </p>
              <div class="blog-meta">
                <span class="author">by Security Team</span>
                <span class="date">January 20, 2025</span>
              </div>
              <a href="#" class="read-more">Read More <i class="ri-arrow-right-s-line"></i></a>
            </div>
          </article>
          
          <article class="blog-card">
            <div class="blog-image automation"></div>
            <div class="blog-content">
              <span class="blog-category">Automation</span>
              <h2>The Future of Business Process Automation</h2>
              <p class="blog-excerpt">
                Discover emerging trends in business process automation and how intelligent automation 
                is reshaping workflows across industries.
              </p>
              <div class="blog-meta">
                <span class="author">by Automation Team</span>
                <span class="date">January 18, 2025</span>
              </div>
              <a href="#" class="read-more">Read More <i class="ri-arrow-right-s-line"></i></a>
            </div>
          </article>
          
          <article class="blog-card">
            <div class="blog-image digital"></div>
            <div class="blog-content">
              <span class="blog-category">Digital Transformation</span>
              <h2>Digital Transformation Roadmap for SMBs</h2>
              <p class="blog-excerpt">
                A step-by-step guide for small and medium businesses to successfully navigate 
                their digital transformation journey.
              </p>
              <div class="blog-meta">
                <span class="author">by Strategy Team</span>
                <span class="date">January 15, 2025</span>
              </div>
              <a href="#" class="read-more">Read More <i class="ri-arrow-right-s-line"></i></a>
            </div>
          </article>
        </div>
      </section>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Blog',
  mounted() {
    document.title = 'Blog - CoreTech';
  }
};
</script>

<style scoped>
.blog-page {
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

.blog-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 40px;
  margin-bottom: 80px;
}

.blog-card {
  background: #fff;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s, box-shadow 0.3s;
}

.blog-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.15);
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

.blog-image.cybersecurity {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
}

.blog-image.automation {
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%);
}

.blog-image.digital {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.blog-card .blog-content {
  padding: 25px;
}

.blog-category {
  display: inline-block;
  color: #3b82f6;
  font-size: 12px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 1px;
  margin-bottom: 15px;
}

.blog-card h2 {
  font-size: 20px;
  margin-bottom: 15px;
  color: #1e3a8a;
  line-height: 1.4;
}

.blog-excerpt {
  color: #666;
  font-size: 14px;
  line-height: 1.6;
  margin-bottom: 20px;
}

.blog-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  font-size: 13px;
  color: #888;
}

.read-more {
  color: #3b82f6;
  text-decoration: none;
  font-weight: 600;
  font-size: 14px;
  display: inline-flex;
  align-items: center;
  gap: 5px;
  transition: color 0.3s;
}

.read-more:hover {
  color: #1e40af;
}

.read-more i {
  font-size: 16px;
}

@media (max-width: 768px) {
  .hero-section h1 {
    font-size: 32px;
  }
  
  .blog-grid {
    grid-template-columns: 1fr;
  }
  
  .blog-meta {
    flex-direction: column;
    align-items: flex-start;
    gap: 5px;
  }
}
</style>