package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"net"
	"os"
)

type GRPCConfig struct {
	Host string
	Port string
}

type PgConfig struct {
	Port     int
	Host     string
	User     string
	Name     string
	Password string
}

type LogConfig struct {
	LogLevel      zerolog.Level
	LogTimeFormat string
}

func (cfg *GRPCConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *LogConfig) SetUp() *zerolog.Logger {
	settingsLoger := zerolog.ConsoleWriter{TimeFormat: cfg.LogTimeFormat, Out: os.Stdout, NoColor: false}
	logger := zerolog.New(settingsLoger).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(cfg.LogLevel)
	zerolog.TimeFieldFormat = cfg.LogTimeFormat
	return &logger
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("failed to load .env file %w", err)
	}

	return nil
}
