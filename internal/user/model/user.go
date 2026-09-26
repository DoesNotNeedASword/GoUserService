package model

import "time"

type User struct {
	ID           int64
	Name         string
	TgID         int64
	Phone        string
	CreatedAt    *time.Time
	LastActiveAt *time.Time
}

type CreateUser struct {
	Name  string
	TgID  int64
	Phone string
}
