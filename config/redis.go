package config

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

type redisConn struct{}

func (r *redisConn) LoadReadRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_READ_HOST"), os.Getenv("REDIS_READ_PORT")),
		Password: os.Getenv("REDIS_READ_PASSWORD"),
	})
}

func (r *redisConn) LoadWriteRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_WRITE_HOST"), os.Getenv("REDIS_WRITE_PORT")),
		Password: os.Getenv("REDIS_WRITE_PASSWORD"),
	})
}
