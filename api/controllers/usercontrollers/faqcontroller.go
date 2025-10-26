package usercontroller

import (
	"go-restapi-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IndexFaqs - GET /api/faqs
func IndexFaqs(c *gin.Context) {
	var faqs []models.Pertanyaan

	// optional ?limit= query param to limit number of returned faqs
	limitStr := c.Query("limit")
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			models.DB.Limit(l).Find(&faqs)
			c.JSON(http.StatusOK, gin.H{"faqs": faqs})
			return
		}
		// if parse failed or non-positive, fall through to return all
	}

	models.DB.Find(&faqs)
	c.JSON(http.StatusOK, gin.H{"faqs": faqs})
}

// ShowFaq - GET /api/faqs/:id
func ShowFaq(c *gin.Context) {
	id := c.Param("id")
	var faq models.Pertanyaan
	if err := models.DB.First(&faq, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "FAQ tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"faq": faq})
}

// CreateFaq - POST /api/faqs
func CreateFaq(c *gin.Context) {
	var input struct {
		Question string `json:"question" binding:"required"`
		Answer   string `json:"answer" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	faq := models.Pertanyaan{
		Question: input.Question,
		Answer:   input.Answer,
	}

	if err := models.DB.Create(&faq).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"faq": faq})
}

// UpdateFaq - PUT /api/faqs/:id
func UpdateFaq(c *gin.Context) {
	id := c.Param("id")

	// pastikan record ada
	var faq models.Pertanyaan
	if err := models.DB.First(&faq, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "FAQ tidak ditemukan"})
		return
	}

	var input struct {
		Question *string `json:"question"`
		Answer   *string `json:"answer"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// update field yang dikirim
	if input.Question != nil {
		faq.Question = *input.Question
	}
	if input.Answer != nil {
		faq.Answer = *input.Answer
	}

	if err := models.DB.Save(&faq).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "FAQ berhasil diperbarui", "faq": faq})
}

// DeleteFaq - DELETE /api/faqs/:id
func DeleteFaq(c *gin.Context) {
	id := c.Param("id")

	// validasi id
	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID tidak valid"})
		return
	}

	if models.DB.Delete(&models.Pertanyaan{}, id).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak bisa hapus FAQ (mungkin tidak ada)"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "FAQ berhasil dihapus"})
}
