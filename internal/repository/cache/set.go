package cache

import (
	"context"
	"encoding/json"
)

const (
	setCommand = "SETEX"
)

func (repo *CacheRepositoryImpl) Set(ctx context.Context, key string, info interface{}, ttl int64) error {
	infoJSON, err := json.Marshal(info)
	if err != nil {
		return err
	}

	_, err = repo.db.DB().DoContext(ctx, setCommand, repo.redisConfig.CtxTimeout, key, ttl, infoJSON)
	if err != nil {
		return err
	}

	return nil
}
