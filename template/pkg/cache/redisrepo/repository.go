package redisrepo

import (
	"{{bootstrap_template}}/pkg/cache"
	"context"
	"encoding/json"
	"fmt"

	"time"
)

type RedisRepository struct {
	cache cache.RedisClient
}

var ctx = context.Background()

type IRedisRepository interface {
	Get(key string, result interface{}) error
	Set(key string, value interface{}, expiration time.Duration) error
}

func NewRedisRepository(cache cache.RedisClient) *RedisRepository {
	return &RedisRepository{
		cache: cache,
	}
}

func (rr *RedisRepository) Get(key string, result interface{}) error {
	value, err := rr.cache.Get(key)
	if err != nil {
		return fmt.Errorf("error getting value from Redis: %w", err)
	}

	if value == "" {
		return nil // the key not exist
	}

	err = json.Unmarshal([]byte(value), &result)
	if err != nil {
		return fmt.Errorf("error unmarshaling value from Redis: %w", err)
	}

	return nil
}

func (rr *RedisRepository) Set(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error marshaling value to JSON: %w", err)
	}

	err = rr.cache.Set(key, string(data), expiration)
	if err != nil {
		return fmt.Errorf("error setting value in Redis: %w", err)
	}

	return nil
}
