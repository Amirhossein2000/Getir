package cache

import (
	"Getir/internal/ports"
	"github.com/go-redis/redis"
	"time"
)

type repo struct {
	client *redis.Client
}

func NewRepo(client *redis.Client) ports.Cache {
	return &repo{
		client: client,
	}
}

func (r *repo) Set(key, value string) error {
	err := r.client.Set(key, value, time.Second*5).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Get(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
