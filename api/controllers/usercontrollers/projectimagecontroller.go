package usercontroller

import (
	"log"
	"net/http"

	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

// -------- Project Images CRUD --------

func IndexProjectImages(c *gin.Context) {
	var images []models.ProjectImage
	if err := models.DB.Preload("Project").Find(&images).Error; err != nil {
		log.Printf("IndexProjectImages error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"project_images": images})
}

func ShowProjectImage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
		return
	}

	var image models.ProjectImage
	if err := models.DB.Preload("Project").First(&image, id).Error; err != nil {
		log.Printf("ShowProjectImage error (id=%s): %v\n", id, err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"project_image": image})
}

func CreateProjectImage(c *gin.Context) {
	var input models.ProjectImage
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("CreateProjectImage bind error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if input.ProjectID == nil || input.ImageURL == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "project_id dan image_url wajib diisi"})
		return
	}

	if err := models.DB.Create(&input).Error; err != nil {
		log.Printf("CreateProjectImage DB error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// preload project biar langsung ikut di response
	if err := models.DB.Preload("Project").First(&input, input.ID).Error; err != nil {
		log.Printf("CreateProjectImage preload error: %v\n", err)
	}

	c.JSON(http.StatusCreated, gin.H{"project_image": input})
}

func UpdateProjectImage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
		return
	}

	var input models.ProjectImage
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("UpdateProjectImage bind error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if input.ImageURL != nil {
		updates["image_url"] = input.ImageURL
	}
	if input.Caption != nil {
		updates["caption"] = input.Caption
	}
	if input.ProjectID != nil {
		updates["project_id"] = input.ProjectID
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "nothing to update"})
		return
	}

	res := models.DB.Model(&models.ProjectImage{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		log.Printf("UpdateProjectImage DB error (id=%s): %v\n", id, res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
		return
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate ProjectImage"})
		return
	}

	// ambil data terbaru dengan preload project
	var updated models.ProjectImage
	if err := models.DB.Preload("Project").First(&updated, id).Error; err != nil {
		log.Printf("UpdateProjectImage reload error: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{"project_image": updated})
}

func DeleteProjectImage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
		return
	}

	res := models.DB.Delete(&models.ProjectImage{}, id)
	if res.Error != nil {
		log.Printf("DeleteProjectImage DB error (id=%s): %v\n", id, res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
		return
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus ProjectImage"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
