package dto

import (
	"errors"
	"strings"
)

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfileInput struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (i *RegisterInput) Validate() error {
	i.Username = strings.TrimSpace(i.Username)
	i.Email = strings.TrimSpace(strings.ToLower(i.Email))

	if i.Username == "" {
		return errors.New("username cannot be empty")
	}
	if i.Email == "" || !strings.Contains(i.Email, "@") {
		return errors.New("invalid email address")
	}
	if len(i.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}

func (i *LoginInput) Validate() error {
	i.Email = strings.TrimSpace(strings.ToLower(i.Email))

	if i.Email == "" {
		return errors.New("email is required")
	}
	if i.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (i *UpdateProfileInput) Validate() error {
	i.Username = strings.TrimSpace(i.Username)
	i.Email = strings.TrimSpace(strings.ToLower(i.Email))

	if i.UserID <= 0 {
		return errors.New("unauthorized: invalid user id")
	}
	if i.Username == "" {
		return errors.New("username cannot be empty")
	}
	if i.Email == "" || !strings.Contains(i.Email, "@") {
		return errors.New("invalid email address")
	}
	return nil
}
