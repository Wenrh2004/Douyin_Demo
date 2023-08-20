package repo

import (
	"Douyin_Demo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := config.AppConfig.DSN
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("mysql - database connect error  == > " + err.Error())
	}
	// if user model changed then auto migrate to new user model
	//err = db.AutoMigrate(&model.User{})
	//if err != nil {
	//	panic("user model has changed == > transfer to new model struct" + err.Error())
	//}

	SetDefault(DB)

}
