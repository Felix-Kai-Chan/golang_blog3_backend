package controllers

import (
	"golang_blog3/config"
	"golang_blog3/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取所有评论
func GetAllComments(c *gin.Context) {
	var comments []models.Comment

	// 预加载关联的用户和帖子
	config.DB.Preload("User").Preload("Post").Find(&comments)
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// 获取单个评论
func GetComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var comment models.Comment

	if err := config.DB.Preload("User").Preload("Post").First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// 创建评论 (需要认证)
func CreateComment(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 设置当前登录用户ID
	input.UserID = userID.(uint)

	// 验证帖子是否存在
	var post models.Post
	if err := config.DB.First(&post, input.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create comment"})
		return
	}

	// 预加载关联的用户和帖子信息
	config.DB.Preload("User").Preload("Post").First(&input, input.ID)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// 更新评论 (需要认证并检查权限)
func UpdateComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var comment models.Comment

	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// 检查当前用户是否有权限更新此评论
	userID, exists := c.Get("user_id")
	if !exists || comment.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own comments"})
		return
	}

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&comment).Updates(input)
	config.DB.Preload("User").Preload("Post").First(&comment, comment.ID)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// 删除评论 (需要认证并检查权限)
func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var comment models.Comment

	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// 检查当前用户是否有权限删除此评论
	userID, exists := c.Get("user_id")
	if !exists || comment.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own comments"})
		return
	}

	config.DB.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
