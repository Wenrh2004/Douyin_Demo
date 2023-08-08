package repo

import (
	"Douyin_Demo/config"
	"Douyin_Demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = InitDB()
}

// InitDB == > Initialize mysql database connection
func InitDB() *gorm.DB {
	dsn := config.AppConfig.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("mysql - database connect error  == > " + err.Error())
	}
	// if user model changed then auto migrate to new user model
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("user model has changed == > transfer to new model struct" + err.Error())
	}
	DB = db
	return db
}

// GetDB == > get database connection
func GetDB() *gorm.DB {
	return DB
}
