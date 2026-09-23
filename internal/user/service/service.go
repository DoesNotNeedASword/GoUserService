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

func (s *UserService) GetUsers(ctx context.Context) ([]model.User, error) {
	response, err := s.repo.GetUsers(ctx)

	if err != nil {
		fmt.Println("Error", err)
		return response, err
	}
	return response, nil
}

type Repository interface {
	GetUsers(ctx context.Context) ([]model.User, error)
}
