package auth

import (
	"context"
	"github.com/t1pcrips/auth/internal/converter"
	"github.com/t1pcrips/auth/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *AuthApiImpl) Login(ctx context.Context, info *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	tokens, err := i.service.Login(ctx, converter.ToUserFromLoginApiAuth(info))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// обработчик ебануть

	return &auth_v1.LoginResponse{
		RefreshToken: tokens.RefreshToken,
		AccessToken:  tokens.AccessToken,
	}, nil
}
