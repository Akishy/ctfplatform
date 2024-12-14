package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/pkg/cache"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/pkg/db"
)

type Config struct {
	db.PostgresConfig
	cache.RedisConfig
	SecretKey string `env:"SECRET_KEY"`
	AppPort   int    `env:"APP_PORT"`
}

func New() (*Config, error) {
	cfg := Config{}
	err := cleanenv.ReadConfig("./config.env", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
