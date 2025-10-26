package usercontroller

import (
    "net/http"
    "go-restapi-gin/models"
    "github.com/gin-gonic/gin"
)

func IndexCompany(c *gin.Context) {
    var companies []models.CompanyProfile
    models.DB.Find(&companies)
    c.JSON(http.StatusOK, gin.H{"data": companies})
}

func ShowCompany(c *gin.Context) {
    var company models.CompanyProfile
    if err := models.DB.First(&company, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": company})
}

func CreateCompany(c *gin.Context) {
    var input models.CompanyProfile
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.DB.Create(&input)
    c.JSON(http.StatusOK, gin.H{"data": input})
}

func UpdateCompany(c *gin.Context) {
    var company models.CompanyProfile
    if err := models.DB.First(&company, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
        return
    }

    var input models.CompanyProfile
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Model(&company).Updates(input)
    c.JSON(http.StatusOK, gin.H{"data": company})
}

func DeleteCompany(c *gin.Context) {
    var company models.CompanyProfile
    if err := models.DB.First(&company, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
        return
    }

    models.DB.Delete(&company)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
