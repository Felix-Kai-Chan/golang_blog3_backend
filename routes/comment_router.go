package routes

import (
	"golang_blog3/config"
	"golang_blog3/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 使用标准 HTTP 状态码
		return
	}

	// 检查关联的 User 和 Post 是否存在
	var user models.User
	if err := config.DB.First(&user, comment.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var post models.Post
	if err := config.DB.First(&post, comment.PostID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not found"})
		return
	}

	// 保存评论
	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"comment": comment}) // 使用标准 HTTP 状态码和一致的响应格式
}

// GetComment 获取单条评论
func GetComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var comment models.Comment

	if err := config.DB.Preload("User").Preload("Post").First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment}) // 使用一致的响应格式
}

// UpdateComment 更新评论
func UpdateComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var comment models.Comment

	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.Content = input.Content
	if err := config.DB.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment}) // 使用一致的响应格式
}

// DeleteComment 删除评论（软删除）
func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var comment models.Comment

	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if err := config.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// GetCommentsByPost 获取某篇文章的所有评论
func GetCommentsByPost(c *gin.Context) {
	postId, _ := strconv.Atoi(c.Param("postId"))

	// 检查文章是否存在
	var post models.Post
	if err := config.DB.First(&post, postId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var comments []models.Comment
	// 使用 Preload 加载关联的用户信息（可选）
	if err := config.DB.Preload("User").Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments}) // 使用一致的响应格式
}

// GetComments 获取所有评论
func GetComments(c *gin.Context) {
	var comments []models.Comment

	if err := config.DB.Preload("User").Preload("Post").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}
