package manager

import (
	"github.com/spf13/viper"
)

// GetTokenConfig == > get whole configuration properties
func GetTokenExpireConfig() int {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file can not be founded")
		}
		panic("config file read error+" + err.Error())
	}
	tokenEXpr := viper.GetInt("token.expire")
	return tokenEXpr
}

// GetTokenSecretConfig == > get whole configuration properties
func GetTokenSecretConfig() string {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file can not be founded")
		}
		panic("config file read error+" + err.Error())
	}
	secretKey := viper.GetString("token.mySecret")
	return secretKey
}
