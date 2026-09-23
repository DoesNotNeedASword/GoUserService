package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	GRPCAddr       string
	DatabaseURL    string
	MigrationsPath string
	RunMigrations  bool
	LogLevel       string
}

func Load() Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v \n", err)
		os.Exit(1)
	}

	return Config{
		GRPCAddr:       envString("CORE_GRPC_ADDR", ":50051"),
		DatabaseURL:    envString("DATABASE_URL", ""),
		MigrationsPath: envString("MIGRATIONS_PATH", "file://migrations"),
		LogLevel:       envString("LOG_LEVEL", "info"),
	}
}

func envString(name string, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		fmt.Fprintf(os.Stderr, "URL not set")
		return defaultValue
	}
	return value
}
