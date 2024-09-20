package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

// Initialize a new Redis client
func NewRedisClient(host, port, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password, // No password set
		DB:       db,       // Use default DB
	})

	// Test the connection
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis!")
	return rdb
}
