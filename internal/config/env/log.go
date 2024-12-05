package env

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/t1pcrips/auth/internal/config"
	"os"
	"strconv"
)

const (
	logLevel      = "LOG_LEVEL"
	logTimeFormat = "LOG_TIME_FORMAT"
)

type LogConfigSearcher struct{}

func NewLogConfigSearcher() *LogConfigSearcher {
	return &LogConfigSearcher{}
}

func (s *LogConfigSearcher) Get() (*config.LogConfig, error) {
	level := os.Getenv(logLevel)
	if len(level) == 0 {
		return nil, fmt.Errorf("logLevel not found")
	}

	intLevel, err := strconv.Atoi(level)
	if err != nil {
		return nil, fmt.Errorf("failed to conver level to int: %w", err)
	}

	timeFormat := os.Getenv(logTimeFormat)
	if len(timeFormat) == 0 {
		return nil, fmt.Errorf("timeFormat not found")
	}

	return &config.LogConfig{
		LogLevel:      zerolog.Level(intLevel),
		LogTimeFormat: timeFormat,
	}, nil
}
