package main

import (
	"errors"
	"os"
)

type Config struct {
	DatabaseURL string
}

func parseConfigFromEnv() (*Config, error) {
	DatabaseURL := os.Getenv("DATABASE_URL")
	if DatabaseURL == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	return &Config{
		DatabaseURL: DatabaseURL,
	}, nil
}
