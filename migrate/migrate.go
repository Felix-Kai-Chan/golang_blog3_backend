package migrate

import (
	"golang_blog3/config"
	"golang_blog3/models"
	"log"
)

// AutoMigrate 自动创建 User、Post、Comment 表结构
func AutoMigrate() {
	db := config.DB

	log.Println("🔧 正在执行自动迁移...")

	err := db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)
	if err != nil {
		log.Fatalf("❌ 数据库迁移失败: %v", err)
	}

	log.Println("✅ 所有模型表已成功创建/更新")
}
