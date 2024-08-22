package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()


type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() (RedisClient, error) {

	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	socket := fmt.Sprintf(`%v:%v`, redisAddress, redisPort)
	client := redis.NewClient(&redis.Options{
		Addr:     socket,
		Password: redisPassword,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return RedisClient{}, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	return  RedisClient{
		client: client,
	}, nil

}

func (rc *RedisClient) Set(key string, value string, expiration time.Duration) error {
	err := rc.client.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("error setting value in Redis: %w", err)
	}
	return nil
}

func (rc *RedisClient) Get(key string) (string, error) {
	value, err := rc.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("error getting value from Redis: %w", err)
	}
	return value, nil
}
