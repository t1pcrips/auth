package service

import (
	"auth/internal/common"
	"auth/internal/repository"
	"auth/internal/repository/user"
	deps "auth/pkg/user_v1"
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServer struct {
	deps.UnimplementedUserServer
	repository *user.UserRepo
	logger     *zerolog.Logger
}

func NewUserServer(repository *user.UserRepo, logger *zerolog.Logger) *UserServer {
	return &UserServer{
		repository: repository,
		logger:     logger,
	}
}

func (s *UserServer) Create(ctx context.Context, req *deps.CreateRequest) (*deps.CreateResponse, error) {
	s.logRequest("CREATE", req)

	if err := s.ValidateCreateRequest(req); err != nil {
		return nil, err
	}

	id, err := s.repository.Create(ctx, &repository.UserCreateRequest{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     req.GetRole().String(),
	})
	if err != nil {
		s.logger.Err(err).Msg("failed to create user")
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &deps.CreateResponse{
		Id: id,
	}, nil
}

func (s *UserServer) Get(ctx context.Context, req *deps.GetRequest) (*deps.GetResponse, error) {
	s.logRequest("GET", req)

	id := req.GetId()

	get, err := s.repository.Get(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, fmt.Errorf("failed no rows: %w", err)):
			s.logger.Info().Msg("failed to get user with this request")
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			s.logger.Err(err).Msg("internal server error with get user")
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &deps.GetResponse{
		Id:        get.ID,
		Name:      get.Name,
		Email:     get.Email,
		Role:      common.RoleToValue(get.Role),
		CreatedAt: timestamppb.New(get.CreatedAt),
		UpdatedAt: timestamppb.New(get.UpdatedAt),
	}, nil
}

func (s *UserServer) Update(ctx context.Context, req *deps.UpdateRequest) (*emptypb.Empty, error) {
	s.logRequest("UPDATE", req)

	_, err := s.repository.Get(ctx, req.GetId())
	if err != nil {
		switch {
		case errors.Is(err, fmt.Errorf("failed no rows: %w", err)):
			s.logger.Info().Msg("failed to get user with this request")
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			s.logger.Err(err).Msg("internal server error with check user")
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	err = s.repository.Update(ctx, &repository.UserUpdateRequest{
		ID:    req.GetId(),
		Name:  common.ToPointer(req.GetName().GetValue()),
		Email: common.ToPointer(req.GetEmail().GetValue()),
		Role:  common.ToPointer(req.GetRole().String()),
	})

	if err != nil {
		switch {
		case errors.Is(err, fmt.Errorf("failed nothing to update pls use data")):
			s.logger.Info().Msg("failed to get user with this request")
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			s.logger.Err(err).Msg("internal server error with update user")
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *UserServer) Delete(ctx context.Context, req *deps.DeleteRequest) (*emptypb.Empty, error) {
	s.logRequest("DELETE", req)

	_, err := s.repository.Get(ctx, req.GetId())
	if err != nil {
		switch {
		case errors.Is(err, fmt.Errorf("failed no rows: %w", err)):
			s.logger.Info().Msg("failed to get user with this request")
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			s.logger.Err(err).Msg("internal server error with check user")
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	err = s.repository.Delete(ctx, req.GetId())
	if err != nil {
		s.logger.Err(err).Msg("failed to delete user")
		return nil, status.Error(codes.NotFound, "faild to delete user with this id")
	}

	return &emptypb.Empty{}, nil
}

func (s *UserServer) logRequest(method string, req interface{}) {
	s.logger.Debug().
		Str("method", method).
		Interface("request", req).Msg("try to process user request")
}

func (s *UserServer) ValidateCreateRequest(req *deps.CreateRequest) error {
	if req.GetPassword() != req.GetPasswordConfirm() {
		s.logger.Info().Msg("passwords do not match")
		return status.Error(codes.InvalidArgument, "passwords do not match")
	}

	if len(req.GetName()) == 0 {
		s.logger.Info().Msg("name is required")
		return status.Error(codes.InvalidArgument, "name is required")
	}

	if len(req.GetEmail()) == 0 {
		s.logger.Info().Msg("email is required")
		return status.Error(codes.InvalidArgument, "email is required")
	}

	if len(req.GetPassword()) == 0 {
		s.logger.Info().Msg("password is required")
		return status.Error(codes.InvalidArgument, "password is required")
	}

	if len(req.GetRole().String()) == 0 {
		s.logger.Info().Msg("role is required")
		return status.Error(codes.InvalidArgument, "role is required")
	}

	return nil
}
