package db

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// InitRedis initializes the Redis connection
func InitRedis(addr string, password string, db int) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis successfully!")
}
