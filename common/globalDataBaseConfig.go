package common

import (
	"Douyin_Demo/model"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DataBase == > mysql-database Yaml map to JSON
type DataBase struct {
	passWord   string
	userName   string
	host       string
	driverName string
	port       int
	dataBase   string
}

// GetProjectConfig == > get whole configuration properties
func getProjectConfig() *DataBase {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file can not be founded")
		}
		panic("config file read error+" + err.Error())
	}
	return &DataBase{
		passWord:   viper.GetString("mysql.password"),
		userName:   viper.GetString("mysql.username"),
		host:       viper.GetString("mysql.host"),
		driverName: viper.GetString("mysql.driverName"),
		port:       viper.GetInt("mysql.port"),
		dataBase:   viper.GetString("mysql.dataBase"),
	}
}

var DB *gorm.DB

// InitDB == > Initialize mysql database connection
func InitDB() *gorm.DB {
	var config = getProjectConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		config.userName, config.passWord, config.host, config.port, config.dataBase)
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

func GetDB() *gorm.DB {
	return DB
}
