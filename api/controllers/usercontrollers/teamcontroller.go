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

// helper to convert string → *string
func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// IndexTeamMembers - GET /team-members
func IndexTeamMembers(c *gin.Context) {
	var teams []models.TeamMember
	if err := models.DB.Find(&teams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"teams": teams})
}

// ShowTeamMember - GET /team-members/:id
func ShowTeamMember(c *gin.Context) {
	var team models.TeamMember
	id := c.Param("id")
	if err := models.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team member not found"})
		return
	}
	c.JSON(http.StatusOK, team)
}

// CreateTeamMember - POST /team-members
func CreateTeamMember(c *gin.Context) {
	var team models.TeamMember

	// Required field
	team.Name = c.PostForm("name")
	if team.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	// Optional fields
	team.Position = strPtr(c.PostForm("position"))
	team.Bio = strPtr(c.PostForm("bio"))
	team.LinkedIn = strPtr(c.PostForm("linkedin"))
	team.GitHub = strPtr(c.PostForm("github"))

	// Handle photo upload
	if file, err := c.FormFile("photo"); err == nil {
		uploadDir := "uploads/team_photos"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("team_%d_%s%s", time.Now().Unix(),
			strings.ReplaceAll(team.Name, " ", "_"), ext)
		fullPath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo"})
			return
		}
		team.Photo = strPtr(fullPath)
	}

	if err := models.DB.Create(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, team)
}

// UpdateTeamMember - PUT /team-members/:id
func UpdateTeamMember(c *gin.Context) {
	var team models.TeamMember
	id := c.Param("id")

	if err := models.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team member not found"})
		return
	}

	// Update fields
	if name := c.PostForm("name"); name != "" {
		team.Name = name
	}
	if position := c.PostForm("position"); position != "" {
		team.Position = strPtr(position)
	}
	team.Bio = strPtr(c.PostForm("bio"))
	team.LinkedIn = strPtr(c.PostForm("linkedin"))
	team.GitHub = strPtr(c.PostForm("github"))

	// Handle photo upload
	if file, err := c.FormFile("photo"); err == nil {
		if team.Photo != nil && *team.Photo != "" {
			if err := os.Remove(*team.Photo); err != nil {
				fmt.Printf("Failed to delete old photo: %v\n", err)
			}
		}

		uploadDir := "uploads/team_photos"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("team_%d_%s%s", time.Now().Unix(),
			strings.ReplaceAll(team.Name, " ", "_"), ext)
		fullPath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo"})
			return
		}
		team.Photo = strPtr(fullPath)
	}

	if err := models.DB.Save(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

// DeleteTeamMember - DELETE /team-members/:id
func DeleteTeamMember(c *gin.Context) {
	var team models.TeamMember
	id := c.Param("id")

	if err := models.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team member not found"})
		return
	}

	if team.Photo != nil && *team.Photo != "" {
		if err := os.Remove(*team.Photo); err != nil {
			fmt.Printf("Failed to delete photo file: %v\n", err)
		}
	}

	if err := models.DB.Delete(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Team member deleted successfully"})
}