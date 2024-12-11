package auth

import (
	"context"
	"github.com/t1pcrips/auth/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *AuthApiImpl) RefreshTokens(ctx context.Context, info *auth_v1.RefreshTokensRequest) (*auth_v1.RefreshTokensResponse, error) {
	tokens, err := i.service.RefreshTokens(ctx, info.GetRefreshToken())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// обработать эту херню

	return &auth_v1.RefreshTokensResponse{
		RefreshToken: tokens.RefreshToken,
		AccessToken:  tokens.AccessToken,
	}, err
}
