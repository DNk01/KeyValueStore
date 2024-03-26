package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type AppConfig struct {
	Port       string
	SwaggerUI  string
	DefaultTTL time.Duration
}

func LoadConfig() AppConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.SetDefault("port", ":8080")
	viper.SetDefault("swaggerUI", "http://localhost:8080/swagger/doc.json")
	viper.SetDefault("defaultTTL", 60)

	return AppConfig{
		Port:       viper.GetString("port"),
		SwaggerUI:  viper.GetString("swaggerUI"),
		DefaultTTL: time.Duration(viper.GetInt("defaultTTL")) * time.Second,
	}
}
