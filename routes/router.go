package routes

import (
	"golang_blog3/config"      // 引入 config.DB
	"golang_blog3/controllers" // 引入 UsersController
	"golang_blog3/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.PrometheusMetrics())

	// 只创建 UsersController
	userCtrl := controllers.NewUsersController(config.DB)

	// =============== 公开路由 ===============
	// ✅ register 用 controller
	r.POST("/register", userCtrl.CreateUser)

	// ❌ login 不用 controller！继续用原来的 Login 函数
	r.POST("/login", Login)

	// 帖子列表（公开）
	r.GET("/posts", GetPosts)
	// 👇 保留公开单个帖子访问（按你要求）
	r.GET("/posts/:id", GetPost) // // 公开 ← 保留！

	// 评论（公开）
	r.GET("/comments", GetComments)
	r.GET("/comments/:id", GetComment)

	// =============== 受保护路由（/api）===============
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		// ✅ 用户相关操作全部用 controller
		protected.GET("/users/:id", userCtrl.GetUser)
		protected.PUT("/users/:id", userCtrl.UpdateUser)
		protected.DELETE("/users/:id", userCtrl.DeleteUser)

		// 帖子：增删改 + **受保护的 GET**
		protected.POST("/posts", CreatePost)
		protected.GET("/posts/:id", GetPost) // // 受保护（冗余路径，但允许）← 保留！
		protected.PUT("/posts/:id", UpdatePost)
		protected.DELETE("/posts/:id", DeletePost)

		// 评论
		protected.POST("/comments", CreateComment)
		protected.PUT("/comments/:id", UpdateComment)
		protected.DELETE("/comments/:id", DeleteComment)
	}

	return r
}
