package redis

import (
	"errors"
	"fmt"
)

var errConfigHostEmpty = errors.New("Config Host must not be empty")
var errConfigPortZero = errors.New("Config Port must not be more than 0")
var errConfigPoolSizeZero = errors.New("Config PoolSize must be more than 0")
var errConfigIdleTimeoutZero = errors.New("Config IdleTimeout must be more than 0")

func errRedisConnection(err error) error {
	return fmt.Errorf("Redis connection error: %v", err)
}
