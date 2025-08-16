package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/bondooapp/backrooms/util"
	"github.com/bondooapp/backrooms/util/xlog"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
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
func NewRedisClient(ctx context.Context, rp *RedisParam) (*RedisClient, error) {
	// Get redis db index.
	db, err := strconv.Atoi(rp.DB)
	if err != nil {
		xlog.Fatal(ctx, err, "backrooms: failed to get redis db index")
		return nil, err
	}

	// get redis pool size.
	poolSize, err := strconv.Atoi(rp.PoolSize)
	if err != nil {
		xlog.Fatal(ctx, err, "backrooms: failed to get redis pool size")
		return nil, err
	}

	// Create new redis client.
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rp.Host, rp.Port),
		Password: rp.Password,
		DB:       db,
		PoolSize: poolSize,
	})

	// Set redis connection timeout.
	connCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Test connection.
	if _, err := rdb.Ping(connCtx).Result(); err != nil {
		_ = rdb.Close()
		xlog.Fatal(ctx, err, "backrooms: failed to connect redis")
		return nil, err
	}

	return &RedisClient{Client: rdb}, nil
}
