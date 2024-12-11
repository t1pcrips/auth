package auth

import (
	"github.com/t1pcrips/auth/internal/service"
	"github.com/t1pcrips/auth/pkg/auth_v1"
)

type AuthApiImpl struct {
	auth_v1.UnimplementedAuthServer
	service service.AuthService
}

func NewAuthApiImpl(authService service.AuthService) *AuthApiImpl {
	return &AuthApiImpl{
		service: authService,
	}
}
