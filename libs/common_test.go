package libs

import (
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

// MockConn is a mock implementation of redis.Conn interface.
type MockConn struct{}

func (m *MockConn) Close() error { return nil }
func (m *MockConn) Err() error   { return nil }
func (m *MockConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	if commandName == "GET" && args[0] == "existing-key" {
		return "value", nil
	}
	if commandName == "GET" && args[0] == "missing-key" {
		return nil, redis.ErrNil
	}
	return nil, nil
}
func (m *MockConn) Send(commandName string, args ...interface{}) error {
	return nil
}
func (m *MockConn) Flush() error                            { return nil }
func (m *MockConn) Receive() (reply interface{}, err error) { return nil, nil }
func (m *MockConn) ReceiveWithTimeout(timeout time.Duration) (reply interface{}, err error) {
	return nil, nil
}

func TestCheckKeyInRedis(t *testing.T) {
	mockConn := &MockConn{}

	tests := []struct {
		name    string
		key     string
		want    string
		wantErr bool
	}{
		{
			name:    "Key exists",
			key:     "existing-key",
			want:    "something",
			wantErr: false,
		},
		{
			name:    "Key does not exist",
			key:     "missing-key",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckKeyInRedis(mockConn, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckKeyInRedis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckKeyInRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}
