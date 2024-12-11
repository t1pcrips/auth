package access

import (
	"github.com/t1pcrips/auth/internal/repository"
	"github.com/t1pcrips/auth/internal/service"
)

type AccessServiceImpl struct {
	jwtService       service.JWTService
	accessRepository repository.AccessRepository
}

func NewAccessServiceImpl(jwtService service.JWTService, accessRepository repository.AccessRepository) service.AccessService {
	return &AccessServiceImpl{
		jwtService:       jwtService,
		accessRepository: accessRepository,
	}
}
