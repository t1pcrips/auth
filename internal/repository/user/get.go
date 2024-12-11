package user

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/t1pcrips/auth/internal/errs"
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/internal/repository/user/converter_user"
	"github.com/t1pcrips/auth/internal/repository/user/model_user"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

func (repo *UserRepositoryImpl) GetByParams(ctx context.Context, info *model.Params) (*model.GetUserResponse, error) {
	var builderGetUser squirrel.SelectBuilder

	switch {
	case info.Email != nil:
		builderGetUser = squirrel.Select(idColumn, nameColumn, emailColumn, passwordColumn, roleColumn).
			From(tableUsers).
			PlaceholderFormat(squirrel.Dollar).
			Where(squirrel.Eq{emailColumn: *info.Email})
	default:
		builderGetUser = squirrel.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
			From(tableUsers).
			PlaceholderFormat(squirrel.Dollar).
			Where(squirrel.Eq{idColumn: *info.Id})
	}

	query, args, err := builderGetUser.ToSql()
	if err != nil {
		return nil, errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "user_v1 repository - get user_v1",
		QueryRow: query,
	}

	var resp model_user.GetUserResponse

	err = repo.db.DB().ScanOneContext(ctx, &resp, q, args...)
	if err != nil {
		return nil, errs.ErrQueryRowScan
	}

	return converter_user.ToGetUserServiceFromRepo(&resp), nil
}
