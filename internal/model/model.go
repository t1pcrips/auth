package model

import (
	"time"
)

type UserRole string

const (
	UNKNOWN UserRole = "UNKNOWN"
	ADMIN   UserRole = "ADMIN"
	USER    UserRole = "USER"
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
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time `db:"updated_at"`
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
