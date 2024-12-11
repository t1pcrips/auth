package auth

import (
	"context"
	"github.com/t1pcrips/auth/internal/model"
)

func (s *AuthServiceImpl) RefreshTokens(ctx context.Context, token string) (*model.Tokens, error) {
	user, err := s.jwtService.ValidateRefreshToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return s.generateTokensPair(ctx, user)
}
