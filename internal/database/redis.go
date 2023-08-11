package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisPool struct {
	rdb    *redis.Client
	rdbCtx context.Context
}

func NewRedisPool(options *redis.Options) (*RedisPool, error) {
	rdbCtx := context.Background()

	rdb := redis.NewClient(options)

	err := rdb.Ping(rdbCtx).Err()
	if err != nil {
		return nil, err
	}

	return &RedisPool{
		rdb, rdbCtx,
	}, nil
}
