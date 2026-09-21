package domain

import (
	"context"
	"time"
)

type Status string

const (
	StatusNEW     Status = "NEW"
	StatusDONE    Status = "DONE"
	StatusPENDING Status = "PENDING"
)

type Task struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskRepository interface {
	Create(ctx context.Context, task *Task) (*Task, error)
	GetTaskByID(ctx context.Context, id int64) (*Task, error)
	GetListTasksByUserID(ctx context.Context, userID int64) ([]Task, error)
	UpdateTask(ctx context.Context, task *Task, userID int64) (*Task, error)
	DeleteTask(ctx context.Context, id int64, userID int64) error
}
