package access

import (
	"github.com/t1pcrips/auth/internal/service"
	"github.com/t1pcrips/auth/pkg/access_v1"
)

type AccessApiImpl struct {
	access_v1.UnimplementedAccessServer
	service service.AccessService
}

func NewAccessApiImpl(accessService service.AccessService) *AccessApiImpl {
	return &AccessApiImpl{
		service: accessService,
	}
}
