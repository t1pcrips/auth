package user

import (
	"context"
	"github.com/t1pcrips/auth/internal/converter"
	"github.com/t1pcrips/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *UserApiImpl) Get(ctx context.Context, info *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	response, err := i.service.Get(ctx, info.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resultResponse := converter.ToDstGetFromGetApi(response)

	return resultResponse, nil
}
