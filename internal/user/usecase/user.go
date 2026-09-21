package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AndroDeMohawk/notes-api/internal/user/domain"
	"github.com/AndroDeMohawk/notes-api/internal/user/dto"
	"github.com/AndroDeMohawk/notes-api/pkg/auth"
)

type TokenManager interface {
	GenerateToken(userID int64, ttl time.Duration) (string, error)
}

type UserUseCase struct {
	userRepo     domain.UserRepository
	tokenManager TokenManager
	tokenTTL     time.Duration
}

func NewUserUseCase(repo domain.UserRepository, tokenManager TokenManager, tokenTTL time.Duration) *UserUseCase {
	return &UserUseCase{
		userRepo:     repo,
		tokenManager: tokenManager,
		tokenTTL:     tokenTTL,
	}
}

func (uc *UserUseCase) Register(ctx context.Context, input dto.RegisterInput) (*domain.User, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}
	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &domain.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: hashedPassword,
		Role:         domain.RoleUser,
	}

	createdUser, err := uc.userRepo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	createdUser.PasswordHash = ""
	return createdUser, nil
}

func (uc *UserUseCase) Login(ctx context.Context, input dto.LoginInput) (string, error) {
	if err := input.Validate(); err != nil {
		return "", fmt.Errorf("validation error: %w", err)
	}

	user, err := uc.userRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !auth.CheckPasswordHash(input.Password, user.PasswordHash) {
		return "", errors.New("invalid email or password")
	}

	// Генерируем токен
	token, err := uc.tokenManager.GenerateToken(user.ID, uc.tokenTTL)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func (uc *UserUseCase) GetProfile(ctx context.Context, userID int64) (*domain.User, error) {
	if userID <= 0 {
		return nil, errors.New("unauthorized: invalid user id")
	}

	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user.PasswordHash = ""
	return user, nil
}

func (uc *UserUseCase) UpdateProfile(ctx context.Context, input dto.UpdateProfileInput) (*domain.User, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	user, err := uc.userRepo.GetByID(ctx, input.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	user.Username = input.Username
	user.Email = input.Email

	updatedUser, err := uc.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user profile: %w", err)
	}

	updatedUser.PasswordHash = ""
	return updatedUser, nil
}
