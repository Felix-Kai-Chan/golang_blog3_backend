package routes

import (
	"github.com/gin-gonic/gin"
	"golang_blog3/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 在所有路由注册前
	r.Use(middleware.PrometheusMetrics())

	// 公开路由 - 不需要JWT认证
	r.POST("/register", Register)      // 用户注册 - 任何人都可以访问
	r.POST("/login", Login)            // 用户登录 - 任何人都可以访问
	r.GET("/posts", GetPosts)          // 查看所有帖子 - 公开访问
	r.GET("/posts/:id", GetPost)       // 查看单个帖子 - 公开访问
	r.GET("/comments", GetComments)    // 查看所有评论 - 公开访问
	r.GET("/comments/:id", GetComment) // 查看单条评论 - 公开访问

	// 受保护的路由 - 需要有效的JWT Token才能访问
	protected := r.Group("/api")        // 创建一个受保护的路由组
	protected.Use(middleware.JWTAuth()) // 应用JWT认证中间件到整个组
	{
		// 用户相关的受保护操作
		r.GET("/users/:id", GetUser)               // 查看用户信息 - 需要登录
		protected.PUT("/users/:id", UpdateUser)    // 修改用户信息 - 需要登录
		protected.DELETE("/users/:id", DeleteUser) // 删除用户账号 - 需要登录

		// 帖子相关的受保护操作
		protected.POST("/posts", CreatePost)       // 发布新帖子 - 需要登录
		protected.PUT("/posts/:id", UpdatePost)    // 修改帖子 - 需要登录
		protected.DELETE("/posts/:id", DeletePost) // 删除帖子 - 需要登录

		// 评论相关的受保护操作
		protected.POST("/comments", CreateComment)       // 发表评论 - 需要登录
		protected.PUT("/comments/:id", UpdateComment)    // 修改评论 - 需要登录
		protected.DELETE("/comments/:id", DeleteComment) // 删除评论 - 需要登录
	}

	return r
}
