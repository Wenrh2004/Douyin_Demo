package manager

import (
	"github.com/spf13/viper"
)

// GetTokenSecretConfig == > get whole configuration properties
func GetYamlConfigByString(content string) string {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file can not be founded")
		}
		panic("config file read error+" + err.Error())
	}
	secretKey := viper.GetString(content)
	return secretKey
}

// GetTokenSecretConfig == > get whole configuration properties
func GetYamlConfigByInt(content string) int {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file can not be founded")
		}
		panic("config file read error+" + err.Error())
	}
	secretKey := viper.GetInt(content)
	return secretKey
}
