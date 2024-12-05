package user

import (
	"github.com/t1pcrips/auth/internal/repository"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

const (
	tableUsers      = "users"
	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
	returningId     = "RETURNING id"
)

type UserRepositoryImpl struct {
	db database.Client
}

func NewUserRepositoryImpl(db database.Client) repository.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}
