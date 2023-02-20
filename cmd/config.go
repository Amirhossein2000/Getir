package main

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server  ServerConfig
	Redis   RedisConfig
	MongoDB MongoConfig
}

type ServerConfig struct {
	Address string
}

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

type MongoConfig struct {
	Address    string
	Database   string
	Collection string
}

func GetConfig() (*Config, error) {
	strRedisDB := os.Getenv("REDIS_DB")
	redisDB := 0
	if strRedisDB != "" {
		var err error
		redisDB, err = strconv.Atoi(strRedisDB)
		if err != nil {
			return nil, fmt.Errorf("REDIS_DB should be a number, err: %w", err.Error())
		}
	}

	c := &Config{
		Server: ServerConfig{
			Address: os.Getenv("SERVER_ADDRESS"),
		},
		Redis: RedisConfig{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDB,
		},
		MongoDB: MongoConfig{
			Address:    os.Getenv("MONGO_ADDRESS"),
			Database:   os.Getenv("MONGO_DATABASE"),
			Collection: os.Getenv("MONGO_COLLECTION"),
		},
	}

	return c, nil
}
