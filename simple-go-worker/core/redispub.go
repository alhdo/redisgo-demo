package core

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisConfig struct {
	Addr string
}

func InitRedisPool(config RedisConfig) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Addr)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}
