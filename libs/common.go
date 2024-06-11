package libs

import "github.com/gomodule/redigo/redis"

func CheckKeyInRedis(redisConn redis.Conn, key string) (string, error) {
	// Simulate the actual interaction with Redis
	_, err := redisConn.Do("GET", key)
	if err != nil {
		return "", err
	}
	return "something", nil
}
