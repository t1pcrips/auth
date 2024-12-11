package user

import (
	"context"
	"github.com/t1pcrips/auth/internal/errs"
	"github.com/t1pcrips/auth/internal/model"
)

func (s *UserServiceImpl) Update(ctx context.Context, info *model.UpdatUsereRequest) error {
	err := s.txManeger.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error

		txErr = s.userRepository.Update(ctx, info)
		if txErr != nil {
			return txErr
		}

		return nil
	})
	if err != nil {
		return errs.ErrUpdateUser
	}

	return nil
}
