package repository

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	ID           int64       `db:"id"`
	Name         string      `db:"name"`
	TgID         int64       `db:"tg_id"`
	PartyID      *int        `db:"party_id"`
	CreatedAt    pgtype.Date `db:"created_at"`
	LastActiveAt pgtype.Date `db:"last_active_at"`
}
