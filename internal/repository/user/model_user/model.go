package model_user

import (
	"time"
)

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
	Role     string
}

type GetUserResponse struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateUserRequest struct {
	ID    int64
	Name  string
	Email string
	Role  string
}
