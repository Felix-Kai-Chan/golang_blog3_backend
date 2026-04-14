package models

import (
	"time"

	"gorm.io/gorm"
)

// SetupHooks 设置模型钩子
func SetupHooks(db *gorm.DB) error {
	// 如果有需要设置的联接表，可以在这里添加
	// 目前没有定义 UserFollower 模型，所以暂时注释掉相关代码
	/*
		err := db.SetupJoinTable(&User{}, "Followers", &UserFollower{})
		if err != nil {
			return err
		}
		err = db.SetupJoinTable(&User{}, "Following", &UserFollower{})
		if err != nil {
			return err
		}
	*/
	return nil
}

// User 模型的钩子函数
func (u *User) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

// Post 模型的钩子函数
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
	return nil
}

func (p *Post) BeforeUpdate(tx *gorm.DB) error {
	p.UpdatedAt = time.Now()
	return nil
}

// Comment 模型的钩子函数
func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now
	return nil
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) error {
	c.UpdatedAt = time.Now()
	return nil
}
