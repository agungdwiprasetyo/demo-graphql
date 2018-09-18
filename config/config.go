package config

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

// Config abstraction
type Config interface {
	LoadReadDB() *sqlx.DB
	LoadWriteDB() *sqlx.DB
	LoadReadRedis() *redis.Client
	LoadWriteRedis() *redis.Client
}

type conf struct {
	databaseConn
	redisConn
}

// New init config
func New() Config {
	return new(conf)
}
