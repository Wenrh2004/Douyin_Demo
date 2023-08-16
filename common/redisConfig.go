package common

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var redisClient *redis.Client

// Redis ==> redis configuration
type Redis struct {
	host     string
	post     int
	userName string
	passWord string
	dataBase int
}

func getRedisConfig() *Redis {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file can not be founded")
		}
		panic("config file read error+" + err.Error())
	}
	return &Redis{
		host:     viper.GetString("redis.host"),
		post:     viper.GetInt("redis.port"),
		userName: viper.GetString("redis.username"),
		passWord: viper.GetString("redis.password"),
		dataBase: viper.GetInt("redis.database"),
	}
}

// InitRedis ==> Initialize redis database connection.
func InitRedis() *redis.Client {
	var config = getRedisConfig()
	dsn := fmt.Sprintf("redis://%s:%s@%s:%d/%d", config.userName, config.passWord, config.host, config.post, config.dataBase)
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		panic("redis - database connect error  == > " + err.Error())
	}

	rdb := redis.NewClient(opt)
	redisClient = rdb
	return rdb
}
