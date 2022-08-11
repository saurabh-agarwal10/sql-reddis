package config

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

const RedisDB = 0

var RedisClient *redis.Client

func ConnectRedis() error {
	if RedisClient != nil {
		return nil
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       RedisDB,
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		return err
	}

	log.Println("Redis Successfully Connected.", pong)

	return nil
}
