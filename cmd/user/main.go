package main

import (
	"Test2/internal/shared/config"
	"Test2/internal/shared/database"
	"Test2/internal/shared/logger"
	"Test2/internal/shared/migration"
	"Test2/internal/user/handler"
	"Test2/internal/user/repository"
	"Test2/internal/user/service"
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	logger := logger.New(cfg.LogLevel)
	err := run(cfg, logger)
	migration.Run(cfg.DatabaseURL, cfg.MigrationsPath, logger)
	if err != nil {
		log.Fatalln(err)
	}

}

func run(cfg config.Config, log *slog.Logger) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	router := gin.Default()
	pool, err := database.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()
	repo := repository.New(pool)
	service := service.New(repo)
	handler := handler.New(service)
	router.GET("/users", handler.GetUsers)

	router.Run("localhost:8080")

	return nil
}
