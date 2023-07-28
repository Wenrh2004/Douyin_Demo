/*
 * Copyright (c) 2023.
 * Project: Douyin_Demo
 * File: db_init.go
 * Last date: 2023/7/28 下午6:52
 * Developer: KingYen
 *
 * Created by KingYen on 2023/7/28 18:52:9.
 */

package common

import (
	"Douyin_Demo/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "Douyin_Demo"
	username := "root"
	password := "Wenrh240004"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	//迁移
	db.AutoMigrate(&model.User{})

	DB = db

	return db

}

// 获取数据库
func GetDB() *gorm.DB {
	return DB
}
