package repository

import (
	"context"
	"github.com/t1pcrips/auth/internal/model"
)

//go:generate ../../bin/mockery --output ./mocks --inpackage-suffix --all --case snake

type UserRepository interface {
	Create(ctx context.Context, info *model.CreateUserRequest) (int64, error)
	Get(ctx context.Context, chatId int64) (*model.GetUserResponse, error)
	Update(ctx context.Context, info *model.UpdatUsereRequest) error
	Delete(ctx context.Context, userId int64) error
}
