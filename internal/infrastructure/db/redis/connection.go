package redis

import (
	"github.com/redis/go-redis/v9"
)

func NewConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // TODO add password
		DB:       0,
	})
}
