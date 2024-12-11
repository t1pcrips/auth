package user

import (
	"context"
	"github.com/t1pcrips/auth/internal/converter"
	"github.com/t1pcrips/auth/internal/errs"
	"github.com/t1pcrips/auth/internal/model"
)

func (s *UserServiceImpl) Get(ctx context.Context, userId int64) (*model.GetUserResponse, error) {
	info, err := s.userRepository.GetByParams(ctx, converter.ToParamsById(userId))
	if err != nil {
		return nil, errs.ErrGetUser
	}

	return info, nil
}
