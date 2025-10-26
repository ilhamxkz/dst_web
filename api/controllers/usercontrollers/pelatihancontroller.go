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

// GET /api/pelatihan
func IndexPelatihan(c *gin.Context) {
	var pelatihan []models.Pelatihan
	if err := models.DB.Find(&pelatihan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"pelatihan": pelatihan})
}

// GET /api/pelatihan/:id
func ShowPelatihan(c *gin.Context) {
	var pelatihan models.Pelatihan
	id := c.Param("id")
	if err := models.DB.First(&pelatihan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data pelatihan tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, pelatihan)
}

// POST /api/pelatihan
func CreatePelatihan(c *gin.Context) {
	var pelatihan models.Pelatihan

	pelatihan.Category = c.PostForm("category")
	pelatihan.Jenis = c.PostForm("jenis")
	pelatihan.Deskripsi = c.PostForm("deskripsi")

	// Upload gambar
	file, err := c.FormFile("image")
	if err == nil {
		uploadDir := "uploads/pelatihan_photos"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat folder upload"})
			return
		}

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("pelatihan_%d%s", time.Now().Unix(), ext)
		filePath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
			return
		}
		pelatihan.Image = filePath
	} else {
		pelatihan.Image = "" // jika tidak upload gambar
	}

	if err := models.DB.Create(&pelatihan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pelatihan)
}

// PUT /api/pelatihan/:id
func UpdatePelatihan(c *gin.Context) {
	var pelatihan models.Pelatihan
	id := c.Param("id")

	if err := models.DB.First(&pelatihan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data pelatihan tidak ditemukan"})
		return
	}

	category := c.PostForm("category")
	jenis := c.PostForm("jenis")
	deskripsi := c.PostForm("deskripsi")

	if category != "" {
		pelatihan.Category = category
	}
	if jenis != "" {
		pelatihan.Jenis = jenis
	}
	if deskripsi != "" {
		pelatihan.Deskripsi = deskripsi
	}

	file, err := c.FormFile("image")
	if err == nil {
		if pelatihan.Image != "" {
			_ = os.Remove(pelatihan.Image)
		}
		uploadDir := "uploads/pelatihan_photos"
		os.MkdirAll(uploadDir, 0755)

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("pelatihan_%d%s", time.Now().Unix(), ext)
		filePath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, filePath); err == nil {
			pelatihan.Image = filePath
		}
	}

	if err := models.DB.Save(&pelatihan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pelatihan)
}

// DELETE /api/pelatihan/:id
func DeletePelatihan(c *gin.Context) {
	var pelatihan models.Pelatihan
	id := c.Param("id")

	if err := models.DB.First(&pelatihan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data pelatihan tidak ditemukan"})
		return
	}

	if pelatihan.Image != "" {
		_ = os.Remove(pelatihan.Image)
	}

	if err := models.DB.Delete(&pelatihan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pelatihan berhasil dihapus"})
}
