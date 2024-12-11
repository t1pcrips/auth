package auth

import (
	"github.com/t1pcrips/auth/internal/repository"
	"github.com/t1pcrips/auth/internal/service"
)

type AuthServiceImpl struct {
	cacheRepository      repository.CacheRepository
	userRepository       repository.UserRepository
	jwtService           service.JWTService
	secretsTimeRedisLive int64
}

func NewAuthServiceImpl(
	cacheRepository repository.CacheRepository,
	userRepository repository.UserRepository,
	jwtService service.JWTService,
	secretsTimeRedisLive int64,
) service.AuthService {
	return &AuthServiceImpl{
		cacheRepository:      cacheRepository,
		userRepository:       userRepository,
		jwtService:           jwtService,
		secretsTimeRedisLive: secretsTimeRedisLive,
	}
}
