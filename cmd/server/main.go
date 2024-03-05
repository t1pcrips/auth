package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	desc "github.com/t1pcrips/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	grpcPort = 50051
)

type server struct {
	desc.UnimplementedUserV1Server
}

// Create . . .
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create a new user - USER ID: %v\n", req.GetInfo())
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

// Get . . .
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get USER ID: %v\n", req.Id)
	return &desc.GetResponse{
		User: &desc.User{
			Id: gofakeit.Int64(),
			Info: &desc.UserInfo{
				Name:                 gofakeit.Name(),
				Email:                gofakeit.Email(),
				Password:             gofakeit.Password(true, false, true, false, false, 7),
				PasswordConfirmation: gofakeit.Password(true, false, true, false, false, 7),
				Role:                 1,
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

// Delete . . .
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Deleate ID: %v\n", req.Id)
	return new(emptypb.Empty), nil
}

// Update . . .
func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update ID: %v\n", req.Id)
	return new(emptypb.Empty), nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("faield to listen - %v\n", err)
	}

	s := grpc.NewServer()
	reflection.RegisterV1(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listeninng at %v\n", listen.Addr())

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve - %v\n", err)
	}
}
