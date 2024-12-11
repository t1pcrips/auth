package access

import (
	"context"
	"github.com/t1pcrips/auth/internal/errs"
	"github.com/t1pcrips/auth/internal/model"
)

func (s *AccessServiceImpl) Create(ctx context.Context, address string) (int64, error) {
	user, err := s.jwtService.ValidateTokenFromMeatadata(ctx)
	if err != nil {
		return 0, err
	}
	
	if user.Role != string(model.ADMIN) {
		return 0, errs.ErrNeedAdminRole
	}

	idAccess, err := s.accessRepository.Create(ctx, &model.AccessAddress{
		EndpointAddress: address,
		Role:            user.Role,
	})
	if err != nil {
		return 0, err
	}

	return idAccess, nil
}
