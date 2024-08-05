package config

import (
	"github.com/spf13/viper"
	"log"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func GetDatabaseURL() string {
	return viper.GetString("database.url")
}

func GetOAuthConfig() map[string]string {
	return map[string]string{
		"clientID":     viper.GetString("oauth.clientID"),
		"clientSecret": viper.GetString("oauth.clientSecret"),
		"redirectURL":  viper.GetString("oauth.redirectURL"),
	}
}
