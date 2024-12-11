package user

import (
	"context"
	"github.com/t1pcrips/auth/internal/converter"
	"github.com/t1pcrips/auth/internal/errs"
	"github.com/t1pcrips/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *UserApiImpl) Create(ctx context.Context, info *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	if info.GetPassword() != info.GetPasswordConfirm() {
		return nil, status.Error(codes.InvalidArgument, errs.ErrInvalidPasswords.Error())
	}

	userId, err := i.service.Create(ctx, converter.ToCreateRequestApiFromDst(info))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return converter.IdToCreateResponse(userId), nil
}
