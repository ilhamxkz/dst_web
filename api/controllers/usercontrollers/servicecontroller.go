package usercontroller

import (
    "net/http"
    "go-restapi-gin/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "fmt"
)

func IndexServices(c *gin.Context) {
    var services []models.Service
    page := c.DefaultQuery("page", "1")
    limit := c.DefaultQuery("limit", "6")

    var p, l int
    fmt.Sscan(page, &p)
    fmt.Sscan(limit, &l)

    offset := (p - 1) * l

    if err := models.DB.Order("created_at DESC").Limit(l).Offset(offset).Find(&services).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"services": services, "page": p})
}



func ShowService(c *gin.Context) {
    id := c.Param("id")
    var service models.Service
    if err := models.DB.First(&service, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"message": "Service tidak ditemukan"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, gin.H{"service": service})
}

func CreateService(c *gin.Context) {
    title := c.PostForm("title")
    description := c.PostForm("description")

    // Upload logo
    var iconPath string
    file, _ := c.FormFile("icon")
    if file != nil {
        iconPath = "uploads/icons/" + file.Filename
        c.SaveUploadedFile(file, iconPath)
    }

    // Upload image
    var imgPath string
    img, _ := c.FormFile("image")
    if img != nil {
        imgPath = "uploads/images/" + img.Filename
        c.SaveUploadedFile(img, imgPath)
    }

    service := models.Service{
        Title:       title,
        Description: &description,
        Icon:        iconPath,
        Images:       imgPath,
    }
    models.DB.Create(&service)
    c.JSON(http.StatusOK, gin.H{"service": service})
}

func UpdateService(c *gin.Context) {
    id := c.Param("id")
    var service models.Service
    if err := models.DB.First(&service, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Service tidak ditemukan"})
        return
    }

    title := c.PostForm("title")
    description := c.PostForm("description")

    if title != "" {
        service.Title = title
    }
    if description != "" {
        service.Description = &description
    }

    // Update logo
    file, _ := c.FormFile("icon")
    if file != nil {
        iconPath := "uploads/icons/" + file.Filename
        c.SaveUploadedFile(file, iconPath)
        service.Icon = iconPath
    }

    // Update image
    img, _ := c.FormFile("image")
    if img != nil {
        imgPath := "uploads/images/" + img.Filename
        c.SaveUploadedFile(img, imgPath)
        service.Images = imgPath
    }

    models.DB.Save(&service)
    c.JSON(http.StatusOK, gin.H{"service": service})
}

func DeleteService(c *gin.Context) {
    id := c.Param("id")
    if models.DB.Delete(&models.Service{}, id).RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak bisa hapus service"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Service berhasil dihapus"})
}
