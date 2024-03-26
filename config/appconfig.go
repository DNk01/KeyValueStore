package main

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	Port      string
	SwaggerUI string
}

func LoadConfig() AppConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.SetDefault("port", "8080")
	viper.SetDefault("swaggerUI", "http://localhost:8080/swagger/doc.json")

	return AppConfig{
		Port:      viper.GetString("port"),
		SwaggerUI: viper.GetString("swaggerUI"),
	}
}
