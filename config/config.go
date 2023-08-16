package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type AppConfigSchema struct {
	// DataBase
	Database       MySQL `mapstructure:"mysql"`
	DSN            string
	AWS            AWS    `mapstructure:"aws"`
	CONSUL_ADDRESS string `mapstructure:"consulAddress"`
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

type AWS struct {
	AccessKey         string `mapstructure:"accessKey"`
	Secret            string `mapstructure:"secret"`
	Region            string `mapstructure:"region"`
	BucketName        string `mapstructure:"bucketName"`
	LambdaFunctionUrl string `mapstructure:"lambdaFunctionUrl"`
}

const (
	FeedServiceName = "feed-service"
	FeedServicePort = ":9010"

	PublishServiceName = "publish-service"
	PublishServicePort = ":9011"
)

// AppConfig == > global variable
var AppConfig = AppConfigSchema{}

// init ==> reads a configuration file, unmarshals the configuration data, and sets the AppConfig.DSN variable to a connection string that specifies the database connection parameters
func init() {
	readConfig()
	unmarshallConfig()
	AppConfig.DSN = AppConfig.Database.UserName + ":" + AppConfig.Database.Password + "@tcp(" + AppConfig.Database.Host + ":" + AppConfig.Database.Port + ")/" + AppConfig.Database.Database + "?charset=utf8&parseTime=True"
}

// readConfig == > get whole configuration properties
func readConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic("get current path error+" + err.Error())
	}
	// trim the path to the project root
	rootDir := findProjectRoot(path)
	if rootDir == "" {
		panic("can not find project root directory")
	}

	// set the configuration file path
	configFilePath := filepath.Join(rootDir, "application.yaml")

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFilePath)

	setDefaultConfig()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file can not be founded")
		}
		panic("config file read error+" + err.Error())
	}
}

func findProjectRoot(cwd string) string {
	// Traverse up the directory tree until the project root directory is found
	for {
		if cwd == "" || cwd == "/" {
			return ""
		}

		// Check if the current directory contains the "Douyin_Demo" directory
		if filepath.Base(cwd) == "Douyin_Demo" {
			return cwd
		}

		// Move up one level in the directory tree
		cwd = filepath.Dir(cwd)
	}
}

// unmarshallConfig == > unmarshall configuration properties
func unmarshallConfig() {
	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic("config file Unmarshal error+" + err.Error())
	}

	// fmt.Println("AppConfig", AppConfig)
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

	// aws
	viper.SetDefault("aws.accessKey", "accessKey")
	viper.SetDefault("aws.secret", "secret")
	viper.SetDefault("aws.region", "region")
	viper.SetDefault("aws.bucketName", "bucketName")
	viper.SetDefault("aws.lambdaFunctionUrl", "lambdaFunctionUrl")

	// consul
	viper.SetDefault("CONSUL_ADDRESS", "http://localhost:8500")
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
