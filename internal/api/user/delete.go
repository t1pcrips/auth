package user

import (
	"context"
	"github.com/t1pcrips/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *UserApiImpl) Delete(ctx context.Context, info *user_v1.DeleteRequest) (*emptypb.Empty, error) {
	err := i.service.Delete(ctx, info.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
