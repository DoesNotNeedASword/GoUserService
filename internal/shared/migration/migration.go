package migration

import (
	"errors"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
)

func Run(databaseURL, migrationPath string, logger *slog.Logger) error {
	m, err := migrate.New(databaseURL, migrationPath)
	if err != nil { return err }
	defer m.Close()

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("up to date")
			return nil
		}
		return err
	}
	logger.Info("migrate up")
	return nil
}
