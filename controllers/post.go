package controllers

import (
	"golang_blog3/config"
	"golang_blog3/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取所有帖子
func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	// 预加载关联的用户和评论
	config.DB.Preload("User").Preload("Comments").Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// 获取单个帖子
func GetPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post

	if err := config.DB.Preload("User").Preload("Comments").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// 创建帖子 (需要认证)
func CreatePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 设置当前登录用户ID
	input.UserID = userID.(uint)

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create post"})
		return
	}

	// 预加载关联的用户信息
	config.DB.Preload("User").First(&input, input.ID)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// 更新帖子 (需要认证并检查权限)
func UpdatePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// 检查当前用户是否有权限更新此帖子
	userID, exists := c.Get("user_id")
	if !exists || post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own posts"})
		return
	}

	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&post).Updates(input)
	config.DB.Preload("User").First(&post, post.ID)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// 删除帖子 (需要认证并检查权限)
func DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// 检查当前用户是否有权限删除此帖子
	userID, exists := c.Get("user_id")
	if !exists || post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own posts"})
		return
	}

	config.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
