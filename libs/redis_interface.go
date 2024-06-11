package libs

import "github.com/gomodule/redigo/redis"

// RedisClient defines the methods for interacting with Redis.
type RedisClient interface {
	GetConnection(host string) (redis.Conn, error)
}
