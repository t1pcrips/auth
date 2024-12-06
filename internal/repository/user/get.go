package user

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/internal/repository/user/converter_user"
	"github.com/t1pcrips/auth/internal/repository/user/model_user"
	"github.com/t1pcrips/auth/pkg/errs"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

func (repo *UserRepositoryImpl) Get(ctx context.Context, userId int64) (*model.GetUserResponse, error) {
	builderGetUser := squirrel.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableUsers).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: userId})

	query, args, err := builderGetUser.ToSql()
	if err != nil {
		return nil, errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "user_v1 repository - get user_v1",
		QueryRow: query,
	}

	var info model_user.GetUserResponse

	err = repo.db.DB().ScanOneContext(ctx, &info, q, args...)
	if err != nil {
		return nil, errs.ErrQueryRowScan
	}
	return converter_user.ToGetUserServiceFromRepo(&info), nil
}
