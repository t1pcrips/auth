package env

import (
	"errors"
	"github.com/t1pcrips/auth/internal/config"
	"os"
)

const (
	hostGRPC = "GRPC_HOST"
	portGRPC = "GRPC_PORT"
)

type GRPCConfigSearcher struct{}

func NewGRPCConfigSearcher() *GRPCConfigSearcher {
	return &GRPCConfigSearcher{}
}

func (s *GRPCConfigSearcher) Get() (*config.GRPCConfig, error) {
	grpcHost := os.Getenv(hostGRPC)
	if len(grpcHost) == 0 {
		return nil, errors.New("gRPC Host not found")
	}

	grpcPort := os.Getenv(portGRPC)
	if len(grpcPort) == 0 {
		return nil, errors.New("gRPC Port not found")
	}

	return &config.GRPCConfig{
		Host: grpcHost,
		Port: grpcPort,
	}, nil
}