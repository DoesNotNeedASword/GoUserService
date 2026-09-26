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

func (r *UserRepository) GetUser(ctx context.Context, id int64) (*model.User, error) {
	var u User
	const q = `select id, name, tg_id, created_at, last_active_at from users where id = $1`
	err := r.pool.QueryRow(ctx, q, id).Scan(
		&u.ID,
		&u.Name,
		&u.TgID,
		&u.CreatedAt,
		&u.LastActiveAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		return nil, err
	}
	user := toDomain(u)
	return &user, nil
}

func (r *UserRepository) GetUserByTgID(ctx context.Context, tgID int64) (*model.User, error) {
	var u User
	const q = `select id, name, tg_id, created_at, last_active_at from users where tg_id = $1`
	err := r.pool.QueryRow(ctx, q, tgID).Scan(
		&u.ID,
		&u.Name,
		&u.TgID,
		&u.CreatedAt,
		&u.LastActiveAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		fmt.Fprintf(os.Stderr, "Error %v \n", err)
		return nil, err
	}
	user := toDomain(u)
	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, u model.CreateUser) (*model.User, error) {
	const q = `
        INSERT INTO users (name, tg_id, phone)
        VALUES ($1, $2, $3)
        RETURNING id, name, tg_id, phone, created_at, last_active_at`

	var user model.User
	err := r.pool.QueryRow(ctx, q, u.Name, u.TgID, u.Phone).Scan(
		&user.ID,
		&user.Name,
		&user.TgID,
		&user.Phone,
		&user.CreatedAt,
		&user.LastActiveAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return &user, nil
}

func toDomain(u User) model.User {
	return model.User{
		ID:   u.ID,
		Name: u.Name,
		TgID: u.TgID,
	}
}
