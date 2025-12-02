package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/online-compiler/backend/configs"
)

var RedisClient *redis.Client
var ctx = context.Background()

// InitRedis initializes Redis connection
func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     configs.AppConfig.RedisURL,
		Password: configs.AppConfig.RedisPassword,
		DB:       configs.AppConfig.RedisDB,
	})

	// Test connection
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return nil
}

// CheckRateLimit checks if the IP has exceeded rate limit
func CheckRateLimit(ip string) (bool, error) {
	// If Redis is not connected, allow all requests
	if RedisClient == nil {
		return true, nil
	}

	key := fmt.Sprintf("rate:execute:%s", ip)

	count, err := RedisClient.Get(ctx, key).Int()
	if err == redis.Nil {
		// First request
		err = RedisClient.Set(ctx, key, 1, time.Duration(configs.AppConfig.RateLimitWindow)*time.Second).Err()
		return true, err
	} else if err != nil {
		// If Redis error, allow request (fail open)
		return true, nil
	}

	if count >= configs.AppConfig.RateLimitRequests {
		return false, nil
	}

	// Increment counter
	err = RedisClient.Incr(ctx, key).Err()
	if err != nil {
		// If Redis error, allow request (fail open)
		return true, nil
	}
	return true, nil
}

// CacheResult caches execution result
func CacheResult(codeHash string, result interface{}) error {
	// If Redis is not connected, skip caching
	if RedisClient == nil {
		return nil
	}

	key := fmt.Sprintf("cache:result:%s", codeHash)
	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return RedisClient.Set(ctx, key, data, time.Hour).Err()
}

// GetCachedResult retrieves cached execution result
func GetCachedResult(codeHash string) ([]byte, error) {
	// If Redis is not connected, return cache miss
	if RedisClient == nil {
		return nil, redis.Nil
	}

	key := fmt.Sprintf("cache:result:%s", codeHash)
	return RedisClient.Get(ctx, key).Bytes()
}

// GetRedisClient returns the Redis client instance
func GetRedisClient() *redis.Client {
	return RedisClient
}
