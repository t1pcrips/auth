package converter

import (
	"github.com/t1pcrips/auth/internal/model"
	dst "github.com/t1pcrips/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCreateRequestApiFromDst(info *dst.CreateRequest) *model.CreateUserRequest {
	return &model.CreateUserRequest{
		Name:     info.GetName(),
		Email:    info.GetEmail(),
		Password: info.GetPassword(),
		Role:     ToServiceRole(info.Role),
	}
}

func ToUpdateRequestApiFromDst(info *dst.UpdateRequest) *model.UpdatUsereRequest {
	return &model.UpdatUsereRequest{
		ID:    info.GetId(),
		Name:  info.GetName(),
		Email: info.GetEmail(),
		Role:  ToServiceRole(info.GetRole()),
	}
}

func ToDstGetFromGetApi(info *model.GetUserResponse) *dst.GetResponse {
	return &dst.GetResponse{
		Name:      info.Name,
		Email:     info.Email,
		Role:      ToRoleFromServiceRole(info.Role),
		CreatedAt: timestamppb.New(info.CreatedAt),
		UpdatedAt: timestamppb.New(info.UpdatedAt),
	}
}

func IdToCreateResponse(id int64) *dst.CreateResponse {
	return &dst.CreateResponse{
		Id: id,
	}
}

func ToRoleFromServiceRole(role model.UserRole) dst.Role {
	resultRole := dst.Role_value[string(role)]
	return dst.Role(resultRole)
}

func ToServiceRole(role dst.Role) model.UserRole {
	resultRole := dst.Role_name[int32(role)]
	return model.UserRole(resultRole)
}
