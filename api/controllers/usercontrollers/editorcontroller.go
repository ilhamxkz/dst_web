package usercontroller

import (
	"net/http"
	"go-restapi-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IndexPosts - GET /api/posts
func IndexPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// ShowPost - GET /api/posts/:id
func ShowPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

// CreatePost - POST /api/posts
func CreatePost(c *gin.Context) {
	var input struct {
		Title    *string `json:"title"`
		Content  string  `json:"content" binding:"required"`
		AuthorID *int    `json:"author_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	post := models.Post{
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: input.AuthorID,
	}

	if err := models.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// UpdatePost - PUT /api/posts/:id
func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	// pastikan record ada
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post tidak ditemukan"})
		return
	}

	var input struct {
		Title    *string `json:"title"`
		Content  *string `json:"content"`
		AuthorID *int    `json:"author_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// update field yang dikirim
	if input.Title != nil {
		post.Title = input.Title
	}
	if input.Content != nil {
		post.Content = *input.Content
	}
	if input.AuthorID != nil {
		post.AuthorID = input.AuthorID
	}

	if err := models.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post berhasil diperbarui", "post": post})
}

// DeletePost - DELETE /api/posts/:id
func DeletePost(c *gin.Context) {
	id := c.Param("id")

	// konversi id ke int agar bisa digunakan jika perlu
	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID tidak valid"})
		return
	}

	if models.DB.Delete(&models.Post{}, id).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak bisa hapus post (mungkin tidak ada)"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post berhasil dihapus"})
}
