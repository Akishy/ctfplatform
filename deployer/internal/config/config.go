package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	PostgresURL string
	RedisURL    string
	AppPort     string
}

func NewConfig() *Config {
	viper.SetDefault("APP_PORT", "8080")
	viper.SetEnvPrefix("ADMIN")
	viper.AutomaticEnv()

	config := &Config{
		PostgresURL: viper.GetString("POSTGRES_URL"),
		RedisURL:    viper.GetString("REDIS_URL"),
		AppPort:     viper.GetString("APP_PORT"),
	}

	if config.PostgresURL == "" || config.RedisURL == "" {
		log.Fatal("Database configuration is missing")
	}

	return config
}
