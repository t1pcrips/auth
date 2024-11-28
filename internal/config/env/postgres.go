package env

import (
	"auth/internal/config"
	"errors"
	"os"
	"strconv"
)

const (
	hostPG   = "PG_HOST"
	portPG   = "PG_PORT"
	user     = "POSTGRES_USER"
	name     = "POSTGRES_DB"
	password = "POSTGRES_PASSWORD"
)

type PgConfigSearcher struct{}

func NewPgConfigSearcher() *PgConfigSearcher {
	return &PgConfigSearcher{}
}

func (s *PgConfigSearcher) Get() (*config.PgConfig, error) {
	dbHost := os.Getenv(hostPG)
	if len(dbHost) == 0 {
		return nil, errors.New("pgdb Host not found")
	}

	dbPort := os.Getenv(portPG)
	if len(dbPort) == 0 {
		return nil, errors.New("pgdb Port not found")
	}

	dbPortInt, err := strconv.Atoi(dbPort)
	if err != nil {
		return nil, errors.New("incorrect port, use integer port")
	}

	dbUser := os.Getenv(user)
	if len(dbUser) == 0 {
		return nil, errors.New("pgdb User not found")
	}

	dbName := os.Getenv(name)
	if len(dbName) == 0 {
		return nil, errors.New("pgdb Name not found")
	}

	dbPass := os.Getenv(password)
	if len(dbPass) == 0 {
		return nil, errors.New("db Pass not found")
	}

	return &config.PgConfig{
		Port:     dbPortInt,
		Host:     dbHost,
		User:     dbUser,
		Name:     dbName,
		Password: dbPass,
	}, nil
}
