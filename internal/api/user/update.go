package user

import (
	"context"
	"github.com/t1pcrips/auth/internal/converter"
	"github.com/t1pcrips/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *UserApiImpl) Update(ctx context.Context, info *user_v1.UpdateRequest) (*emptypb.Empty, error) {
	err := i.service.Update(ctx, converter.ToUpdateRequestApiFromDst(info))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
