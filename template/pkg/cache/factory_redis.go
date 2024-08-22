package cache

import (
	"errors"
)

var (
	errInvalidRedisInstance = errors.New("invalid redis instance")
)

const (
	InstanceRedis int = iota
)

func NewRedisFactory(instance int) (RedisClient, error){
	switch instance {
	case InstanceRedis:
		return NewRedisClient()
	default:
		return RedisClient{}, errInvalidRedisInstance
	}
}
