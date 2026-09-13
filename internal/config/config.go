package config

import (
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
		Version: "dev",
		Port:    ":8080",
	}, nil
}
