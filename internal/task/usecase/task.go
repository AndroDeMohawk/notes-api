package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/AndroDeMohawk/notes-api/internal/task/domain"
	"github.com/AndroDeMohawk/notes-api/internal/task/dto"
)

type TaskUС struct {
	taskRepo domain.TaskRepository
}

func NewTaskUseCase(repo domain.TaskRepository) *TaskUС {
	return &TaskUС{taskRepo: repo}
}

func (uc *TaskUС) CreateTask(ctx context.Context, cti *dto.CreateTaskInput) (*domain.Task, error) {
	err := cti.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate input error: %w", err)
	}
	task := &domain.Task{
		Title:       cti.Title,
		Description: cti.Description,
		UserID:      cti.UserID,
		Status:      domain.StatusNEW,
	}

	newTask, err := uc.taskRepo.Create(ctx, task)
	if err != nil {
		return nil, fmt.Errorf("create task error: %w", err)
	}
	return newTask, nil
}

func (uc *TaskUС) DeleteTask(ctx context.Context, id, userID int64) error {
	if id <= 0 || userID <= 0 {
		return errors.New("invalid input")
	}
	err := uc.taskRepo.DeleteTask(ctx, id, userID)
	if err != nil {
		return fmt.Errorf("delete task error: %w", err)
	}
	return nil
}

func (uc *TaskUС) UpdateTask(ctx context.Context, uti *dto.UpdateTaskInput) (*domain.Task, error) {
	err := uti.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate input error: %w", err)
	}
	task := &domain.Task{
		ID:          uti.ID,
		UserID:      uti.UserID,
		Title:       uti.Title,
		Description: uti.Description,
		Status:      domain.Status(uti.Status),
	}
	updatedTask, err := uc.taskRepo.Create(ctx, task)
	if err != nil {
		return nil, fmt.Errorf("update task error: %w", err)
	}
	return updatedTask, nil
}

func (uc *TaskUС) GetUsersTaskList(ctx context.Context, userID int64) ([]domain.Task, error) {
	if userID <= 0 {
		return nil, errors.New("invalid input")
	}
	tasks, err := uc.taskRepo.GetListTasksByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get tasks list error: %w", err)
	}
	return tasks, nil
}

func (uc *TaskUС) GetTaskById(ctx context.Context, id, userId int64) (*domain.Task, error) {
	if id <= 0 || userId <= 0 {
		return nil, errors.New("invalid input")
	}
	task, err := uc.taskRepo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get task error: %w", err)
	}
	if task.UserID != userId {
		return nil, errors.New("invalid user, access denied")
	}
	return task, nil
}
