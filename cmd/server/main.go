package main

import (
	"golang_blog3/config"
	"golang_blog3/middleware"
	"golang_blog3/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	config.InitDB()

	// 创建 Gin 引擎实例
	r := gin.Default()

	// 设置应用路由（包含公开路由和受保护路由）
	r = routes.SetupRouter()

	// 用于监控系统性能，此路由无需认证
	r.GET("/metrics", middleware.MetricsHandler())

	// 启动 Web 服务器，监听 8080 端口
	// 服务器将处理所有路由（包括公开路由、受保护路由和监控路由）
	r.Run(":8080")
}
