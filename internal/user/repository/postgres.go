package repository

import (
	"context"
	"fmt"

	"github.com/AndroDeMohawk/notes-api/internal/database/db"
	"github.com/AndroDeMohawk/notes-api/internal/user/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	q *db.Queries
}

func NewPostgresRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		q: db.New(pool),
	}
}

func (r *UserRepository) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	create, err := r.q.CreateUser(ctx, db.CreateUserParams{
		Username:     u.Username,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Role:         string(u.Role),
	})
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	user := toDomainUser(create)
	return &user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) { // FIX: Переименовано в GetByID
	u, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	user := toDomainUser(u)
	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	user := toDomainUser(u)
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, u *domain.User) (*domain.User, error) {
	update, err := r.q.UpdateUser(ctx, db.UpdateUserParams{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	user := toDomainUser(update)
	return &user, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	err := r.q.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

// Теперь возвращает значение domain.User, а не указатель
func toDomainUser(u db.User) domain.User {
	return domain.User{
		ID:           u.ID,
		Username:     u.Username,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Role:         domain.Role(u.Role),
		CreatedAt:    u.CreatedAt.Time,
	}
}
