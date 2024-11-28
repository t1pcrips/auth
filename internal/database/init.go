package database

import (
	"auth/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const (
	pgURLFormat = "postgres://%s:%s@%s:%d/%s"
)

func InitPostgresConnection(ctx context.Context, cfg *config.PgConfig) (*pgxpool.Pool, func(), error) {
	if cfg == nil {
		return nil, nil, fmt.Errorf("faild to create new database with nil config")
	}

	pgxConfig, err := pgxpool.ParseConfig(getPgURL(cfg))
	if err != nil {
		return nil, nil, fmt.Errorf("faild to parse pgx config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	closer := func() {
		if err := pool.Close; err != nil {
			log.Printf("Failed to close database pool: %v", err)
		}
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to ping to database: %w", err)
	}

	return pool, closer, nil
}

// "postgres://username:password@localhost:5432/database_name"
func getPgURL(cfg *config.PgConfig) string {
	return fmt.Sprintf(
		pgURLFormat,
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)
}
