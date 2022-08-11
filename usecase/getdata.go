package usecase

import (
	"log"

	"sqlredis-server/config"
)

func GetColour(productName string) (string, error) {
	log.Println("Searching for value in redis")
	value, err := config.RedisClient.Get(productName).Result()
	//if value is not found, or in case of error, fallback to db
	if value == "" || err != nil {
		log.Println("Couldn't find the value in Redis, falling back to db")
		value, err = GetColourOfProduct(productName)
		if err != nil {
			log.Println("Error in GetColourOfProduct", err)
			return "", err
		}

		err := SetColour(productName, value)
		if err != nil {
			log.Println("Error in SetColour", err)
		}
	}

	return value, nil
}
