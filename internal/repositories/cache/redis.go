package cache

import (
	"github.com/go-redis/redis"
)

type RedisParams struct {
	Address  string
	Password string
	DB       int
}

func NewRedis(params *RedisParams) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     params.Address,
		Password: params.Password,
		DB:       params.DB,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
