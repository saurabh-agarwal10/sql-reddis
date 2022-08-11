package main

import (
	"log"
	"sqlredis-server/config"
	"sqlredis-server/usecase"

	"github.com/joho/godotenv"
)

func main() {
	// load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded.")
	}

	log.Printf(".env file loaded.")
	
	// mysql connection
	err = config.ConnectMysql()
	if err != nil {
		log.Fatal("Unable to connect mysql ", err)
	}

	defer func() {
		log.Println("Closing MYSQL Connection")
		_ = config.MysqlConnection.Close()
	}()

	// redis connection
	err = config.ConnectRedis()
	if err != nil {
		log.Fatal("Unable to connect to redis ", err)
	}

	defer func() {
		log.Println("Closing Redis Connection")
		_ = config.RedisClient.Close()
	}()

	//get product detail
	productName := "Apple"
	val, err := usecase.GetColour(productName)
	if err != nil {
		log.Println("Error in GetColour", err)
		return
	}

	log.Println("Value is:", val)

}
