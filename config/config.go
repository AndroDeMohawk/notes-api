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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "dev"
	}
	return &Config{
		Version: version,
		Port:    port,
	}, nil
}
