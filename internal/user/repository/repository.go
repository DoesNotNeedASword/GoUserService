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

func (r *UserRepository) GetUsers(ctx context.Context) (model.User, error) {
	var id int64
	var name string
	err := r.pool.QueryRow(context.Background(), "select id, name from users").Scan(&id, &name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, user.ErrNotFound
		}
		fmt.Fprintf(os.Stderr, "Error %v \n", err)
		return model.User{}, err
	}
	user := model.User{ID: id, Name: name}
	fmt.Println(id, name)
	return user, nil
}
