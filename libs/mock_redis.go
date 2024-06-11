package libs

import "github.com/gomodule/redigo/redis"

// MockRedisClient is a mock implementation of RedisClient interface.
type MockRedisClient struct {
	Conn redis.Conn
	Err  error
}

// GetConnection returns the mock Redis connection.
func (m *MockRedisClient) GetConnection(host string) (redis.Conn, error) {
	return m.Conn, m.Err
}
