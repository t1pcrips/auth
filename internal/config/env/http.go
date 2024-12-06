package env

import (
	"errors"
	"github.com/t1pcrips/auth/internal/config"
	"os"
	"strconv"
)

const (
	portHTTP = "HTTP_PORT"
	hostHTTP = "HTTP_HOST"
)

type HTTPCfgSearcher struct{}

func NewHTTPCfgSearcher() *HTTPCfgSearcher {
	return &HTTPCfgSearcher{}
}

func (cfg *HTTPCfgSearcher) Get() (*config.HTTPConfig, error) {
	host := os.Getenv(hostHTTP)
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}

	portString := os.Getenv(portHTTP)
	if len(portString) == 0 {
		return nil, errors.New("http port not found")
	}
	_, err := strconv.Atoi(portString)
	if err != nil {
		return nil, errors.New("faild to convert port to integer")
	}

	return &config.HTTPConfig{
		Host: host,
		Port: portString,
	}, nil
}
