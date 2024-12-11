package access

import (
	"github.com/t1pcrips/auth/internal/repository"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

const (
	idColumn              = "id"
	roleColumn            = "role"
	endpointAddressColumn = "endpoint_address"
	tableAccessEndpoints  = "access_endpoints"
	returningId           = "RETURNING id"
)

type AccessRepositoryImpl struct {
	db database.Client
}

func NewAccessRepositoryImpl(db database.Client) repository.AccessRepository {
	return &AccessRepositoryImpl{
		db: db,
	}
}
