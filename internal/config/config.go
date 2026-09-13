package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Version string
	Port    string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		Version: os.Getenv("APP_VERSION"),
		Port:    os.Getenv("PORT"),
	}, nil
}
