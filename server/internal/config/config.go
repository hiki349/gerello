package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ConnStrPostgres string
	MaxAttempts     int
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	connStrPostgres := os.Getenv("POSTGRES_CONN")
	maxAttempts, err := strconv.Atoi(os.Getenv("MAX_ATTEMPTS"))
	if err != nil {
		return nil, err
	}

	return &Config{
		ConnStrPostgres: connStrPostgres,
		MaxAttempts: maxAttempts,
	}, nil
}