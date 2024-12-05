package user

import (
	"context"
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/pkg/errs"
)

func (s *UserServiceImpl) Get(ctx context.Context, userId int64) (*model.GetUserResponse, error) {
	info, err := s.userRepository.Get(ctx, userId)
	if err != nil {
		return nil, errs.ErrGetUser
	}

	return info, nil
}
