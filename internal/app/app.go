package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/AndroDeMohawk/notes-api/config"
	taskHandler "github.com/AndroDeMohawk/notes-api/internal/task/handler"
	taskRepo "github.com/AndroDeMohawk/notes-api/internal/task/repository"
	taskUC "github.com/AndroDeMohawk/notes-api/internal/task/usecase"
	userHandler "github.com/AndroDeMohawk/notes-api/internal/user/handler"
	userRepo "github.com/AndroDeMohawk/notes-api/internal/user/repository"
	userUC "github.com/AndroDeMohawk/notes-api/internal/user/usecase"
	"github.com/AndroDeMohawk/notes-api/pkg/auth"
	"github.com/AndroDeMohawk/notes-api/pkg/middlewares"
)

func Run() error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to create db pool: %w", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("db ping failed: %w", err)
	}
	logger.Info("connected to PostgreSQL successfully")

	jwtManager := auth.NewJWTManager(cfg.JWTSecret)
	tokenTTL := 24 * time.Hour

	uRepo := userRepo.NewPostgresRepository(pool)
	uUseCase := userUC.NewUserUseCase(uRepo, jwtManager, tokenTTL)
	uHandler := userHandler.NewUserHandler(uUseCase)

	tRepo := taskRepo.NewPostgresRepository(pool)
	tUseCase := taskUC.NewTaskUseCase(tRepo)
	tHandler := taskHandler.NewTaskHandler(tUseCase)

	h := &Handlers{
		User:           uHandler,
		Task:           tHandler,
		AuthMiddleware: middlewares.Authorize(jwtManager),
	}

	router := RegisterRoutes(h)

	httpServer := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("starting server", zap.String("port", cfg.Port))

		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("server error: %w", err)
	case <-ctx.Done():
		logger.Info("shutting down server...")
	}

	shutDownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutDownCtx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}

	logger.Info("server shutdown successfully")
	return nil
}
