package user

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/t1pcrips/auth/internal/errs"
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/internal/repository/user/converter_user"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

func (repo *UserRepositoryImpl) Create(ctx context.Context, info *model.CreateUserRequest) (int64, error) {
	repoInfo := converter_user.ToCreateUserRequestFromService(info)

	buidlerCreteUser := squirrel.Insert(tableUsers).
		PlaceholderFormat(squirrel.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(repoInfo.Name, repoInfo.Email, repoInfo.Password, repoInfo.Role).
		Suffix(returningId)

	query, args, err := buidlerCreteUser.ToSql()
	if err != nil {
		return 0, errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "user_v1 repository - create user_v1",
		QueryRow: query,
	}

	var userId int64

	err = repo.db.DB().ScanOneContext(ctx, &userId, q, args...)
	if err != nil {
		return 0, errs.ErrQueryRowScan
	}

	return userId, nil
}
