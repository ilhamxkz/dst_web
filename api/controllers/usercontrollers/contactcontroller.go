package usercontroller

import (
    "net/http"
    "go-restapi-gin/models"

    "github.com/gin-gonic/gin"
)

// Ambil semua pesan
func IndexContact(c *gin.Context) {
    var contacts []models.ContactMessage
    models.DB.Find(&contacts)
    c.JSON(http.StatusOK, gin.H{"data": contacts})
}

// Ambil pesan by ID
func ShowContact(c *gin.Context) {
    var contact models.ContactMessage
    if err := models.DB.First(&contact, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": contact})
}

// Buat pesan baru
func CreateContact(c *gin.Context) {
    var input models.ContactMessage
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.DB.Create(&input)
    c.JSON(http.StatusOK, gin.H{"data": input})
}

// Update pesan
func UpdateContact(c *gin.Context) {
    var contact models.ContactMessage
    if err := models.DB.First(&contact, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    var input models.ContactMessage
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Model(&contact).Updates(input)
    c.JSON(http.StatusOK, gin.H{"data": contact})
}

// Hapus pesan
func DeleteContact(c *gin.Context) {
    var contact models.ContactMessage
    if err := models.DB.First(&contact, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    models.DB.Delete(&contact)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
