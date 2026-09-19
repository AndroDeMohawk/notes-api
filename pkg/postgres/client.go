package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewClient(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("error to parse dsn", err)
	}
	config.MaxConns = 10
	config.MaxConnLifetime = time.Hour
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error to create new pool", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("error to ping pool", err)
	}
	return pool, nil
}
