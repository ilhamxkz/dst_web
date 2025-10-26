<template>
  <div class="not-found-page">
    <div class="container">
      <div class="not-found-content">
        <div class="error-code">404</div>
        <h1>Page Not Found</h1>
        <p>Sorry, the page you are looking for doesn't exist or has been moved.</p>
        <div class="actions">
          <router-link to="/" class="btn btn-primary">
            <i class="ri-home-line"></i>
            Go to Homepage
          </router-link>
          <router-link to="/contact" class="btn btn-secondary">
            <i class="ri-mail-line"></i>
            Contact Support
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NotFound',
  mounted() {
    document.title = 'Page Not Found - CoreTech';
  }
};
</script>

<style scoped>
.not-found-page {
  min-height: 60vh;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #333;
  background: #fff;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.not-found-content {
  text-align: center;
  max-width: 500px;
  margin: 0 auto;
}

.error-code {
  font-size: 120px;
  font-weight: 700;
  color: #3b82f6;
  line-height: 1;
  margin-bottom: 20px;
  text-shadow: 0 4px 8px rgba(59, 130, 246, 0.3);
}

.not-found-content h1 {
  font-size: 36px;
  margin-bottom: 20px;
  color: #1e3a8a;
}

.not-found-content p {
  font-size: 16px;
  color: #666;
  margin-bottom: 40px;
  line-height: 1.6;
}

.actions {
  display: flex;
  gap: 20px;
  justify-content: center;
  flex-wrap: wrap;
}

.btn {
  padding: 14px 28px;
  border-radius: 8px;
  border: none;
  cursor: pointer;
  font-size: 14px;
  font-weight: 600;
  text-decoration: none;
  transition: all 0.3s;
  display: inline-flex;
  align-items: center;
  gap: 8px;
}

.btn-primary {
  background: linear-gradient(90deg, #1e3a8a 0%, #3b82f6 60%);
  color: #fff;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

.btn-secondary {
  background: transparent;
  color: #3b82f6;
  border: 2px solid #3b82f6;
}

.btn-secondary:hover {
  background: #3b82f6;
  color: #fff;
  transform: translateY(-2px);
}

@media (max-width: 768px) {
  .error-code {
    font-size: 80px;
  }
  
  .not-found-content h1 {
    font-size: 28px;
  }
  
  .actions {
    flex-direction: column;
    align-items: center;
  }
  
  .btn {
    min-width: 200px;
    justify-content: center;
  }
}
</style>