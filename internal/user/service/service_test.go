// internal/user/service/service_test.go
package service_test

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"Test2/internal/user/model"
	"Test2/internal/user/service"
	"Test2/internal/user/service/mocks"
)

// silentLogger — тихий логгер, чтобы тесты не сыпали в stdout.
// В тестах на логику логи нам не нужны — только чтобы *slog.Logger был валидным.
func silentLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

// ---------- GetUsers ----------

func TestUserService_GetUsers_OK(t *testing.T) {
	want := []model.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	repo := mocks.NewRepository(t)
	repo.EXPECT().
		GetUsers(mock.Anything).
		Return(want, nil).
		Once()

	svc := service.New(repo, silentLogger())

	got, err := svc.GetUsers(context.Background())

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestUserService_GetUsers_Error(t *testing.T) {
	wantErr := errors.New("db is down")

	repo := mocks.NewRepository(t)
	repo.EXPECT().
		GetUsers(mock.Anything).
		Return(nil, wantErr).
		Once()

	svc := service.New(repo, silentLogger())

	got, err := svc.GetUsers(context.Background())

	require.ErrorIs(t, err, wantErr)
	require.Nil(t, got)
}

// ---------- GetUser ----------

func TestUserService_GetUser_OK(t *testing.T) {
	want := &model.User{ID: 42, Name: "Alice"}

	repo := mocks.NewRepository(t)
	repo.EXPECT().
		GetUser(mock.Anything, int64(42)). // проверяем, что сервис передал правильный id
		Return(want, nil).
		Once()

	svc := service.New(repo, silentLogger())

	got, err := svc.GetUser(context.Background(), 42)

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestUserService_GetUser_Error(t *testing.T) {
	wantErr := errors.New("not found")

	repo := mocks.NewRepository(t)
	repo.EXPECT().
		GetUser(mock.Anything, int64(1)).
		Return(nil, wantErr).
		Once()

	svc := service.New(repo, silentLogger())

	got, err := svc.GetUser(context.Background(), 1)

	require.ErrorIs(t, err, wantErr)
	require.Nil(t, got)
}

// ---------- GetUserByTgID ----------

func TestUserService_GetUserByTgID_OK(t *testing.T) {
	want := &model.User{ID: 7, Name: "Alice"}

	repo := mocks.NewRepository(t)
	repo.EXPECT().
		GetUserByTgID(mock.Anything, int64(123456)).
		Return(want, nil).
		Once()

	svc := service.New(repo, silentLogger())

	got, err := svc.GetUserByTgID(context.Background(), 123456)

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestUserService_GetUserByTgID_Error(t *testing.T) {
	wantErr := errors.New("db error")

	repo := mocks.NewRepository(t)
	repo.EXPECT().
		GetUserByTgID(mock.Anything, int64(1)).
		Return(nil, wantErr).
		Once()

	svc := service.New(repo, silentLogger())

	got, err := svc.GetUserByTgID(context.Background(), 1)

	require.ErrorIs(t, err, wantErr)
	require.Nil(t, got)
}

// ---------- CreateUser ----------

func TestUserService_CreateUser_OK(t *testing.T) {
	input := model.CreateUser{Name: "Alice", TgID: 123456}
	want := &model.User{ID: 1, Name: "Alice", TgID: 123456}

	repo := mocks.NewRepository(t)
	repo.EXPECT().
		CreateUser(mock.Anything, input). // точное совпадение входного DTO
		Return(want, nil).
		Once()

	svc := service.New(repo, silentLogger())

	got, err := svc.CreateUser(context.Background(), input)

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestUserService_CreateUser_Error(t *testing.T) {
	wantErr := errors.New("duplicate key")

	repo := mocks.NewRepository(t)
	repo.EXPECT().
		CreateUser(mock.Anything, mock.Anything).
		Return(nil, wantErr).
		Once()

	svc := service.New(repo, silentLogger())

	got, err := svc.CreateUser(context.Background(), model.CreateUser{Name: "Alice"})

	require.ErrorIs(t, err, wantErr)
	require.Nil(t, got)
}
