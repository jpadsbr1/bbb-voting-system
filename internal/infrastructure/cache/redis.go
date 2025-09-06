package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(redis_url string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: redis_url,
		DB:   0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to ping cache: %v\n", err)
	}

	log.Printf("Connecting to Redis...")
	return &RedisClient{client: client}
}

func (r *RedisClient) Close() {
	log.Printf("Closing Redis...")
	r.client.Close()
}

func (r *RedisClient) GetRedisClient() *redis.Client {
	return r.client
}
