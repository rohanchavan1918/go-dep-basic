package libs

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// GetRedisConnection returns the redis connection
// func GetRedisConnection(host string) (redis.Conn, error) {

// 	connectTimeout := redis.DialConnectTimeout(10 * time.Second)
// 	readTimeout := redis.DialReadTimeout(5 * time.Second)
// 	writeTimeout := redis.DialWriteTimeout(5 * time.Second)

// 	conn, err := redis.Dial("tcp", host, connectTimeout, readTimeout, writeTimeout)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return conn, nil

// }

// RedisClientImpl is the implementation of RedisClient interface.
type RedisClientImpl struct{}

// GetConnection returns the Redis connection.
func (r *RedisClientImpl) GetConnection(host string) (redis.Conn, error) {
	connectTimeout := redis.DialConnectTimeout(10 * time.Second)
	readTimeout := redis.DialReadTimeout(5 * time.Second)
	writeTimeout := redis.DialWriteTimeout(5 * time.Second)

	conn, err := redis.Dial("tcp", host, connectTimeout, readTimeout, writeTimeout)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
