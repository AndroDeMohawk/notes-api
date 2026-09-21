package dto

import (
	"errors"
)

type CreateTaskInput struct {
	UserID      int64  `json:"userID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (cti *CreateTaskInput) Validate() error {
	if cti.UserID == 0 {
		return errors.New("UserID is empty")
	}
	if cti.Title == "" {
		return errors.New("Title is empty")
	}
	return nil
}

type UpdateTaskInput struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (uti *UpdateTaskInput) Validate() error {
	if uti.UserID == 0 {
		return errors.New("UserID is empty")
	}
	if uti.Title == "" {
		return errors.New("Title is empty")
	}
	if uti.Status == "" {
		return errors.New("Status is empty")
	}
	return nil
}
