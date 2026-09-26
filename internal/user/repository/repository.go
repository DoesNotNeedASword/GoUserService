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
	const q = `
        select id, name, tg_id, party_id
        from users
        order by id
    `
	rows, err := r.pool.Query(ctx, q)
	defer rows.Close()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		fmt.Fprintf(os.Stderr, "Error %v \n", err)
		return nil, err
	}
	u, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	if err != nil {
		return nil, err
	}
	users := make([]model.User, len(u))
	for i, u := range u {
		users[i] = toDomain(u)
	}
	return users, nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int64) (model.User, error) {
	var u User
	const q = `select id, name, tg_id, party_id, created_at, last_active_at from users where id = $1`
	err := r.pool.QueryRow(ctx, q, id).Scan(
		&u.ID,
		&u.Name,
		&u.TgID,
		&u.PartyID,
		&u.CreatedAt,
		&u.LastActiveAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, user.ErrNotFound
		}
		fmt.Fprintf(os.Stderr, "Error %v \n", err)
		return model.User{}, err
	}
	user := toDomain(u)
	return user, nil
}

func toDomain(u User) model.User {
	return model.User{
		ID:      u.ID,
		Name:    u.Name,
		TgID:    u.TgID,
		PartyID: u.PartyID,
	}
}
