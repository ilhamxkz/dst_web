package usercontroller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

// IndexHeroBackground - GET /hero_backgrounds
func IndexHeroBackground(c *gin.Context) {
	var heroBackgrounds []models.HeroBackground
	if err := models.DB.Find(&heroBackgrounds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"hero_backgrounds": heroBackgrounds})
}

// ShowHeroBackground - GET /hero_backgrounds/:id
func ShowHeroBackground(c *gin.Context) {
	var heroBackground models.HeroBackground
	id := c.Param("id")
	if err := models.DB.First(&heroBackground, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hero background not found"})
		return
	}
	c.JSON(http.StatusOK, heroBackground)
}

// CreateHeroBackground - POST /hero_backgrounds
func CreateHeroBackground(c *gin.Context) {
	var heroBackground models.HeroBackground

	// Required field
	if file, err := c.FormFile("image"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	} else {
		uploadDir := "uploads/hero_background"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("hero_bg_%d%s", time.Now().Unix(), ext)
		fullPath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
		heroBackground.Image = fullPath
	}

	// Optional field
	heroBackground.AltText = c.PostForm("alt_text")

	if err := models.DB.Create(&heroBackground).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, heroBackground)
}

// UpdateHeroBackground - PUT /hero_backgrounds/:id
func UpdateHeroBackground(c *gin.Context) {
	var heroBackground models.HeroBackground
	id := c.Param("id")

	if err := models.DB.First(&heroBackground, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hero background not found"})
		return
	}

	// Update alt_text
	heroBackground.AltText = c.PostForm("alt_text")

	// Handle image upload
	if file, err := c.FormFile("image"); err == nil {
		if heroBackground.Image != "" {
			if err := os.Remove(heroBackground.Image); err != nil {
				fmt.Printf("Failed to delete old image: %v\n", err)
			}
		}

		uploadDir := "uploads/hero_background"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("hero_bg_%d%s", time.Now().Unix(), ext)
		fullPath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
		heroBackground.Image = fullPath
	}

	if err := models.DB.Save(&heroBackground).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, heroBackground)
}

// DeleteHeroBackground - DELETE /hero_backgrounds/:id
func DeleteHeroBackground(c *gin.Context) {
	var heroBackground models.HeroBackground
	id := c.Param("id")

	if err := models.DB.First(&heroBackground, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hero background not found"})
		return
	}

	if heroBackground.Image != "" {
		if err := os.Remove(heroBackground.Image); err != nil {
			fmt.Printf("Failed to delete image file: %v\n", err)
		}
	}

	if err := models.DB.Delete(&heroBackground).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hero background deleted successfully"})
}
