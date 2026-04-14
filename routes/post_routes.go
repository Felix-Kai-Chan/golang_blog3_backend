package routes

import (
	"golang_blog3/config"
	"golang_blog3/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Preload("User").Preload("Comments").Find(&posts)
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func GetPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	config.DB.Preload("User").Preload("Comments").First(&post, id)

	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func CreatePost(c *gin.Context) {
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		UserID  uint   `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  input.UserID,
	}

	config.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post": post})
}

func UpdatePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	config.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	post.Title = input.Title
	post.Content = input.Content
	config.DB.Save(&post)

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": post})
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.Delete(&models.Post{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
