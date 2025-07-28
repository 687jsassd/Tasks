package models

import (
	"time"

	"gorm.io/gorm"
)

// User 数据模型定义
// 存储用户基本信息和认证数据
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:64;not null" json:"name"`
	Email     string         `gorm:"size:128;uniqueIndex;not null" json:"email"`
	Phone     string         `gorm:"size:20;uniqueIndex;not null" json:"phone"`
	Password  string         `gorm:"size:128;not null" json:"-"`     // 密码不返回给前端
	Role      int            `gorm:"not null;default:0" json:"role"` // 0:普通用户, 1:管理员
	IsActive  bool           `gorm:"not null;default:false" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// 角色常量定义
const (
	RoleUser  = 0 // 普通用户
	RoleAdmin = 1 // 管理员
)
