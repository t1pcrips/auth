package app

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/t1pcrips/auth/internal/config"
	"github.com/t1pcrips/auth/internal/interceptor"
	desc "github.com/t1pcrips/auth/pkg/user_v1"
	"github.com/t1pcrips/platform-pkg/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"sync"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
	grpcServer      *grpc.Server
	configPath      string
}

func NewApp(ctx context.Context, configPath string) (*App, error) {
	a := &App{configPath: configPath}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("failed to run http server: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runGRPCServer()
		if err != nil {
			log.Fatalf("failed to run grpc server: %s", err.Error())
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initConfig,
		a.initGRPCServer,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	err := config.Load(a.configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)

	reflection.Register(a.grpcServer)
	desc.RegisterUserServer(a.grpcServer, a.serviceProvider.UserApiImpl(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := desc.RegisterUserHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts)
	if err != nil {
		return err
	}

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: mux,
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("gRPC Server starts on: %s", a.serviceProvider.GRPCConfig().Address())

	lis, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP Server starts on: %s", a.serviceProvider.HTTPConfig().Address())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
