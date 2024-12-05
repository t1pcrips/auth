package user

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/internal/repository/user/converter_user"
	"github.com/t1pcrips/auth/pkg/errs"
	"github.com/t1pcrips/platform-pkg/pkg/database"
	"time"
)

func (repo *UserRepositoryImpl) Update(ctx context.Context, info *model.UpdatUsereRequest) error {
	repoInfo := converter_user.ToUpdateUserRequestFromService(info)

	builderUpdateUser := squirrel.Update(tableUsers).
		PlaceholderFormat(squirrel.Dollar).
		Set(nameColumn, repoInfo.Name).
		Set(emailColumn, repoInfo.Email).
		Set(roleColumn, repoInfo.Role).
		Set(updatedAtColumn, time.Now()).
		Where(squirrel.Eq{idColumn: repoInfo.ID})

	query, args, err := builderUpdateUser.ToSql()
	if err != nil {
		return errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "user repository - update user",
		QueryRow: query,
	}

	result, err := repo.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return errs.ErrExec
	}

	if result.RowsAffected() == 0 {
		return errs.ErrNoRowsUpdate
	}

	return nil
}
