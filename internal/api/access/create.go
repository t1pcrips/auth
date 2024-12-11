package access

import (
	"context"
	"github.com/t1pcrips/auth/pkg/access_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *AccessApiImpl) Create(ctx context.Context, info *access_v1.CreateRequest) (*access_v1.CreateResponse, error) {
	id, err := i.service.Create(ctx, info.Address)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &access_v1.CreateResponse{
		Id: id,
	}, nil
}
