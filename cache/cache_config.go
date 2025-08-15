package cache

import (
	"context"
	"fmt"
	"github.com/bondooapp/backrooms/util"
	"github.com/bondooapp/backrooms/util/xlog"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

// RedisParam
//
// Redis param.
type RedisParam struct {
	Host     string
	Port     string
	Password string
	DB       string
	PoolSize string
}

// RedisClient
//
// Redis client.
type RedisClient struct {
	Client *redis.Client
}

// LoadRedisParam
//
// Load redis param.
func LoadRedisParam() (*RedisParam, error) {
	_ = godotenv.Load()
	param := &RedisParam{
		Host:     util.GetEnv("REDIS_HOST", "localhost"),
		Port:     util.GetEnv("REDIS_PORT", "6379"),
		Password: util.GetEnv("REDIS_PASSWORD", ""),
		DB:       util.GetEnv("REDIS_DB", "0"),
		PoolSize: util.GetEnv("REDIS_POOL_SIZE", "100"),
	}
	return param, nil
}

// NewRedisClient
//
// New redis client.
func NewRedisClient(rp *RedisParam) (*RedisClient, error) {
	db, _ := strconv.Atoi(rp.DB)
	poolSize, _ := strconv.Atoi(rp.PoolSize)
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rp.Host, rp.Port),
		Password: rp.Password,
		DB:       db,
		PoolSize: poolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		xlog.Fatal(context.Background(), err, "backrooms: failed to connect redis")
		return nil, err
	}

	return &RedisClient{Client: rdb}, nil
}

// Close
//
// Redis close client.
func (rc *RedisClient) Close() error {
	return rc.Client.Close()
}
