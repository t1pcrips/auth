package app

import (
	"context"
	"github.com/t1pcrips/auth/internal/api/access"
	"github.com/t1pcrips/auth/internal/api/auth"
	"github.com/t1pcrips/auth/internal/api/user"
	"github.com/t1pcrips/auth/internal/config"
	"github.com/t1pcrips/auth/internal/config/env"
	"github.com/t1pcrips/auth/internal/repository"
	accessRepository "github.com/t1pcrips/auth/internal/repository/access"
	cacheRepository "github.com/t1pcrips/auth/internal/repository/cache"
	userRepository "github.com/t1pcrips/auth/internal/repository/user"
	"github.com/t1pcrips/auth/internal/service"
	accessService "github.com/t1pcrips/auth/internal/service/access"
	authService "github.com/t1pcrips/auth/internal/service/auth"
	jwtService "github.com/t1pcrips/auth/internal/service/jwt"
	userService "github.com/t1pcrips/auth/internal/service/user"
	"github.com/t1pcrips/platform-pkg/pkg/closer"
	"github.com/t1pcrips/platform-pkg/pkg/database"
	"github.com/t1pcrips/platform-pkg/pkg/database/postgres"
	"github.com/t1pcrips/platform-pkg/pkg/database/transaction"
	"github.com/t1pcrips/platform-pkg/pkg/memory_database"
	"github.com/t1pcrips/platform-pkg/pkg/memory_database/redis"

	"log"
)

type serviceProvider struct {
	pgConfig      *config.PgConfig
	redisConfig   *config.RedisConfig
	grpcConfig    *config.GRPCConfig
	httpConfig    *config.HTTPConfig
	swaggerConfig *config.SwagerConfig
	secretsConfig *config.SecretsConfig

	dbClient       database.Client
	memoryDBClient memory_database.Client
	txManeger      database.TxManeger

	userRepository   repository.UserRepository
	cacheRepository  repository.CacheRepository
	accessRepository repository.AccessRepository

	userService   service.UserService
	jwtService    service.JWTService
	authService   service.AuthService
	accessService service.AccessService

	userApiImpl   *user.UserApiImpl
	accessApiImpl *access.AccessApiImpl
	authApiImpl   *auth.AuthApiImpl
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PgConfig() *config.PgConfig {
	if s.pgConfig == nil {
		cfgSearcher := env.NewPgConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() *config.GRPCConfig {
	if s.grpcConfig == nil {
		cfgSearcher := env.NewGRPCConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() *config.HTTPConfig {
	if s.httpConfig == nil {
		cfgSearcher := env.NewHTTPCfgSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) RedisConfig() *config.RedisConfig {
	if s.redisConfig == nil {
		cfgSearcher := env.NewRedisConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get redis config: %s", err.Error())
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) SwaggerConfig() *config.SwagerConfig {
	if s.swaggerConfig == nil {
		cfgSearcher := env.NewSwaggerConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) SecretsConfig() *config.SecretsConfig {
	if s.secretsConfig == nil {
		cfgSearcher := env.NewSecretsConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get secrets config: %s", err.Error())
		}

		s.secretsConfig = cfg
	}

	return s.secretsConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) database.Client {
	if s.dbClient == nil {
		dbc, err := postgres.New(ctx, s.PgConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create DBClient: %s", err.Error())
		}

		closer.Add(dbc.Close)

		err = dbc.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping postgres database: %s", err.Error())
		}

		s.dbClient = dbc
	}

	return s.dbClient
}

func (s *serviceProvider) MemoryDBClient(ctx context.Context) memory_database.Client {
	if s.memoryDBClient == nil {
		client := redis.NewClientRs(ctx, "tcp", s.RedisConfig().Address(), s.RedisConfig().MaxIdle, s.RedisConfig().IdleTimeout)

		closer.Add(client.Close)

		err := client.DB().Ping(ctx, s.RedisConfig().CtxTimeout)
		if err != nil {
			log.Fatalf("failed to ping redis memory database: %s", err.Error())
		}

		s.memoryDBClient = client
	}

	return s.memoryDBClient
}

func (s *serviceProvider) TxManeger(ctx context.Context) database.TxManeger {
	if s.txManeger == nil {
		s.txManeger = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManeger
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepositoryImpl(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) CacheRepository(ctx context.Context) repository.CacheRepository {
	if s.cacheRepository == nil {
		s.cacheRepository = cacheRepository.NewCacheRepositoryImpl(s.MemoryDBClient(ctx), s.RedisConfig())
	}

	return s.cacheRepository
}

func (s *serviceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewAccessRepositoryImpl(s.DBClient(ctx))
	}

	return s.accessRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewUserServiceImpl(s.UserRepository(ctx), s.TxManeger(ctx))
	}

	return s.userService
}

func (s *serviceProvider) JWTService(ctx context.Context) service.JWTService {
	if s.jwtService == nil {
		s.jwtService = jwtService.NewJWTServiceImpl(s.SecretsConfig(), s.CacheRepository(ctx))
	}

	return s.jwtService
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewAuthServiceImpl(
			s.CacheRepository(ctx),
			s.UserRepository(ctx),
			s.JWTService(ctx),
			s.SecretsConfig().TimeRedisLive,
		)
	}

	return s.authService
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewAccessServiceImpl(s.JWTService(ctx), s.AccessRepository(ctx))
	}

	return s.accessService
}

func (s *serviceProvider) UserApiImpl(ctx context.Context) *user.UserApiImpl {
	if s.userApiImpl == nil {
		s.userApiImpl = user.NewUserApiImpl(s.UserService(ctx))
	}

	return s.userApiImpl
}

func (s *serviceProvider) AuthApiImpl(ctx context.Context) *auth.AuthApiImpl {
	if s.authApiImpl == nil {
		s.authApiImpl = auth.NewAuthApiImpl(s.AuthService(ctx))
	}

	return s.authApiImpl
}

func (s *serviceProvider) AccessApiImpl(ctx context.Context) *access.AccessApiImpl {
	if s.accessApiImpl == nil {
		s.accessApiImpl = access.NewAccessApiImpl(s.AccessService(ctx))
	}

	return s.accessApiImpl
}
