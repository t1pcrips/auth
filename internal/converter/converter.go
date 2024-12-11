package converter

import (
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/pkg/auth_v1"
	"github.com/t1pcrips/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCreateRequestApiFromDst(info *user_v1.CreateRequest) *model.CreateUserRequest {
	return &model.CreateUserRequest{
		Name:     info.GetName(),
		Email:    info.GetEmail(),
		Password: info.GetPassword(),
		Role:     ToServiceRole(info.Role),
	}
}

func ToUpdateRequestApiFromDst(info *user_v1.UpdateRequest) *model.UpdatUsereRequest {
	return &model.UpdatUsereRequest{
		ID:    info.GetId(),
		Name:  info.GetName(),
		Email: info.GetEmail(),
		Role:  ToServiceRole(info.GetRole()),
	}
}

func ToDstGetFromGetApi(info *model.GetUserResponse) *user_v1.GetResponse {
	return &user_v1.GetResponse{
		Id:        info.Id,
		Name:      info.Name,
		Email:     info.Email,
		Role:      ToRoleFromServiceRole(info.Role),
		CreatedAt: timestamppb.New(info.CreatedAt),
		UpdatedAt: timestamppb.New(info.UpdatedAt),
	}
}

func ToUserFromLoginApiAuth(info *auth_v1.LoginRequest) *model.User {
	return &model.User{
		Email:    info.GetEmail(),
		Password: info.GetPassword(),
	}
}

func IdToCreateResponse(id int64) *user_v1.CreateResponse {
	return &user_v1.CreateResponse{
		Id: id,
	}
}

func ToRoleFromServiceRole(role model.UserRole) user_v1.Role {
	resultRole := user_v1.Role_value[string(role)]
	return user_v1.Role(resultRole)
}

func ToServiceRole(role user_v1.Role) model.UserRole {
	resultRole := user_v1.Role_name[int32(role)]
	return model.UserRole(resultRole)
}

func ToParamsByEmail(eamil string) *model.Params {
	em := &eamil
	return &model.Params{
		Id:    nil,
		Email: em,
	}
}

func ToParamsById(id int64) *model.Params {
	i := &id
	return &model.Params{
		Id:    i,
		Email: nil,
	}
}
