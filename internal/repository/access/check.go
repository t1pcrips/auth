package access

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/t1pcrips/auth/internal/errs"
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

func (repo *AccessRepositoryImpl) Check(ctx context.Context, info *model.AccessAddress) error {
	builderChechEndpoint := squirrel.Select("1").
		PlaceholderFormat(squirrel.Dollar).
		From(tableAccessEndpoints).
		Where(squirrel.And{
			squirrel.Eq{endpointAddressColumn: info.EndpointAddress},
			squirrel.Eq{roleColumn: info.Role},
		})

	query, args, err := builderChechEndpoint.ToSql()
	if err != nil {
		return errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "access_v1 repository check endpoint access",
		QueryRow: query,
	}

	var exists string

	err = repo.db.DB().ScanOneContext(ctx, &exists, q, args...)
	if err != nil {
		return errs.ErrQueryRowScan
	}

	if exists != "1" {
		return errs.ErrEndpointNotFound
	}

	return nil
}
