package usercontroller

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ==========================
// GET ALL EXPERIENCES
// ==========================
func IndexExperiences(c *gin.Context) {
	var experiences []models.Experience
	if err := models.DB.Find(&experiences).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i := range experiences {
		if experiences[i].CoverImage != "" && !strings.HasPrefix(experiences[i].CoverImage, "http") {
			experiences[i].CoverImage = "http://localhost:8080/" + experiences[i].CoverImage
		}
	}
	c.JSON(http.StatusOK, experiences)
}

// ==========================
// GET SINGLE EXPERIENCE BY ID
// ==========================
func ShowExperience(c *gin.Context) {
	id := c.Param("id")
	var experience models.Experience
	if err := models.DB.First(&experience, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if experience.CoverImage != "" && !strings.HasPrefix(experience.CoverImage, "http") {
		experience.CoverImage = "http://localhost:8080/" + experience.CoverImage
	}

	c.JSON(http.StatusOK, experience)
}

// ==========================
// CREATE NEW EXPERIENCE
// ==========================
func CreateExperience(c *gin.Context) {
	title := c.PostForm("title")
	clientName := c.PostForm("client_name")
	description := c.PostForm("description")
	startDateStr := c.PostForm("start_date")
	endDateStr := c.PostForm("end_date")
	status := c.PostForm("status")
	yearStr := c.PostForm("year") // ✅ ambil dari form

	// === Handle tanggal ===
	layout := "2006-01-02"
	var startDate, endDate *time.Time

	if startDateStr != "" {
		if parsed, err := time.Parse(layout, startDateStr); err == nil {
			startDate = &parsed
		}
	}
	if endDateStr != "" {
		if parsed, err := time.Parse(layout, endDateStr); err == nil {
			endDate = &parsed
		}
	}

	// === Handle Year ===
	var yearValue *int
	if yearStr != "" {
		if y, err := strconv.Atoi(yearStr); err == nil {
			yearValue = &y
		}
	}

	// === Upload cover image ===
	var coverImagePath string
	file, _ := c.FormFile("cover_image")
	if file != nil {
		coverImagePath = "uploads/images/" + file.Filename
		c.SaveUploadedFile(file, coverImagePath)
	}

	experience := models.Experience{
		Title:       title,
		ClientName:  clientName,
		Description: &description,
		StartDate:   startDate,
		EndDate:     endDate,
		Status:      status,
		Year:        yearValue, // ✅ ditambahkan
		CoverImage:  coverImagePath,
	}

	if err := models.DB.Create(&experience).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Experience created successfully",
		"data":    experience,
	})
}

// ==========================
// UPDATE EXPERIENCE
// ==========================
func UpdateExperience(c *gin.Context) {
	id := c.Param("id")
	var experience models.Experience

	if err := models.DB.First(&experience, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		return
	}

	title := c.PostForm("title")
	clientName := c.PostForm("client_name")
	description := c.PostForm("description")
	startDateStr := c.PostForm("start_date")
	endDateStr := c.PostForm("end_date")
	status := c.PostForm("status")
	yearStr := c.PostForm("year") // ✅ ambil dari form

	layout := "2006-01-02"
	var startDate, endDate *time.Time

	if startDateStr != "" {
		if parsed, err := time.Parse(layout, startDateStr); err == nil {
			startDate = &parsed
		}
	}
	if endDateStr != "" {
		if parsed, err := time.Parse(layout, endDateStr); err == nil {
			endDate = &parsed
		}
	}

	// === Update field ===
	if title != "" {
		experience.Title = title
	}
	if clientName != "" {
		experience.ClientName = clientName
	}
	if description != "" {
		experience.Description = &description
	}
	if startDate != nil {
		experience.StartDate = startDate
	}
	if endDate != nil {
		experience.EndDate = endDate
	}
	if status != "" {
		experience.Status = status
	}

	// === Update Year ===
	if yearStr != "" {
		if y, err := strconv.Atoi(yearStr); err == nil {
			experience.Year = &y
		}
	}

	// === Update cover image ===
	file, _ := c.FormFile("cover_image")
	if file != nil {
		coverImagePath := "uploads/images/" + file.Filename
		c.SaveUploadedFile(file, coverImagePath)
		experience.CoverImage = coverImagePath
	}

	if err := models.DB.Save(&experience).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Experience updated successfully",
		"data":    experience,
	})
}

// ==========================
// DELETE EXPERIENCE
// ==========================
func DeleteExperience(c *gin.Context) {
	id := c.Param("id")
	var experience models.Experience

	if err := models.DB.First(&experience, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		return
	}

	if err := models.DB.Delete(&experience).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Experience deleted successfully"})
}
