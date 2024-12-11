package access

import (
	"context"
	"github.com/t1pcrips/auth/internal/model"
)

func (s *AccessServiceImpl) Check(ctx context.Context, address string) error {
	user, err := s.jwtService.ValidateTokenFromMeatadata(ctx)
	if err != nil {
		return err
	}

	if user.Role == string(model.ADMIN) {
		return nil
	}

	err = s.accessRepository.Check(ctx, &model.AccessAddress{
		EndpointAddress: address,
		Role:            user.Role,
	})
	if err != nil {
		return err
	}

	return nil
}
