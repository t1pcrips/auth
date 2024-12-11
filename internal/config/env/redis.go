package env

import (
	"errors"
	"github.com/t1pcrips/auth/internal/config"
	"os"
	"strconv"
	"time"
)

const (
	hostRD      = "REDIS_HOST"
	portRD      = "REDIS_PORT"
	maxIdle     = "MAX_IDLE"
	idleTimeout = "IDLE_TIMEOUT_SEC"
	ctxTimeout  = "CONTEXT_TIMEOUT_SEC"
)

type RedisConfigSearcher struct{}

func NewRedisConfigSearcher() *RedisConfigSearcher {
	return &RedisConfigSearcher{}
}

func (cfg *RedisConfigSearcher) Get() (*config.RedisConfig, error) {
	host := os.Getenv(hostRD)
	if len(host) == 0 {
		return nil, errors.New("redis host not found")
	}

	portString := os.Getenv(portRD)
	if len(portString) == 0 {
		return nil, errors.New("redis port not found")
	}

	_, err := strconv.Atoi(portString)
	if err != nil {
		return nil, errors.New("incorrect port, use integer port")
	}

	maxIdleString := os.Getenv(maxIdle)
	maxIdleInt, err := strconv.Atoi(maxIdleString)
	if err != nil {
		return nil, errors.New("incorrect maxIdle, use integer maxIdle")
	}

	idleTimeoutString := os.Getenv(idleTimeout)
	idleTimeoutInt, err := strconv.Atoi(idleTimeoutString)
	if err != nil {
		return nil, errors.New("incorrect IdleTimeout, use integer IdleTimeout")
	}

	ctxTimeoutString := os.Getenv(ctxTimeout)
	ctxTimeoutInt, err := strconv.Atoi(ctxTimeoutString)
	if err != nil {
		return nil, errors.New("incorret ctx timeout, use integer ctx timeout")
	}

	return &config.RedisConfig{
		Host:        host,
		Port:        portString,
		MaxIdle:     maxIdleInt,
		IdleTimeout: time.Duration(idleTimeoutInt) * time.Second,
		CtxTimeout:  time.Duration(ctxTimeoutInt) * time.Second,
	}, nil
}
