<template>
  <div class="faq-page">
    <div class="container">
      <!-- Hero Section -->
      <section class="hero-section">
        <h1>Frequently Asked Questions</h1>
        <p class="lead">Find answers to common questions about our services</p>
      </section>

      <!-- FAQ Content -->
      <section class="faq-content">
        <div class="faq-header">
          <h2>All FAQ Questions</h2>
          <p>Browse through our comprehensive FAQ database to find the information you need.</p>
        </div>

        <!-- Loading State -->
        <div v-if="faqLoading" class="loading-state">
          <div class="spinner"></div>
          <p>Loading FAQ data...</p>
        </div>

        <!-- Error State -->
        <div v-else-if="faqError" class="error-state">
          <div class="error-icon">
            <i class="ri-error-warning-line"></i>
          </div>
          <h3>Error Loading FAQ</h3>
          <p>{{ faqError }}</p>
          <button @click="fetchFaqs" class="btn btn-primary">
            <i class="ri-refresh-line"></i>
            Try Again
          </button>
        </div>

        <!-- DataTable -->
        <div v-else class="faq-table-container">
          <div class="table-responsive">
            <table ref="faqTable" class="table table-striped table-hover" style="width: 100%">
              <thead>
                <tr>
                  <th>No</th>
                  <th>Question</th>
                  <th>Answer</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(faq, index) in faqs" :key="faq.id || faq.ID || index">
                  <td>{{ index + 1 }}</td>
                  <td class="question-cell">
                    <strong>{{ faq.question }}</strong>
                  </td>
                  <td class="answer-cell">
                    <div v-html="truncateText(faq.answer, 100)"></div>
                  </td>
                  <td class="actions-cell">
                    <button 
                      @click="viewFullFaq(faq)" 
                      class="btn btn-sm btn-outline-primary"
                      title="View Full Answer"
                    >
                      <i class="ri-eye-line"></i>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </section>
    </div>

    <!-- FAQ Detail Modal -->
    <div 
      v-if="selectedFaq" 
      class="modal-overlay" 
      @click="closeModal"
    >
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h3>{{ selectedFaq.question }}</h3>
          <button @click="closeModal" class="close-btn">
            <i class="ri-close-line"></i>
          </button>
        </div>
        <div class="modal-body">
          <div v-html="selectedFaq.answer"></div>
        </div>
        <div class="modal-footer">
          <button @click="closeModal" class="btn btn-secondary">
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import $ from 'jquery';

// Make jQuery available globally for DataTables
window.$ = window.jQuery = $;

// Import DataTables CSS
import 'datatables.net-dt/css/dataTables.dataTables.css';

// Import DataTables after jQuery is available
import 'datatables.net';

const API_BASE = 'http://localhost:8080/api';

export default {
  name: 'FAQ',
  data() {
    return {
      faqs: [],
      faqLoading: false,
      faqError: null,
      selectedFaq: null,
      dataTableInstance: null
    };
  },
  mounted() {
    document.title = 'FAQ - CoreTech';
    this.fetchFaqs();
  },
  beforeUnmount() {
    if (this.dataTableInstance) {
      this.dataTableInstance.destroy();
    }
  },
  methods: {
    async fetchFaqs() {
      this.faqLoading = true;
      this.faqError = null;
      
      try {
        // Get token from localStorage if available
        const token = localStorage.getItem('token');
        const headers = token ? { Authorization: `Bearer ${token}` } : {};

        const res = await axios.get(`${API_BASE}/faqs`, {
          headers,
          // Remove limit to get all FAQs
          params: {}
        });

        // Handle different response formats
        const arr = Array.isArray(res.data) 
          ? res.data 
          : Array.isArray(res.data?.faqs) 
            ? res.data.faqs 
            : [];
            
        this.faqs = arr.map((f) => ({
          id: f.id ?? f.ID ?? null,
          question: f.question ?? f.question_text ?? '',
          answer: f.answer ?? f.body ?? '',
        }));

        // Initialize DataTable after data is loaded
        this.$nextTick(() => {
          this.initDataTable();
        });
        
      } catch (e) {
        this.faqError = e?.response?.data?.message || e.message || 'Failed to fetch FAQ data';
        console.error('Error fetching FAQs:', e);
      } finally {
        this.faqLoading = false;
      }
    },

    initDataTable() {
      if (this.$refs.faqTable && this.faqs.length > 0) {
        try {
          // Destroy existing instance if it exists
          if (this.dataTableInstance) {
            this.dataTableInstance.destroy();
            this.dataTableInstance = null;
          }

          // Ensure $ is available
          if (typeof $ !== 'undefined') {
            // Initialize DataTable with custom options
            this.dataTableInstance = $(this.$refs.faqTable).DataTable({
              responsive: true,
              pageLength: 10,
              lengthMenu: [[10, 25, 50, 100, -1], [10, 25, 50, 100, "All"]],
              order: [[0, 'asc']],
              columnDefs: [
                {
                  targets: [2], // Answer column
                  orderable: false,
                  searchable: true
                },
                {
                  targets: [3], // Actions column
                  orderable: false,
                  searchable: false,
                  width: '100px'
                },
                {
                  targets: [0], // No column
                  width: '80px'
                }
              ],
              language: {
                search: "Search FAQ:",
                lengthMenu: "Show _MENU_ entries per page",
                info: "Showing _START_ to _END_ of _TOTAL_ FAQs",
                infoEmpty: "No FAQs available",
                infoFiltered: "(filtered from _MAX_ total FAQs)",
                paginate: {
                  first: "First",
                  last: "Last",
                  next: "Next",
                  previous: "Previous"
                }
              }
            });
          }
        } catch (error) {
          console.error('Error initializing DataTable:', error);
        }
      }
    },

    truncateText(text, maxLength) {
      if (!text) return '';
      
      // Remove HTML tags for length calculation
      const plainText = text.replace(/<[^>]*>/g, '');
      
      if (plainText.length <= maxLength) {
        return text;
      }
      
      // Truncate and add ellipsis
      const truncated = plainText.substring(0, maxLength) + '...';
      return truncated;
    },

    viewFullFaq(faq) {
      this.selectedFaq = faq;
      // Prevent body scroll when modal is open
      document.body.style.overflow = 'hidden';
    },

    closeModal() {
      this.selectedFaq = null;
      // Restore body scroll
      document.body.style.overflow = '';
    }
  }
};
</script>

<style scoped>
.faq-page {
  color: #333;
  background: #fff;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

/* Hero Section */
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

/* FAQ Content */
.faq-content {
  margin-bottom: 80px;
}

.faq-header {
  text-align: center;
  margin-bottom: 50px;
}

.faq-header h2 {
  font-size: 36px;
  margin-bottom: 15px;
  color: #1e3a8a;
}

.faq-header p {
  font-size: 16px;
  color: #666;
  max-width: 600px;
  margin: 0 auto;
}

/* Loading State */
.loading-state {
  text-align: center;
  padding: 60px 20px;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid #f3f4f6;
  border-top: 4px solid #3b82f6;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 20px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.loading-state p {
  color: #666;
  font-size: 16px;
}

/* Error State */
.error-state {
  text-align: center;
  padding: 60px 20px;
}

.error-icon {
  font-size: 48px;
  color: #ef4444;
  margin-bottom: 20px;
}

.error-state h3 {
  color: #ef4444;
  margin-bottom: 15px;
}

.error-state p {
  color: #666;
  margin-bottom: 25px;
}

/* Table Styles */
.faq-table-container {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  padding: 20px;
}

.table-responsive {
  overflow-x: auto;
}

.question-cell {
  max-width: 300px;
}

.answer-cell {
  max-width: 400px;
  font-size: 14px;
  color: #666;
}

.actions-cell {
  text-align: center;
}

/* Button Styles */
.btn {
  padding: 8px 16px;
  border-radius: 6px;
  border: none;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  text-decoration: none;
  transition: all 0.3s;
  display: inline-flex;
  align-items: center;
  gap: 6px;
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
  background: #6b7280;
  color: #fff;
}

.btn-secondary:hover {
  background: #4b5563;
}

.btn-sm {
  padding: 6px 12px;
  font-size: 12px;
}

.btn-outline-primary {
  background: transparent;
  color: #3b82f6;
  border: 1px solid #3b82f6;
}

.btn-outline-primary:hover {
  background: #3b82f6;
  color: #fff;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 10000;
  padding: 20px;
}

.modal-content {
  background: #fff;
  border-radius: 8px;
  max-width: 800px;
  width: 100%;
  max-height: 80vh;
  overflow-y: auto;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.3);
}

.modal-header {
  padding: 25px 30px 15px;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.modal-header h3 {
  margin: 0;
  color: #1e3a8a;
  font-size: 20px;
  line-height: 1.4;
  flex: 1;
  margin-right: 20px;
}

.close-btn {
  background: none;
  border: none;
  font-size: 24px;
  color: #6b7280;
  cursor: pointer;
  padding: 0;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 4px;
  transition: all 0.3s;
}

.close-btn:hover {
  background: #f3f4f6;
  color: #374151;
}

.modal-body {
  padding: 20px 30px;
  line-height: 1.6;
  color: #374151;
}

.modal-footer {
  padding: 15px 30px 25px;
  border-top: 1px solid #e5e7eb;
  text-align: right;
}

@media (max-width: 768px) {
  .hero-section h1 {
    font-size: 32px;
  }
  
  .faq-header h2 {
    font-size: 28px;
  }
  
  .modal-overlay {
    padding: 10px;
  }
  
  .modal-header,
  .modal-body,
  .modal-footer {
    padding-left: 20px;
    padding-right: 20px;
  }
  
  .question-cell,
  .answer-cell {
    max-width: 200px;
  }
}
</style>

<!-- Global DataTables styles -->
<style>
/* Override DataTables default styles */
.dataTables_wrapper {
  margin-top: 20px;
}

.dataTables_filter {
  margin-bottom: 20px;
}

.dataTables_filter input {
  border: 2px solid #e5e7eb;
  border-radius: 6px;
  padding: 8px 12px;
  margin-left: 10px;
  font-size: 14px;
}

.dataTables_filter input:focus {
  outline: none;
  border-color: #3b82f6;
}

.dataTables_length select {
  border: 2px solid #e5e7eb;
  border-radius: 6px;
  padding: 6px 12px;
  margin: 0 10px;
}

.dataTables_info {
  color: #6b7280;
}

.dataTables_paginate .paginate_button {
  padding: 8px 12px;
  margin: 0 2px;
  border-radius: 6px;
  border: 1px solid #e5e7eb;
  color: #374151 !important;
  text-decoration: none;
}

.dataTables_paginate .paginate_button:hover {
  background: #f3f4f6 !important;
  border-color: #d1d5db;
}

.dataTables_paginate .paginate_button.current {
  background: #3b82f6 !important;
  border-color: #3b82f6;
  color: #fff !important;
}

.table {
  border-collapse: separate;
  border-spacing: 0;
}

.table th {
  background: #f8f9fa;
  border-bottom: 2px solid #e5e7eb;
  color: #374151;
  font-weight: 600;
  padding: 12px;
}

.table td {
  padding: 12px;
  border-bottom: 1px solid #e5e7eb;
  vertical-align: top;
}

.table-striped tbody tr:nth-of-type(odd) {
  background: #f8f9fa;
}

.table-hover tbody tr:hover {
  background: #e5f3ff;
}
</style>