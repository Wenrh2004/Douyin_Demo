package model

import "github.com/jinzhu/gorm"

// 用户数据	users
type User struct {
	gorm.Model
	Username string `gorm:"varchar(20);not null"`
	Password string `gorm:"size:255;not null"`
}
