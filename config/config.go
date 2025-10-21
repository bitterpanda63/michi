package config

import "github.com/spf13/viper"

func GetAPIToken() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.michi")
	viper.ReadInConfig()
	return viper.GetString("api_token")
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.michi")
	viper.SafeWriteConfig()
}
