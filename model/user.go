/*
 * Copyright (c) 2023.
 * Project: Douyin_Demo
 * File: user.go
 * Last date: 2023/7/28 下午6:52
 * Developer: KingYen
 *
 * Created by KingYen on 2023/7/28 18:52:9.
 */

package model

import "github.com/jinzhu/gorm"

// 用户数据	users
type User struct {
	gorm.Model
	Username string `gorm:"varchar(20);not null"`
	Password string `gorm:"size:255;not null"`
}
