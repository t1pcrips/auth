package converter_user

import (
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/internal/repository/user/model_user"
)

func ToGetUserServiceFromRepo(info *model_user.GetUserResponse) *model.GetUserResponse {
	return &model.GetUserResponse{
		Id:        info.ID,
		Name:      info.Name,
		Email:     info.Email,
		Role:      ToRoleFromString(info.Role),
		UpdatedAt: info.UpdatedAt,
		CreatedAt: info.CreatedAt,
	}
}

func ToCreateUserRequestFromService(info *model.CreateUserRequest) *model_user.CreateUserRequest {
	return &model_user.CreateUserRequest{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     string(info.Role),
	}
}

func ToUpdateUserRequestFromService(info *model.UpdatUsereRequest) *model_user.UpdateUserRequest {
	return &model_user.UpdateUserRequest{
		ID:    info.ID,
		Name:  info.Name,
		Email: info.Email,
		Role:  string(info.Role),
	}
}

func ToRoleFromString(role string) model.UserRole {
	return model.UserRole(role)
}
