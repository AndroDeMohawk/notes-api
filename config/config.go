package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	Port        string
	Env         string        // dev, production
	DatabaseURL string        // DSN для подключения к Postgres
	JWTSecret   string        // Секретный ключ для подписи токенов
	JWTTTL      time.Duration // Время жизни JWT токена
}

func LoadConfig() (*Config, error) {
	// Игнорируем ошибку, если .env файла нет (например, в Docker или на продакшене)
	_ = godotenv.Load(".env")

	port := getEnv("PORT", "8083")
	version := getEnv("APP_VERSION", "dev")
	env := getEnv("APP_ENV", "development")

	// Параметры базы данных
	dbURL := getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/notes_db?sslmode=disable")

	// Безопасность
	jwtSecret := getEnv("JWT_SECRET", "default-super-secret-key-change-me")

	// TTL токена (парсим строку в time.Duration или берём дефолтные 24h)
	jwtTTLStr := getEnv("JWT_TTL", "24h")
	jwtTTL, err := time.ParseDuration(jwtTTLStr)
	if err != nil {
		jwtTTL = 24 * time.Hour
	}

	return &Config{
		Version:     version,
		Port:        port,
		Env:         env,
		DatabaseURL: dbURL,
		JWTSecret:   jwtSecret,
		JWTTTL:      jwtTTL,
	}, nil
}

// Вспомогательная функция для чтения переменных с фоллбэком
func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return defaultValue
}
