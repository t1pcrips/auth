package access

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/t1pcrips/auth/internal/errs"
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

func (repo *AccessRepositoryImpl) Create(ctx context.Context, info *model.AccessAddress) (int64, error) {
	builderCreateEndpoint := squirrel.Insert(tableAccessEndpoints).
		PlaceholderFormat(squirrel.Dollar).
		Columns(roleColumn, endpointAddressColumn).
		Values(info.Role, info.EndpointAddress).
		Suffix(returningId)

	query, args, err := builderCreateEndpoint.ToSql()
	if err != nil {
		return 0, errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "access_v1 repository create endpoint access",
		QueryRow: query,
	}

	var id int64
	err = repo.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, errs.ErrQueryRowScan
	}

	return id, nil
}
