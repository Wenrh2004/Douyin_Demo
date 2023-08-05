package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfigSchema struct {
	// DataBase
	Database MySQL `mapstructure:"mysql"`
	DSN      string
}

// MySQL  == > mysql-database Yaml map to JSON
type MySQL struct {
	Password   string `mapstructure:"passWord"`
	UserName   string `mapstructure:"userName"`
	Host       string `mapstructure:"host"`
	DriverName string `mapstructure:"driverName"`
	Port       string `mapstructure:"port"`
	Database   string `mapstructure:"dataBase"`
}

// AppConfig == > global variable
var AppConfig = AppConfigSchema{}

func init() {
	readConfig()
	unmarshallConfig()
	AppConfig.DSN = AppConfig.Database.UserName + ":" + AppConfig.Database.Password + "@tcp(" + AppConfig.Database.Host + ":" + AppConfig.Database.Port + ")/" + AppConfig.Database.Database + "?charset=utf8&parseTime=True"
}

// readConfig == > get whole configuration properties
func readConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("application.yaml")

	setDefaultConfig()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file can not be founded")
		}
		panic("config file read error+" + err.Error())
	}
}

// unmarshallConfig == > unmarshall configuration properties
func unmarshallConfig() {
	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic("config file Unmarshal error+" + err.Error())
	}

	fmt.Println("AppConfig", AppConfig)
}

// setDefaultConfig == > set default configuration properties
func setDefaultConfig() {
	// mysql
	viper.SetDefault("mysql.username", "root")
	viper.SetDefault("mysql.password", "root")
	viper.SetDefault("mysql.host", "localhost")
	viper.SetDefault("mysql.port", "3306")
	viper.SetDefault("mysql.dataBase", "test")
	viper.SetDefault("mysql.driverName", "mysql")

}

func main() {
	// read config from file
	readConfig()
	unmarshallConfig()
	// get config
	fmt.Println("DSN", AppConfig.DSN)
	fmt.Println("AppConfig", AppConfig)
	// database.port
	fmt.Println("database.port", AppConfig.Database.Port)
}
