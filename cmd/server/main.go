package main

import (
	"auth/internal/config"
	"auth/internal/config/env"
	"auth/internal/database"
	"auth/internal/repository/user"
	"auth/internal/service"
	deps "auth/pkg/user_v1"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "path to config file")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Println(err)
	}

	pgConfig, err := env.NewPgConfigSearcher().Get()
	if err != nil {
		log.Fatal(err)
	}

	grpcConfig, err := env.NewGRPCConfigSearcher().Get()
	if err != nil {
		log.Fatal(err)
	}

	logConfig, err := env.NewLogConfigSearcher().Get()
	if err != nil {
		log.Println(err)
	}

	logger := logConfig.SetUp()

	pool, closer, err := database.InitPostgresConnection(ctx, pgConfig)
	if err != nil {
		log.Fatal(err)
	}

	defer closer()

	repository := user.NewUserRepo(pool, logger)
	userServer := service.NewUserServer(repository, logger)

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	reflection.Register(server)
	deps.RegisterUserServer(server, userServer)

	log.Println("server starts...")
	log.Fatal(server.Serve(lis))
}
