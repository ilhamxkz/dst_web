package usercontroller

import (
	"net/http"
	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

// GetAllTemplates - GET /api/templates
// Menampilkan daftar semua template yang tersedia
func GetAllTemplates(c *gin.Context) {
	templates := []gin.H{
		{
			"id":          1,
			"name":        "services",
			"title":       "Services Page Template",
			"description": "Template halaman services dari users.vue dengan 2 kolom",
			"preview":     "Template dengan layout services dan IT infrastructure",
		},
		{
			"id":          2,
			"name":        "contact",
			"title":       "Contact Page Template",
			"description": "Template halaman contact (bisa ditambahkan nanti)",
			"preview":     "Template untuk halaman kontak",
		},
		{
			"id":          3,
			"name":        "about",
			"title":       "About Page Template", 
			"description": "Template halaman about (bisa ditambahkan nanti)",
			"preview":     "Template untuk halaman tentang perusahaan",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"templates": templates,
		"message":   "Daftar template berhasil diambil",
	})
}

// GetUsersTemplate - GET /api/templates/users
// Mengambil HTML template dari users.vue
func GetUsersTemplate(c *gin.Context) {
	// HTML template dari users.vue yang sudah di-compile
	htmlTemplate := `<div class="page">
  <!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <title>Services Management </title>
  <style>
.services-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 1.5rem;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}

.page-header {
  margin-bottom: 2rem;
}

.page-header h1 {
  font-size: 2rem;
  font-weight: 700;
  color: #1f2937;
  margin: 0;
}

/* Controls */
.controls-section {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 1rem;
  margin-bottom: 1.5rem;
  flex-wrap: wrap;
}

.search-controls {
  display: flex;
  gap: 1rem;
  flex: 1;
  min-width: 300px;
}

.search-input {
  flex: 1;
  padding: 0.75rem;
  border: 1px solid #d1d5db;
  border-radius: 0.5rem;
  font-size: 1rem;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.search-input:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.entries-select {
  padding: 0.75rem;
  border: 1px solid #d1d5db;
  border-radius: 0.5rem;
  background-color: white;
  min-width: 140px;
}

/* Buttons */
.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.75rem 1.5rem;
  font-size: 0.95rem;
  font-weight: 500;
  border-radius: 0.5rem;
  border: none;
  cursor: pointer;
  transition: all 0.2s;
  text-decoration: none;
  min-height: 44px;
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.btn-sm {
  padding: 0.5rem 1rem;
  font-size: 0.875rem;
  min-height: 36px;
}

.btn-primary {
  background-color: #3b82f6;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background-color: #2563eb;
  transform: translateY(-1px);
}

.btn-success {
  background-color: #10b981;
  color: white;
}

.btn-success:hover:not(:disabled) {
  background-color: #059669;
}

.btn-danger {
  background-color: #ef4444;
  color: white;
}

.btn-danger:hover:not(:disabled) {
  background-color: #dc2626;
}

.btn-secondary {
  background-color: #6b7280;
  color: white;
}

.btn-secondary:hover:not(:disabled) {
  background-color: #4b5563;
}

.btn-outline {
  background-color: white;
  color: #374151;
  border: 1px solid #d1d5db;
}

.btn-outline:hover:not(:disabled) {
  background-color: #f9fafb;
  border-color: #9ca3af;
}

/* Error Message */
.error-message {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #fef2f2;
  border: 1px solid #fecaca;
  color: #dc2626;
  padding: 1rem;
  border-radius: 0.5rem;
  margin-bottom: 1.5rem;
}

.error-close {
  background: none;
  border: none;
  color: #dc2626;
  font-size: 1.5rem;
  cursor: pointer;
  padding: 0;
  width: 24px;
  height: 24px;
}

/* Loading */
.loading-indicator {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  padding: 2rem;
  color: #6b7280;
}

.spinner {
  width: 20px;
  height: 20px;
  border: 2px solid #e5e7eb;
  border-top: 2px solid #3b82f6;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* Forms */
.form-card {
  background-color: #f9fafb;
  border: 1px solid #e5e7eb;
  border-radius: 0.75rem;
  padding: 1.5rem;
  margin-bottom: 1.5rem;
}

.form-card h3 {
  margin: 0 0 1.5rem 0;
  font-size: 1.25rem;
  font-weight: 600;
  color: #1f2937;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  font-weight: 500;
  color: #374151;
  margin-bottom: 0.5rem;
}

.form-input, .form-textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #d1d5db;
  border-radius: 0.5rem;
  font-size: 1rem;
  transition: border-color 0.2s, box-shadow 0.2s;
  box-sizing: border-box;
}

.form-input:focus, .form-textarea:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.form-textarea {
  resize: vertical;
  min-height: 80px;
}

.form-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1.5rem;
  flex-wrap: wrap;
}

/* Table */
.table-container {
  background-color: white;
  border: 1px solid #e5e7eb;
  border-radius: 0.75rem;
  overflow: hidden;
  margin-bottom: 1.5rem;
}

.services-table {
  width: 100%;
  border-collapse: collapse;
}

.services-table th {
  background-color: #f9fafb;
  font-weight: 600;
  color: #374151;
  padding: 1rem;
  text-align: left;
  border-bottom: 1px solid #e5e7eb;
}

.services-table td {
  padding: 1rem;
  border-bottom: 1px solid #f3f4f6;
  vertical-align: top;
}

.services-table tr:last-child td {
  border-bottom: none;
}

.services-table tbody tr:hover {
  background-color: #f9fafb;
}

.col-number {
  width: 60px;
}

.col-actions {
  width: 160px;
}

.service-name {
  font-weight: 500;
  color: #1f2937;
}

.service-description {
  color: #6b7280;
  max-width: 300px;
  word-wrap: break-word;
}

.actions-cell {
  display: flex;
  gap: 0.5rem;
  flex-wrap: wrap;
}

/* Empty State */
.empty-state {
  text-align: center;
  padding: 3rem 2rem;
  color: #6b7280;
}

.empty-state p {
  font-size: 1.1rem;
  margin: 0;
}

/* Pagination */
.pagination-container {
  display: flex;
  justify-content: center;
  margin-bottom: 1rem;
}

.pagination {
  display: flex;
  gap: 0.5rem;
  align-items: center;
  flex-wrap: wrap;
}

/* Status */
.status-text {
  text-align: center;
  color: #6b7280;
  font-size: 0.95rem;
}

/* Responsive */
@media (max-width: 768px) {
  .services-container {
    padding: 1rem;
  }
  
  .controls-section {
    flex-direction: column;
    align-items: stretch;
  }
  
  .search-controls {
    flex-direction: column;
    min-width: auto;
  }
  
  .services-table {
    font-size: 0.875rem;
  }
  
  .services-table th,
  .services-table td {
    padding: 0.75rem 0.5rem;
  }
  
  .form-actions {
    flex-direction: column;
  }
  
  .actions-cell {
    flex-direction: column;
  }
  
  .pagination {
    justify-content: center;
    gap: 0.25rem;
  }
}
  </style>
</head>
<body>
  <div class="services-container">
    <header class="page-header">
      <h1>Services Management</h1>
    </header>

    <!-- Controls -->
    <div class="controls-section">
      <div class="search-controls">
        <input id="searchInput" type="text" class="search-input" placeholder="Search services by name or description...">
        <select id="entriesSelect" class="entries-select">
          <option value="5">5 per page</option>
          <option value="10" selected>10 per page</option>
          <option value="25">25 per page</option>
          <option value="50">50 per page</option>
        </select>
      </div>
      
      <button id="btnAddNew" class="btn btn-primary">Add New Service</button>
    </div>

    <!-- Error Display -->
    <div id="errorMessage" class="error-message" style="display:none;">
      <span id="errorText"></span>
      <button id="errorClose" class="error-close">&times;</button>
    </div>

    <!-- Loading Indicator -->
    <div id="loadingIndicator" class="loading-indicator" style="display:none;">
      <div class="spinner"></div>
      <span>Loading...</span>
    </div>

    <!-- Add Service Form -->
    <div id="addForm" class="form-card" style="display:none;">
      <h3>Add New Service</h3>
      <form id="addFormElement">
        <div class="form-group">
          <label for="new-name">Name *</label>
          <input id="new-name" type="text" class="form-input" placeholder="Enter service name" maxlength="100" required>
        </div>
        <div class="form-group">
          <label for="new-description">Description *</label>
          <textarea id="new-description" class="form-textarea" placeholder="Enter service description" maxlength="500" rows="3" required></textarea>
        </div>
        <div class="form-actions">
          <button id="btnSaveNew" type="submit" class="btn btn-success">Save Service</button>
          <button id="btnCancelAdd" type="button" class="btn btn-secondary">Cancel</button>
        </div>
      </form>
    </div>

    <!-- Edit Service Form -->
    <div id="editForm" class="form-card" style="display:none;">
      <h3>Edit Service</h3>
      <form id="editFormElement">
        <input type="hidden" id="edit-id">
        <div class="form-group">
          <label for="edit-name">Name *</label>
          <input id="edit-name" type="text" class="form-input" placeholder="Enter service name" maxlength="100" required>
        </div>
        <div class="form-group">
          <label for="edit-description">Description *</label>
          <textarea id="edit-description" class="form-textarea" placeholder="Enter service description" maxlength="500" rows="3" required></textarea>
        </div>
        <div class="form-actions">
          <button id="btnUpdate" type="submit" class="btn btn-success">Update Service</button>
          <button id="btnCancelEdit" type="button" class="btn btn-secondary">Cancel</button>
        </div>
      </form>
    </div>

    <!-- Services Table -->
    <div class="table-container">
      <table id="servicesTable" class="services-table" style="display:none;">
        <thead>
          <tr>
            <th class="col-number">#</th>
            <th class="col-name">Name</th>
            <th class="col-description">Description</th>
            <th class="col-actions">Actions</th>
          </tr>
        </thead>
        <tbody id="servicesTbody"></tbody>
      </table>
      
      <div id="emptyState" class="empty-state" style="display:none;">
        <p id="emptyText"></p>
      </div>
    </div>

    <!-- Pagination -->
    <div id="paginationWrapper" class="pagination-container" style="display:none;">
      <div id="pagination" class="pagination"></div>
    </div>

    <!-- Status Text -->
    <div id="statusText" class="status-text" style="margin-top:8px;"></div>
  </div>

</body>
</html>
`

	c.JSON(http.StatusOK, gin.H{
		"success":       true,
		"html":          htmlTemplate,
		"template_name": "services",
		"template_id":   1,
		"message":       "Template HTML services berhasil diambil",
	})
}

// GetContactTemplate - GET /api/templates/contact (contoh template lain)
func GetContactTemplate(c *gin.Context) {
	// Contoh template contact sederhana
	htmlTemplate := `<div class="contact-page">
  <div class="container">
    <h1>Contact Us</h1>
    
    <div class="contact-grid">
      <div class="contact-info">
        <h3>Get in Touch</h3>
        <p>We'd love to hear from you. Send us a message and we'll respond as soon as possible.</p>
        
        <div class="contact-item">
          <h4>Address</h4>
          <p>123 Business Street<br>Jakarta, Indonesia</p>
        </div>
        
        <div class="contact-item">
          <h4>Phone</h4>
          <p>+62 21 1234 5678</p>
        </div>
        
        <div class="contact-item">
          <h4>Email</h4>
          <p>info@dayasinergi.com</p>
        </div>
      </div>
      
      <div class="contact-form">
        <h3>Send Message</h3>
        <form>
          <div class="form-group">
            <label>Name</label>
            <input type="text" placeholder="Your Name">
          </div>
          <div class="form-group">
            <label>Email</label>
            <input type="email" placeholder="Your Email">
          </div>
          <div class="form-group">
            <label>Message</label>
            <textarea placeholder="Your Message" rows="5"></textarea>
          </div>
          <button type="submit">Send Message</button>
        </form>
      </div>
    </div>
  </div>
</div>

<style>
.contact-page { padding: 40px 0; font-family: Arial, sans-serif; }
.container { max-width: 1200px; margin: 0 auto; padding: 0 20px; }
.contact-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 40px; margin-top: 30px; }
.contact-info h3, .contact-form h3 { color: #333; margin-bottom: 20px; }
.contact-item { margin-bottom: 20px; }
.contact-item h4 { color: #5bc34a; margin-bottom: 5px; }
.form-group { margin-bottom: 15px; }
.form-group label { display: block; margin-bottom: 5px; font-weight: bold; }
.form-group input, .form-group textarea { width: 100%; padding: 10px; border: 1px solid #ddd; border-radius: 5px; }
button { background: #5bc34a; color: white; padding: 12px 20px; border: none; border-radius: 5px; cursor: pointer; }
button:hover { background: #4a9c3a; }
@media (max-width: 768px) { .contact-grid { grid-template-columns: 1fr; } }
</style>`

	c.JSON(http.StatusOK, gin.H{
		"success":       true,
		"html":          htmlTemplate,
		"template_name": "contact",
		"template_id":   2,
		"message":       "Template HTML contact berhasil diambil",
	})
}

// GetAboutTemplate - GET /api/templates/about (contoh template lain)
func GetAboutTemplate(c *gin.Context) {
	// Contoh template about sederhana
	htmlTemplate := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>About Daya Sinergi Teknomadiri</title>

		<style>
.about-page { padding: 40px 0; font-family: Arial, sans-serif; }
.container { max-width: 1000px; margin: 0 auto; padding: 0 20px; }
.about-content section { margin-bottom: 40px; }
.mv-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 30px; margin-top: 20px; }
.mission, .vision { padding: 20px; background: #f8f9fa; border-radius: 8px; }
.values-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 20px; margin-top: 20px; }
.value-item { padding: 20px; border: 1px solid #e9ecef; border-radius: 8px; text-align: center; }
.value-item h4 { color: #5bc34a; margin-bottom: 10px; }
h1 { color: #333; text-align: center; margin-bottom: 40px; }
h2 { color: #333; border-bottom: 2px solid #5bc34a; padding-bottom: 10px; }
h3 { color: #5bc34a; }
@media (max-width: 768px) { .mv-grid { grid-template-columns: 1fr; } }
</style>
	</head>
	<body>
	<div class="about-page">
  <div class="container">
    <h1>About Daya Sinergi Teknomadiri</h1>
    
    <div class="about-content">
      <section class="company-intro">
        <h2>Who We Are</h2>
        <p>Daya Sinergi Teknomadiri adalah perusahaan teknologi yang berfokus pada pengembangan solusi software dan layanan IT infrastructure. Dengan pengalaman bertahun-tahun, kami telah membantu berbagai perusahaan dalam transformasi digital mereka.</p>
      </section>
      
      <section class="mission-vision">
        <div class="mv-grid">
          <div class="mission">
            <h3>Our Mission</h3>
            <p>Memberikan solusi teknologi terbaik yang dapat meningkatkan efisiensi dan produktivitas bisnis klien melalui inovasi dan layanan berkualitas tinggi.</p>
          </div>
          <div class="vision">
            <h3>Our Vision</h3>
            <p>Menjadi partner teknologi terpercaya yang membantu perusahaan mencapai kesuksesan melalui implementasi teknologi yang tepat dan berkelanjutan.</p>
          </div>
        </div>
      </section>
      
      <section class="values">
        <h2>Our Values</h2>
        <div class="values-grid">
          <div class="value-item">
            <h4>Innovation</h4>
            <p>Selalu mencari solusi terbaru dan terdepan dalam teknologi</p>
          </div>
          <div class="value-item">
            <h4>Quality</h4>
            <p>Berkomitmen memberikan hasil kerja dengan standar kualitas tertinggi</p>
          </div>
          <div class="value-item">
            <h4>Partnership</h4>
            <p>Membangun hubungan jangka panjang dengan klien sebagai partner bisnis</p>
          </div>
        </div>
      </section>
    </div>
  </div>
</div>
</body>
</html>`

	c.JSON(http.StatusOK, gin.H{
		"success":       true,
		"html":          htmlTemplate,
		"template_name": "about",
		"template_id":   3,
		"message":       "Template HTML about berhasil diambil",
	})
}

// SaveTemplateAsPost - POST /api/templates/save
// Menyimpan template yang sudah diedit sebagai post
func SaveTemplateAsPost(c *gin.Context) {
	var input struct {
		Title        *string `json:"title"`
		Content      string  `json:"content" binding:"required"`
		AuthorID     *int    `json:"author_id"`
		TemplateName *string `json:"template_name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data input tidak valid: " + err.Error(),
		})
		return
	}

	// Jika tidak ada title, buat default berdasarkan template
	if input.Title == nil {
		defaultTitle := "Edited Template"
		if input.TemplateName != nil {
			defaultTitle = "Edited " + *input.TemplateName + " Template"
		}
		input.Title = &defaultTitle
	}

	post := models.Post{
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: input.AuthorID,
	}

	if err := models.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal menyimpan ke database: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Template berhasil disimpan sebagai post",
		"post":    post,
		"post_id": post.ID,
	})
}

// GetTemplateByName - GET /api/templates/:name (dynamic template getter)
func GetTemplateByName(c *gin.Context) {
	templateName := c.Param("name")

	switch templateName {
	case "services":
		GetUsersTemplate(c)
	case "contact":
		GetContactTemplate(c)
	case "about":
		GetAboutTemplate(c)
	default:
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Template '" + templateName + "' tidak ditemukan",
			"available_templates": []string{"users", "contact", "about"},
		})
	}
}