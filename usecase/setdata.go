package usecase

import (
	"sqlredis-server/config"
)

func SetColour(productName string, val string) error {
	return config.RedisClient.Set(productName, val, 0).Err()
}
