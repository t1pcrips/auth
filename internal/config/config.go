package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"net"
	"time"
)

const (
	postgresDSNExample = "postgres://%s:%s@%s:%d/%s"
)

type GRPCConfig struct {
	Host string
	Port string
}

type HTTPConfig struct {
	Host string
	Port string
}

type SwagerConfig struct {
	Host string
	Port string
}

type PgConfig struct {
	Host     string
	Port     int
	User     string
	Name     string
	Password string
}

type RedisConfig struct {
	Host        string
	Port        string
	MaxIdle     int
	IdleTimeout time.Duration
	CtxTimeout  time.Duration
}

type LogConfig struct {
	LogLevel      zerolog.Level
	LogTimeFormat string
}

func (cfg *GRPCConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *HTTPConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *RedisConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *SwagerConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *PgConfig) DSN() string {
	return fmt.Sprintf(
		postgresDSNExample,
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("failed to load .env file %w", err)
	}

	return nil
}
