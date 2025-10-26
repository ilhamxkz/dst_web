package usercontroller

import (
	"go-restapi-gin/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// IndexClients - GET /clients
func IndexClients(c *gin.Context) {
	var clients []models.Client
	if err := models.DB.Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

// ShowClient - GET /clients/:id
func ShowClient(c *gin.Context) {
	var client models.Client
	id := c.Param("id")
	if err := models.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	c.JSON(http.StatusOK, client)
}

// CreateClient - POST /clients (form-data)
func CreateClient(c *gin.Context) {
	name := c.PostForm("name")
	website := c.PostForm("website")

	// simpan logo
	var logoPath string
	file, err := c.FormFile("logo")
	if err == nil {
		os.MkdirAll("uploads/clients", os.ModePerm)
		path := "uploads/clients/" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logoPath = "/" + path
	}

	client := models.Client{Name: name, Website: website, Logo: logoPath}
	if err := models.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, client)
}

// UpdateClient - PUT /clients/:id (form-data)
func UpdateClient(c *gin.Context) {
	var client models.Client
	id := c.Param("id")
	if err := models.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	name := c.PostForm("name")
	website := c.PostForm("website")

	if name != "" {
		client.Name = name
	}
	if website != "" {
		client.Website = website
	}

	file, err := c.FormFile("logo")
	if err == nil {
		os.MkdirAll("uploads/clients", os.ModePerm)
		path := "uploads/clients/" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		client.Logo = "/" + path
	}

	if err := models.DB.Save(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

// DeleteClient - DELETE /clients/:id
func DeleteClient(c *gin.Context) {
	var client models.Client
	id := c.Param("id")
	if err := models.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	if err := models.DB.Delete(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Client deleted"})
}
