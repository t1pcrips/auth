package main

import (
	deps "auth/pkg/auth_v1"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"net"
)

type server struct {
	deps.UnimplementedAuthServer
}

func (s *server) Create(ctx context.Context, req *deps.CreateRequest) (*deps.CreateResponse, error) {
	log.Printf("create: %v\n", req)
	return &deps.CreateResponse{
		Id: rand.Int63(),
	}, nil
}

func (s *server) Get(ctx context.Context, req *deps.GetRequest) (*deps.GetResponse, error) {
	log.Printf("Get: %+v\n", req)
	return &deps.GetResponse{Id: rand.Int63(), UpdatedAt: timestamppb.Now()}, nil
}

func (s *server) Update(ctx context.Context, req *deps.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("update: %+v\n", req)
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *deps.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("delete: %+v", req)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err.Error())
	}

	s := grpc.NewServer()
	reflection.Register(s)
	deps.RegisterAuthServer(s, &server{})

	log.Fatal(s.Serve(lis))
}
