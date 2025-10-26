package usercontroller

import (
	"log"
	"net/http"
	"strconv"

	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

// List semua users
func Users(c *gin.Context) {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		log.Printf("Users DB Find error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// Show user by id
func Show(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
		return
	}

	if err := models.DB.First(&user, id).Error; err != nil {
		log.Printf("Show user error (id=%s): %v\n", id, err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Create user (admin)
func Create(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Create bind error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Basic validation
	if input.Username == "" || input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username, email, password required"})
		return
	}

	if err := models.DB.Create(&input).Error; err != nil {
		log.Printf("Create DB error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": input})
}

// Update user by id
func Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Update bind error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Build update map to avoid overwriting with zero-values (esp. password)
	updates := map[string]interface{}{}
	if input.Username != "" {
		updates["username"] = input.Username
	}
	if input.Email != "" {
		updates["email"] = input.Email
	}
	if input.Password != "" {
		updates["password"] = input.Password
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "nothing to update"})
		return
	}

	res := models.DB.Model(&models.User{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		log.Printf("Update DB error (id=%s): %v\n", id, res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
		return
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

// Delete user by id (gunakan param id, jangan baca body)
func Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
		return
	}

	// convert to int for logging (optional)
	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		log.Printf("Delete parse id error: %v\n", err)
		// tetap coba delete by string id (gorm accepts), tapi beri pesan
	}

	res := models.DB.Delete(&models.User{}, id)
	if res.Error != nil {
		log.Printf("Delete DB error (id=%s): %v\n", id, res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
		return
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
