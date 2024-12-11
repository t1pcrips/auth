package model

import (
	"time"
)

type UserRole string

const (
	UNKNOWN UserRole = "UNKNOWN"
	ADMIN   UserRole = "ROLE_ADMIN"
	USER    UserRole = "ROLE_USER"
)

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
	Role     UserRole
}

type GetUserResponse struct {
	Id        int64
	Name      string
	Email     string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetUserRequest struct {
	Id    int64
	Name  string
	Email string
	Role  UserRole
}

type UpdatUsereRequest struct {
	ID    int64
	Name  string
	Email string
	Role  UserRole
}
