package http

import "Test2/internal/user/model"

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	TgID  int64  `json:"tg_id" binding:"required"`
}

func (r CreateUserRequest) ToModel() model.CreateUser {
	return model.CreateUser{
		Name:  r.Name,
		TgID:  r.TgID,
		Phone: r.Phone,
	}
}
