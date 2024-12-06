package user

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/t1pcrips/auth/pkg/errs"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

func (repo *UserRepositoryImpl) Delete(ctx context.Context, userId int64) error {
	builderDeleteUser := squirrel.Delete(tableUsers).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: userId})

	query, args, err := builderDeleteUser.ToSql()

	q := database.Query{
		Name:     "user_v1 repository - delete user_v1",
		QueryRow: query,
	}

	result, err := repo.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return errs.ErrExec
	}

	if result.RowsAffected() == 0 {
		return errs.ErrNoRowsDelete
	}

	return nil
}
