package usercontroller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

// IndexTestimoni - GET /testimonis
func IndexTestimoni(c *gin.Context) {
	var testimonials []models.Testimoni
	if err := models.DB.Find(&testimonials).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"testimonials": testimonials})
}

// ShowTestimoni - GET /testimonis/:id
func ShowTestimoni(c *gin.Context) {
	var testimonial models.Testimoni
	id := c.Param("id")
	if err := models.DB.First(&testimonial, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimonial not found"})
		return
	}
	c.JSON(http.StatusOK, testimonial)
}

// CreateTestimoni - POST /testimonis
func CreateTestimoni(c *gin.Context) {
	var testimonial models.Testimoni

	// Required fields
	testimonial.Nama = c.PostForm("nama")
	if testimonial.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama is required"})
		return
	}

	testimonial.Pesan = c.PostForm("pesan")
	if testimonial.Pesan == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pesan is required"})
		return
	}

	// Optional fields
	testimonial.Perusahaan = c.PostForm("perusahaan")

	// Handle photo upload
	if file, err := c.FormFile("url_foto_profil"); err == nil {
		uploadDir := "uploads/testimonial photo"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("testimonial_%d_%s%s", time.Now().Unix(),
			strings.ReplaceAll(testimonial.Nama, " ", "_"), ext)
		fullPath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo"})
			return
		}
		testimonial.UrlFotoProfil = fullPath
	}

	if err := models.DB.Create(&testimonial).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, testimonial)
}

// UpdateTestimoni - PUT /testimonis/:id
func UpdateTestimoni(c *gin.Context) {
	var testimonial models.Testimoni
	id := c.Param("id")

	if err := models.DB.First(&testimonial, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimonial not found"})
		return
	}

	// Update fields
	if nama := c.PostForm("nama"); nama != "" {
		testimonial.Nama = nama
	}
	if pesan := c.PostForm("pesan"); pesan != "" {
		testimonial.Pesan = pesan
	}
	testimonial.Perusahaan = c.PostForm("perusahaan")

	// Handle photo upload
	if file, err := c.FormFile("url_foto_profil"); err == nil {
		if testimonial.UrlFotoProfil != "" {
			if err := os.Remove(testimonial.UrlFotoProfil); err != nil {
				fmt.Printf("Failed to delete old photo: %v\n", err)
			}
		}

		uploadDir := "uploads/testimonial photo"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("testimonial_%d_%s%s", time.Now().Unix(),
			strings.ReplaceAll(testimonial.Nama, " ", "_"), ext)
		fullPath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo"})
			return
		}
		testimonial.UrlFotoProfil = fullPath
	}

	if err := models.DB.Save(&testimonial).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, testimonial)
}

// DeleteTestimoni - DELETE /testimonis/:id
func DeleteTestimoni(c *gin.Context) {
	var testimonial models.Testimoni
	id := c.Param("id")

	if err := models.DB.First(&testimonial, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimonial not found"})
		return
	}

	if testimonial.UrlFotoProfil != "" {
		if err := os.Remove(testimonial.UrlFotoProfil); err != nil {
			fmt.Printf("Failed to delete photo file: %v\n", err)
		}
	}

	if err := models.DB.Delete(&testimonial).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Testimonial deleted successfully"})
}