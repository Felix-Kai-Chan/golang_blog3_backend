package routes

import (
	"github.com/gin-gonic/gin"
	"golang_blog3/config"      // 引入 config.DB
	"golang_blog3/controllers" // 引入 UsersController
	"golang_blog3/middleware"
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

	// ❌ login 不用 controller！继续用原来的 Login 函数（因为你没写 Login 方法）
	r.POST("/login", Login) //

	// 帖子和评论
	r.GET("/posts", GetPosts)
	r.GET("/posts/:id", GetPost)
	r.GET("/comments", GetComments)
	r.GET("/comments/:id", GetComment)

	// =============== 受保护路由（/api）===============
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		// ✅ 用户相关操作全部用 controller（GetUser, UpdateUser, DeleteUser）
		protected.GET("/users/:id", userCtrl.GetUser)
		protected.PUT("/users/:id", userCtrl.UpdateUser)
		protected.DELETE("/users/:id", userCtrl.DeleteUser)

		// 帖子和评论
		protected.POST("/posts", CreatePost)
		protected.PUT("/posts/:id", UpdatePost)
		protected.DELETE("/posts/:id", DeletePost)

		protected.POST("/comments", CreateComment)
		protected.PUT("/comments/:id", UpdateComment)
		protected.DELETE("/comments/:id", DeleteComment)
	}

	return r
}
