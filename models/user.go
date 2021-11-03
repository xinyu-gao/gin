package models


import (
	"time"
)

type User struct {
	ID         int       `gorm:"column:id;primary_key"`      // 用户id
	Name       string    `gorm:"column:name;NOT NULL"`       // 用户名
	Password   string    `gorm:"column:password;NOT NULL"`   // 密码
	AvatarUrl  string    `gorm:"column:avatar_url"`          // 头像URL
	Phone      string    `gorm:"column:phone"`               // 手机号码
	Email      string    `gorm:"column:email"`               // 邮箱
	Status     string    `gorm:"column:status;NOT NULL"`     // 用户状态 1：可用；0：冻结；
	SourceFrom string    `gorm:"column:source_from"`         // 用户来源
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL"` // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at"`          // 更新时间
	DeletedAt  time.Time `gorm:"column:deleted_at"`          // 删除时间
}

func (m *User) TableName() string {
	return "users"
}