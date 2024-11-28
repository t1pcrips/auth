package user

import (
	"auth/internal/repository"
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"time"
)

const (
	tableName       = "users"
	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type UserRepo struct {
	pool   *pgxpool.Pool
	logger *zerolog.Logger
}

func NewUserRepo(pool *pgxpool.Pool, logger *zerolog.Logger) *UserRepo {
	return &UserRepo{
		pool:   pool,
		logger: logger,
	}
}

func (repo *UserRepo) Create(ctx context.Context, req *repository.UserCreateRequest) (int64, error) {
	builderInsert := squirrel.Insert(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn, createdAtColumn, updatedAtColumn).
		Values(req.Name, req.Email, req.Password, req.Role, time.Now(), time.Now()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed insert builds the query into a SQL string and bound args: %w", err)
	}

	var userId int64

	err = repo.pool.QueryRow(ctx, query, args...).Scan(&userId)
	if err != nil {
		return 0, fmt.Errorf("failed to acquires a connection and executes a query: %w", err)
	}

	return userId, nil
}

func (repo *UserRepo) Get(ctx context.Context, userId int64) (*repository.UserGetResponse, error) {
	builderSelect := squirrel.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: userId})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed select builds the query into a SQL string and bound args: %w", err)
	}

	rows, err := repo.pool.Query(ctx, query, args...)
	defer rows.Close()

	resp, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[repository.UserGetResponse])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("failed no rows: %w", err)
		}
		return nil, fmt.Errorf("failed to collet in one row with struct by name: %w", err)
	}
	return &resp, nil
}

func (repo *UserRepo) Update(ctx context.Context, req *repository.UserUpdateRequest) error {
	builderUpdate := squirrel.Update(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Set(updatedAtColumn, time.Now()).
		Set(nameColumn, req.Name).
		Set(emailColumn, req.Email).
		Set(roleColumn, req.Role).
		Where(squirrel.Eq{idColumn: req.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return fmt.Errorf("failed update builds the query into a SQL string and bound args: %w", err)
	}

	_, err = repo.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed update exec acquires a connection from the Pool and executes the given SQL: %w", err)
	}

	return nil
}

func (repo *UserRepo) Delete(ctx context.Context, userId int64) error {
	builderDelete := squirrel.Delete(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: userId})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return fmt.Errorf("failed delete builds the query into a SQL string and bound args: %w", err)
	}

	result, err := repo.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed delete exec acquires a connection from the Pool and executes the given SQL: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("no rows deleted")
	}

	return nil
}
