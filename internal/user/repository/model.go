package repository

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	ID           int64
	Name         string
	CreatedAt    pgtype.Date
	LastActiveAt pgtype.Date
}
