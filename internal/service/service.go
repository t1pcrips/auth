package service

import (
	"context"
	"github.com/t1pcrips/auth/internal/model"
)

//go:generate ../../bin/mockery --output ./mocks --inpackage-suffix --all --case snake

type UserService interface {
	Create(ctx context.Context, info *model.CreateUserRequest) (int64, error)
	Get(ctx context.Context, chatId int64) (*model.GetUserResponse, error)
	Update(ctx context.Context, info *model.UpdatUsereRequest) error
	Delete(ctx context.Context, userId int64) error
}

type AuthService interface {
	Login(ctx context.Context, info *model.User) (*model.Tokens, error)
	RefreshTokens(ctx context.Context, token string) (*model.Tokens, error)
}

type JWTService interface {
	GenerateAccessToken(user *model.User) (string, error)
	GenerateRefreshToken(ctx context.Context, user *model.User) (string, error)
	ValidateRefreshToken(ctx context.Context, signedToken string) (*model.User, error)
	ValidateTokenFromMeatadata(ctx context.Context) (*model.User, error)
}

type AccessService interface {
	Create(ctx context.Context, address string) (int64, error)
	Check(ctx context.Context, address string) error
}
