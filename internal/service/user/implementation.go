package user

import (
	"github.com/t1pcrips/auth/internal/repository"
	"github.com/t1pcrips/auth/internal/service"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	txManeger      database.TxManeger
}

func NewUserServiceImpl(
	userRepository repository.UserRepository,
	txManeger database.TxManeger,
) service.UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		txManeger:      txManeger,
	}
}
