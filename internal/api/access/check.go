package access

import (
	"context"
	"github.com/t1pcrips/auth/pkg/access_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *AccessApiImpl) Check(ctx context.Context, info *access_v1.CheckRequest) (*emptypb.Empty, error) {
	err := i.service.Check(ctx, info.GetAddress())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
