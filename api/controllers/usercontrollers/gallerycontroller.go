package usercontroller

import (
	"go-restapi-gin/models"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadClientLogo handles upload logo client, returns URL
func UploadClientLogo(c *gin.Context) {
	file, err := c.FormFile("logo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, os.ModePerm)
	}
	dstPath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, dstPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// URL bisa disesuaikan sesuai static serving
	url := "/uploads/" + file.Filename
	c.JSON(http.StatusOK, gin.H{"url": url})
}

// IndexGalleries - GET /galleries
func IndexGalleries(c *gin.Context) {
	var galleries []models.Gallery
	if err := models.DB.Find(&galleries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, galleries)
}

// ShowGallery - GET /galleries/:id
func ShowGallery(c *gin.Context) {
	var gallery models.Gallery
	id := c.Param("id")
	if err := models.DB.First(&gallery, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
		return
	}
	c.JSON(http.StatusOK, gallery)
}

// CreateGallery - POST /galleries
func CreateGallery(c *gin.Context) {
	var input models.Gallery
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// UpdateGallery - PUT /galleries/:id
func UpdateGallery(c *gin.Context) {
	var gallery models.Gallery
	id := c.Param("id")
	if err := models.DB.First(&gallery, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
		return
	}
	var input models.Gallery
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gallery.Title = input.Title
	gallery.ImageURL = input.ImageURL
	gallery.Caption = input.Caption
	if err := models.DB.Save(&gallery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gallery)
}

// DeleteGallery - DELETE /galleries/:id
func DeleteGallery(c *gin.Context) {
	var gallery models.Gallery
	id := c.Param("id")
	if err := models.DB.First(&gallery, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
		return
	}
	if err := models.DB.Delete(&gallery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Gallery deleted"})
}
