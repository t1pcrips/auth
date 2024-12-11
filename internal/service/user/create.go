package user

import (
	"context"
	"github.com/t1pcrips/auth/internal/errs"
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/internal/utils"
)

func (s *UserServiceImpl) Create(ctx context.Context, info *model.CreateUserRequest) (int64, error) {
	hashedPassword, err := utils.SecureHash(info.Password)
	if err != nil {
		return 0, err
	}

	info.Password = hashedPassword

	userId, err := s.userRepository.Create(ctx, info)
	if err != nil {
		return 0, errs.ErrCreateUser
	}

	return userId, nil
}
