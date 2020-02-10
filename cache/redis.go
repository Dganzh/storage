package cache

import (
	"github.com/garyburd/redigo/redis"
)

const (
	redisHost = "127.0.0.1:6379"
)

var (
	pool *redis.Pool
)


func newRedisPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			return c, err
		},
		MaxIdle: 50,
		MaxActive: 30,
	}
}

func init() {
	pool = newRedisPool()
}

func RedisPool() *redis.Pool {
	return pool
}
