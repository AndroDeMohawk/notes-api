package repository

import (
	"context"
	"fmt"

	"github.com/AndroDeMohawk/notes-api/internal/database/db"
	"github.com/AndroDeMohawk/notes-api/internal/task/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	q *db.Queries
}

func NewPostgresRepository(pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		q: db.New(pool),
	}
}

func (r *TaskRepository) Create(ctx context.Context, t *domain.Task) (*domain.Task, error) {
	create, err := r.q.CreateTask(ctx, db.CreateTaskParams{
		UserID:      t.UserID,
		Title:       t.Title,
		Description: t.Description,
		Status:      string(t.Status),
	})
	if err != nil {
		return nil, fmt.Errorf("error creating task: %w", err)
	}

	task := toDomainTask(create)
	return &task, nil
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, id int64) (*domain.Task, error) {
	t, err := r.q.GetTaskByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting task: %w", err)
	}

	task := toDomainTask(t)
	return &task, nil
}

func (r *TaskRepository) GetListTasksByUserID(ctx context.Context, userId int64) ([]domain.Task, error) {
	dbTasks, err := r.q.ListTasksByUserID(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("error getting tasks: %w", err)
	}

	tasks := make([]domain.Task, 0, len(dbTasks))
	for _, t := range dbTasks {
		tasks = append(tasks, toDomainTask(t))
	}

	return tasks, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, t *domain.Task, userId int64) (*domain.Task, error) {
	update, err := r.q.UpdateTask(ctx, db.UpdateTaskParams{
		ID:          t.ID,
		UserID:      userId,
		Title:       t.Title,
		Description: t.Description,
		Status:      string(t.Status),
	})
	if err != nil {
		return nil, fmt.Errorf("error updating task: %w", err)
	}

	task := toDomainTask(update)
	return &task, nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id int64, userId int64) error {
	err := r.q.DeleteTask(ctx, db.DeleteTaskParams{
		ID:     id,
		UserID: userId,
	})
	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}
	return nil
}

// Возвращаем структуру domain.Task по значению
func toDomainTask(t db.Task) domain.Task {
	return domain.Task{
		ID:          t.ID,
		UserID:      t.UserID,
		Title:       t.Title,
		Description: t.Description,
		Status:      domain.Status(t.Status),
		CreatedAt:   t.CreatedAt.Time,
	}
}
