package env

import (
	"errors"
	"github.com/t1pcrips/auth/internal/config"
	"os"
	"strconv"
)

const (
	accessSecret        = "ACCESS_SECRET_KEY_PATH"
	refreshSecret       = "REFRESH_SECRET_KEY_PATH"
	timeAccessSecret    = "TIME_ACCESS_JWT"
	timeRefreshSecret   = "TIME_REFRESH_JWT"
	timeRedisLiveSecret = "TIME_REDIS_LIVE"
)

type SecretsConfigSearcher struct{}

func NewSecretsConfigSearcher() *SecretsConfigSearcher {
	return &SecretsConfigSearcher{}
}

func (cfg *SecretsConfigSearcher) Get() (*config.SecretsConfig, error) {
	access := os.Getenv(accessSecret)
	if len(access) == 0 {
		return nil, errors.New("access secret not found")
	}

	refresh := os.Getenv(refreshSecret)
	if len(refresh) == 0 {
		return nil, errors.New("refresh secret not found")
	}

	accessTime := os.Getenv(timeAccessSecret)
	if len(access) == 0 {
		return nil, errors.New("access secret not found")
	}

	accessIntTime, err := strconv.Atoi(accessTime)
	if err != nil {
		return nil, errors.New("use integer")
	}

	refreshTime := os.Getenv(timeRefreshSecret)
	if len(refresh) == 0 {
		return nil, errors.New("refresh secret not found")
	}

	refreshIntTime, err := strconv.Atoi(refreshTime)
	if err != nil {
		return nil, errors.New("use integer")
	}

	timeRedisLive := os.Getenv(timeRedisLiveSecret)
	if len(refresh) == 0 {
		return nil, errors.New("timeRedisLive secret not found")
	}

	timeRedisLiveInt, err := strconv.Atoi(timeRedisLive)
	if err != nil {
		return nil, errors.New("use integer")
	}

	return &config.SecretsConfig{
		JWTAccess:      access,
		JWTRefresh:     refresh,
		JWTAccessTime:  int64(accessIntTime),
		JWTRefreshTime: int64(refreshIntTime),
		TimeRedisLive:  int64(timeRedisLiveInt),
	}, nil
}
