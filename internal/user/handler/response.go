package handler

import "Test2/internal/user/model"

type UsersListResponse struct {
	Users []model.User
}

type UserResponse struct {
	User model.User
}
