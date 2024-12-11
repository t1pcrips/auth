package cache

import (
	"github.com/t1pcrips/auth/internal/config"
	"github.com/t1pcrips/platform-pkg/pkg/memory_database"

	"github.com/t1pcrips/auth/internal/repository"
)

type CacheRepositoryImpl struct {
	db          memory_database.Client
	redisConfig *config.RedisConfig
}

func NewCacheRepositoryImpl(db memory_database.Client, redisConfig *config.RedisConfig) repository.CacheRepository {
	return &CacheRepositoryImpl{
		db:          db,
		redisConfig: redisConfig,
	}
}
