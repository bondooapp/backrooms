package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func (r *RedisClient) Get(ctx context.Context, key string, target interface{}) error {
	val, err := r.Client.Get(ctx, key).Result()
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

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
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

	return r.Client.Set(ctx, key, data, ttl).Err()
}

func (r *RedisClient) Delete(ctx context.Context, key string) error {
	_, err := r.Client.Del(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("backrooms: delete key in redis failed: %w", err)
	}
	return nil
}
