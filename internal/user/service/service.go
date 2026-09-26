package service

import (
	"Test2/internal/user/model"
	"context"
	"log/slog"
)

type UserService struct {
	repo   Repository
	logger *slog.Logger
}

func New(repo Repository, log *slog.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: log,
	}
}

func (s *UserService) GetUsers(ctx context.Context) ([]model.User, error) {
	response, err := s.repo.GetUsers(ctx)

	if err != nil {
		s.logger.ErrorContext(ctx, "get user failed", slog.Any("err", err))
		return response, err
	}
	return response, nil
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	response, err := s.repo.GetUser(ctx, id)

	if err != nil {
		s.logger.ErrorContext(ctx, "get user failed", slog.Any("err", err))
		return response, err
	}
	return response, nil
}

func (s *UserService) GetUserByTgID(ctx context.Context, tgID int64) (*model.User, error) {
	response, err := s.repo.GetUserByTgID(ctx, tgID)

	if err != nil {
		s.logger.ErrorContext(ctx, "get user failed", slog.Any("err", err))
		return response, err
	}
	return response, nil
}

func (s *UserService) CreateUser(ctx context.Context, u model.CreateUser) (*model.User, error) {
	user, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		s.logger.ErrorContext(ctx, "create user failed", slog.Any("err", err))
		return nil, err
	}
	return user, nil
}

//go:generate mockery --name=Repository --output=./mocks --outpkg=mocks --with-expecter
type Repository interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUser(ctx context.Context, id int64) (*model.User, error)
	GetUserByTgID(ctx context.Context, tgID int64) (*model.User, error)
	CreateUser(ctx context.Context, u model.CreateUser) (*model.User, error)
}
