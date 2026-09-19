package repository

import (
	"github.com/AndroDeMohawk/notes-api/i"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

//func (r *UserRepository) Create(ctx context.Context, user domain.User) error {
//
//}
