package env

import (
	"errors"
	"github.com/t1pcrips/auth/internal/config"
	"os"
	"strconv"
)

const (
	hostPG     = "PG_HOST"
	portPG     = "PG_PORT"
	userPG     = "POSTGRES_USER"
	namePG     = "POSTGRES_DB"
	passwordPG = "POSTGRES_PASSWORD"
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

	dbUser := os.Getenv(userPG)
	if len(dbUser) == 0 {
		return nil, errors.New("pgdb User not found")
	}

	dbName := os.Getenv(namePG)
	if len(dbName) == 0 {
		return nil, errors.New("pgdb Name not found")
	}

	dbPass := os.Getenv(passwordPG)
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
