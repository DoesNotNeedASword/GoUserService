package repository

import (
	"Test2/internal/user"
	"Test2/internal/user/model"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	rows, err := r.pool.Query(context.Background(), "select id, name from users")
	defer rows.Close()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		fmt.Fprintf(os.Stderr, "Error %v \n", err)
		return nil, err
	}
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.User])

	return users, nil
}
