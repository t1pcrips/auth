package app

import (
	"auth/internal/client/database"
	"auth/internal/client/database/postgres"
	"auth/internal/client/database/transaction"
	"auth/internal/config"
	"auth/internal/config/env"
	"context"
	"log"
)

type serviceProvider struct {
	pgConfig   *config.PgConfig
	grpcConfig *config.GRPCConfig

	dbClient  database.Client
	txManeger database.TxManeger
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

func (s *serviceProvider) DBClient(ctx context.Context) database.Client {
	if s.dbClient == nil {
		dbc, err := postgres.NewClientPG(ctx, s.PgConfig().DSN())
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
		s.txManeger = transaction.NewTransactionManeger(s.DBClient(ctx).DB())
	}

	return s.txManeger
}
