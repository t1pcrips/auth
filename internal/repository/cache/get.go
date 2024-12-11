package cache

import (
	"context"
	"github.com/gomodule/redigo/redis"
)

const (
	getCommand = "GET"
)

func (repo *CacheRepositoryImpl) Get(ctx context.Context, key string) ([]byte, error) {
	data, err := redis.Bytes(repo.db.DB().DoContext(ctx, getCommand, repo.redisConfig.CtxTimeout, key))
	if err != nil {
		return nil, err
	}

	return data, nil
}
