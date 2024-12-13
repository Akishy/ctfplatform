package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/deployer/internal/config"
	"go.uber.org/fx"
	"log"
	"time"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	options, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	client := redis.NewClient(options)

	// Проверим подключение к Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis")
	return client
}

var Module = fx.Provide(NewRedisClient)
