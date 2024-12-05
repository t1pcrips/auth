package user

import (
	"github.com/t1pcrips/auth/internal/service"
	dst "github.com/t1pcrips/auth/pkg/user_v1"
)

type UserApiImpl struct {
	dst.UnimplementedUserServer
	service service.UserService
}

func NewUserApiImpl(service service.UserService) *UserApiImpl {
	return &UserApiImpl{
		service: service,
	}
}
