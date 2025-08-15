package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// Get
//
// Redis gets value by key.
func (rc *RedisClient) Get(ctx context.Context, key string, target interface{}) error {
	val, err := rc.Client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		return fmt.Errorf("backrooms: get value in redis failed: %w", err)
	}

	if val == "" {
		return nil
	}

	if err := json.Unmarshal([]byte(val), target); err != nil {
		return fmt.Errorf("backrooms: unmarshal value in redis failed: %w", err)
	}
	return nil
}

// Set
//
// Redis sets value by key-value and ttl.
func (rc *RedisClient) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	default:
		var err error
		data, err = json.Marshal(value)
		if err != nil {
			return fmt.Errorf("backrooms: marshal value in redis failed: %w", err)
		}
	}

	return rc.Client.Set(ctx, key, data, ttl).Err()
}

// Delete
//
// Redis delete value by key.
func (rc *RedisClient) Delete(ctx context.Context, key string) error {
	_, err := rc.Client.Del(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("backrooms: delete key in redis failed: %w", err)
	}
	return nil
}
