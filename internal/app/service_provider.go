package app

import (
	"context"
	"github.com/t1pcrips/auth/internal/api/user"
	"github.com/t1pcrips/auth/internal/config"
	"github.com/t1pcrips/auth/internal/config/env"
	"github.com/t1pcrips/auth/internal/repository"
	userRepository "github.com/t1pcrips/auth/internal/repository/user"
	"github.com/t1pcrips/auth/internal/service"
	userService "github.com/t1pcrips/auth/internal/service/user"
	"github.com/t1pcrips/platform-pkg/pkg/database"
	"github.com/t1pcrips/platform-pkg/pkg/database/postgres"
	"github.com/t1pcrips/platform-pkg/pkg/database/transaction"
	"log"
)

type serviceProvider struct {
	pgConfig      *config.PgConfig
	redisConfig   *config.RedisConfig
	grpcConfig    *config.GRPCConfig
	httpConfig    *config.HTTPConfig
	swaggerConfig *config.SwagerConfig

	dbClient  database.Client
	txManeger database.TxManeger

	userRepository repository.UserRepository

	userService service.UserService
	userApiImpl *user.UserApiImpl
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

func (s *serviceProvider) DBClient(ctx context.Context) database.Client {
	if s.dbClient == nil {
		dbc, err := postgres.New(ctx, s.PgConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create DBClient: %s", err.Error())
		}

		err = dbc.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping postgres database: %s", err.Error())
		}

		s.dbClient = dbc
	}

	return s.dbClient
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

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewUserServiceImpl(s.UserRepository(ctx), s.TxManeger(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserApiImpl(ctx context.Context) *user.UserApiImpl {
	if s.userApiImpl == nil {
		s.userApiImpl = user.NewUserApiImpl(s.UserService(ctx))
	}

	return s.userApiImpl
}
