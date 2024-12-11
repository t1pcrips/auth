package user

import (
	"github.com/t1pcrips/auth/internal/service"
	"github.com/t1pcrips/auth/pkg/user_v1"
)

type UserApiImpl struct {
	user_v1.UnimplementedUserServer
	service service.UserService
}

func NewUserApiImpl(service service.UserService) *UserApiImpl {
	return &UserApiImpl{
		service: service,
	}
}
