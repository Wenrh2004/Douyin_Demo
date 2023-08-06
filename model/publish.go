package model

import "gorm.io/gorm"

// Publish 视频数据
type Publish struct {
	gorm.Model
	VideoId  int64  `json:"video_id" gorm:"primary key;comment:视频id;not null"`
	UserId   int64  `json:"user_id" gorm:"not null"`
	Title    string `json:"title" gorm:"varchar(20);not null"`
	PlayUrl  string `json:"play_url" gorm:"varchar(255);not null"`
	CoverUrl string `json:"cover_url" gorm:"varchar(255);not null"`
}
