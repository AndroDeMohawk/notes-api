package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Version string
	Port    string
}

type Db struct {
	Dsn string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	return &Config{
		App: AppConfig{
			Version: os.Getenv("APP_VERSION"),
			Port:    os.Getenv("APP_PORT"),
		},
	}, nil
}
