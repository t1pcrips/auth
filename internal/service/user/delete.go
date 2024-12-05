package user

import (
	"context"
	"github.com/t1pcrips/auth/pkg/errs"
)

func (s *UserServiceImpl) Delete(ctx context.Context, userId int64) error {
	err := s.userRepository.Delete(ctx, userId)
	if err != nil {
		return errs.ErrDeleteUser
	}

	return nil
}
