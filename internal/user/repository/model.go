package repository

import (
	"time"
)

type User struct {
	ID           int64      `db:"id"`
	Name         string     `db:"name"`
	TgID         int64      `db:"tg_id"`
	PhoneNumber  string     `db:"phone"`
	CreatedAt    *time.Time `db:"created_at"`
	LastActiveAt *time.Time `db:"last_active_at"`
}
