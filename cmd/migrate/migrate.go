package main

import (
	"Douyin_Demo/config"
	"Douyin_Demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// use gorm to migrate database
func main() {
	db, err := gorm.Open(mysql.Open(config.AppConfig.DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("mysql - database connect error  == > " + err.Error())
	}

	// migrate user model
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("user model has changed == > transfer to new model struct" + err.Error())
	}

	// migrate publish model
	err = db.AutoMigrate(&model.Publish{})
	if err != nil {
		panic("publish model has changed == > transfer to new model struct" + err.Error())
	}

	// migrate userprofile model
	err = db.AutoMigrate(&model.UserProfile{})
	if err != nil {
		panic("userprofile model has changed == > transfer to new model struct" + err.Error())
	}

}
