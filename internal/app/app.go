package app

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	grpcServer *grpc.Server
	configPath string
}

func NewApp() {}

func (a *App) Run() error {}

func (a *App) initDeps() error {}

func (a *App) initConfig() error {}

func (a *App) initGRPCServer() error {
	_ = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
}

func (a *App) initServiceProvider() {}
