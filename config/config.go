package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func initConfig() {
	// Define the config path
	configPath := filepath.Join(os.Getenv("HOME"), ".michi")

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(configPath, 0755); err != nil {
		log.Printf("Error creating config directory: %v", err)
		return
	}

	// Set up Viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.SetDefault("api_token", "")
	viper.SetDefault("mistral_api_token", "")

	// Write the config file
	if err := viper.SafeWriteConfig(); err != nil {
		log.Printf("Error writing config: %v", err)
	}
}

func readConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.michi")
	return viper.ReadInConfig()
}

func GetMistralAPIToken() string {
	if err := readConfig(); err != nil {
		initConfig()
		_ = readConfig()
	}
	return viper.GetString("mistral_api_token")
}

func SetMistralAPIToken(token string) error {
	if err := readConfig(); err != nil {
		initConfig()
	}
	viper.Set("mistral_api_token", token)
	return viper.WriteConfig()
}

func GetClaudeAPIToken() string {
	if err := readConfig(); err != nil {
		initConfig()
		_ = readConfig()
	}
	return viper.GetString("claude_api_token")
}

func SetClaudeAPIToken(token string) error {
	if err := readConfig(); err != nil {
		initConfig()
	}
	viper.Set("claude_api_token", token)
	return viper.WriteConfig()
}
