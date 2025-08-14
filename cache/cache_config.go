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

type RedisParam struct {
	Host     string
	Port     string
	Password string
	DB       string
	PoolSize string
}

type RedisClient struct {
	Client *redis.Client
}

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

func NewRedisClient(cfg *RedisParam) (*RedisClient, error) {
	db, _ := strconv.Atoi(cfg.DB)
	poolSize, _ := strconv.Atoi(cfg.PoolSize)
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       db,
		PoolSize: poolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		xlog.Fatal(context.Background(), err, "connection to redis failed")
		return nil, err
	}

	return &RedisClient{Client: rdb}, nil
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}
