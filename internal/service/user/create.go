package user

import (
	"context"
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/pkg/errs"
)

func (s *UserServiceImpl) Create(ctx context.Context, info *model.CreateUserRequest) (int64, error) {
	var userId int64

	err := s.txManeger.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error

		userId, txErr = s.userRepository.Create(ctx, info)
		if txErr != nil {
			return errs.ErrCreateUser
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return userId, nil
}
