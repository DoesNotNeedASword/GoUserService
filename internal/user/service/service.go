package service

import (
	"Test2/internal/user/model"
	"context"
	"fmt"
)

type UserService struct {
	repo Repository
}

func New(repo Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUsers(ctx context.Context) (model.User, error) {

	response, err := s.repo.GetUsers(ctx)

	user := model.User{ID: response.ID, Name: response.Name}

	if err != nil {
		fmt.Println("Error", err)
		return user, err
	}
	fmt.Println(user.ID, user.Name)
	return model.User{ID: user.ID, Name: user.Name}, nil
}

type Repository interface {
	GetUsers(ctx context.Context) (model.User, error)
}
