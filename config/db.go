package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang_blog3/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局数据库连接实例
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	// 从环境变量获取数据库连接信息，默认值用于本地开发
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// 如果环境变量未设置，则使用默认值
	if os.Getenv("DB_USER") == "" {
		// ✅ 修改为你的实际数据库信息
		dsn = "root:@tcp(127.0.0.1:3306)/golang_blok3?charset=utf8mb4&parseTime=True&loc=Local" // ✅ 修改这里
	}

	// 配置 GORM 日志级别
	newLogger := logger.Default
	if os.Getenv("GIN_MODE") == "debug" {
		newLogger = newLogger.LogMode(logger.Info)
	}

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("❌ 数据库连接失败:", err)
	}

	// 获取底层 sql.DB 实例以配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取底层数据库连接失败:", err)
	}

	// 配置连接池
	sqlDB.SetMaxIdleConns(10)                  // 设置空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(100)                 // 设置数据库的最大打开连接数
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // 设置连接可以重用的最长时间

	// 将数据库连接赋值给全局变量
	DB = db

	// 自动迁移数据库表结构
	log.Println("🔧 开始执行数据库迁移...")
	err = DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)
	if err != nil {
		log.Fatal("❌ 数据库迁移失败:", err)
	}
	log.Println("✅ 数据库迁移完成")
}
