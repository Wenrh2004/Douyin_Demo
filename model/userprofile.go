package model

import "gorm.io/gorm"

// UserProfile 用户profile数据
type UserProfile struct {
	gorm.Model
	UserId          int64  `json:"user_id" gorm:"primaryKey;not null"`
	Name            string `json:"name" gorm:"varchar(20);not null"`
	Avatar          string `json:"Avatar" gorm:"varchar(255);"`
	BackgroundImage string `json:"background_image" gorm:"varchar(255);"`
	Signature       string `json:"signature" gorm:"varchar(255);"`
}
