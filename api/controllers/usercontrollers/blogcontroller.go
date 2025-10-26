package usercontroller

import (
	"log"
	"net/http"

	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

// -------- Blog Posts CRUD --------

func IndexBlogPosts(c *gin.Context) {
    var posts []models.BlogPost
    if err := models.DB.Preload("Author").Find(&posts).Error; err != nil {
        log.Printf("IndexBlogPosts error: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func ShowBlogPost(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
        return
    }
    var post models.BlogPost
    if err := models.DB.Preload("Author").Preload("Comments").First(&post, id).Error; err != nil {
        log.Printf("ShowBlogPost error (id=%s): %v\n", id, err)
        c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"post": post})
}

func CreateBlogPost(c *gin.Context) {
    var input models.BlogPost
    if err := c.ShouldBindJSON(&input); err != nil {
        log.Printf("CreateBlogPost bind error: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    if input.Title == "" || input.Content == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "title dan content wajib diisi"})
        return
    }

    if err := models.DB.Create(&input).Error; err != nil {
        log.Printf("CreateBlogPost DB error: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"post": input})
}

func UpdateBlogPost(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
        return
    }

    var input models.BlogPost
    if err := c.ShouldBindJSON(&input); err != nil {
        log.Printf("UpdateBlogPost bind error: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    updates := map[string]interface{}{}
    if input.Title != "" {
        updates["title"] = input.Title
    }
    if input.Content != "" {
        updates["content"] = input.Content
    }
    if input.Slug != nil {
        updates["slug"] = input.Slug
    }
    if input.CoverImage != nil {
        updates["cover_image"] = input.CoverImage
    }
    if input.AuthorID != nil {
        updates["author_id"] = input.AuthorID
    }

    if len(updates) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"message": "nothing to update"})
        return
    }

    res := models.DB.Model(&models.BlogPost{}).Where("id = ?", id).Updates(updates)
    if res.Error != nil {
        log.Printf("UpdateBlogPost DB error (id=%s): %v\n", id, res.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
        return
    }
    if res.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate BlogPost"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func DeleteBlogPost(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
        return
    }
    res := models.DB.Delete(&models.BlogPost{}, id)
    if res.Error != nil {
        log.Printf("DeleteBlogPost DB error (id=%s): %v\n", id, res.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
        return
    }
    if res.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus BlogPost"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

// -------- Blog Comments CRUD --------

func IndexBlogComments(c *gin.Context) {
    var comments []models.BlogComment
    if err := models.DB.Preload("Post").Find(&comments).Error; err != nil {
        log.Printf("IndexBlogComments error: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func ShowBlogComment(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
        return
    }
    var comment models.BlogComment
    if err := models.DB.Preload("Post").First(&comment, id).Error; err != nil {
        log.Printf("ShowBlogComment error (id=%s): %v\n", id, err)
        c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"comment": comment})
}

func CreateBlogComment(c *gin.Context) {
    var input models.BlogComment
    if err := c.ShouldBindJSON(&input); err != nil {
        log.Printf("CreateBlogComment bind error: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    // Only require comment content; allow PostID to be null
    if input.Comment == nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "comment wajib diisi"})
        return
    }

    if err := models.DB.Create(&input).Error; err != nil {
        log.Printf("CreateBlogComment DB error: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"comment": input})
}

func UpdateBlogComment(c *gin.Context) {
    id := c.Param("id")

    var input models.BlogComment
    if err := c.ShouldBindJSON(&input); err != nil {
        log.Printf("UpdateBlogComment bind error: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    updates := map[string]interface{}{}
    if input.Name != nil {
        updates["name"] = input.Name
    }
    if input.Email != nil {
        updates["email"] = input.Email
    }
    if input.Comment != nil {
        updates["comment"] = input.Comment
    }

    if len(updates) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"message": "nothing to update"})
        return
    }

    res := models.DB.Model(&models.BlogComment{}).Where("id = ?", id).Updates(updates)
    if res.Error != nil {
        log.Printf("UpdateBlogComment DB error (id=%s): %v\n", id, res.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
        return
    }
    if res.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate BlogComment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func DeleteBlogComment(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
        return
    }
    res := models.DB.Delete(&models.BlogComment{}, id)
    if res.Error != nil {
        log.Printf("DeleteBlogComment DB error (id=%s): %v\n", id, res.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
        return
    }
    if res.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus BlogComment"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}


