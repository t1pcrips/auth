package repository

import (
	"context"
	"github.com/t1pcrips/auth/internal/model"
)

//go:generate ../../bin/mockery --output ./mocks --inpackage-suffix --all --case snake

type UserRepository interface {
	Create(ctx context.Context, info *model.CreateUserRequest) (int64, error)
	GetByParams(ctx context.Context, info *model.Params) (*model.GetUserResponse, error)
	Update(ctx context.Context, info *model.UpdatUsereRequest) error
	Delete(ctx context.Context, userId int64) error
}

type CacheRepository interface {
	Set(ctx context.Context, key string, value interface{}, ttl int64) error
	Get(ctx context.Context, key string) ([]byte, error)
}

type AccessRepository interface {
	Create(ctx context.Context, info *model.AccessAddress) (int64, error)
	Check(ctx context.Context, info *model.AccessAddress) error
}
