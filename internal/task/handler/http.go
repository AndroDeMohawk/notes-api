package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AndroDeMohawk/notes-api/internal/task/dto"
	"github.com/AndroDeMohawk/notes-api/internal/task/usecase"
	"github.com/AndroDeMohawk/notes-api/pkg/middlewares"
	"github.com/AndroDeMohawk/notes-api/pkg/response"
)

type TaskHandler struct {
	taskUC *usecase.TaskUС
}

func NewTaskHandler(uc *usecase.TaskUС) *TaskHandler {
	return &TaskHandler{
		taskUC: uc,
	}
}

// POST /tasks
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		response.SendJson(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var input dto.CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.SendJson(w, "invalid request body", http.StatusBadRequest)
		return
	}

	input.UserID = userID

	task, err := h.taskUC.CreateTask(r.Context(), &input)
	if err != nil {
		response.SendJson(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.SendJson(w, task, http.StatusCreated)
}

// GET /tasks
func (h *TaskHandler) GetUserTasks(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		response.SendJson(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	tasks, err := h.taskUC.GetUsersTaskList(r.Context(), userID)
	if err != nil {
		response.SendJson(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.SendJson(w, tasks, http.StatusOK)
}

// GET /tasks/{id}
func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		response.SendJson(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Извлекаем ID из URL (стандартный r.PathValue в Go 1.22+)
	taskID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || taskID <= 0 {
		response.SendJson(w, "invalid task id", http.StatusBadRequest)
		return
	}

	task, err := h.taskUC.GetTaskById(r.Context(), taskID, userID)
	if err != nil {
		response.SendJson(w, err.Error(), http.StatusNotFound)
		return
	}

	response.SendJson(w, task, http.StatusOK)
}

// PUT /tasks/{id}
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		response.SendJson(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	taskID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || taskID <= 0 {
		response.SendJson(w, "invalid task id", http.StatusBadRequest)
		return
	}

	var input dto.UpdateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.SendJson(w, "invalid request body", http.StatusBadRequest)
		return
	}

	input.ID = taskID
	input.UserID = userID

	task, err := h.taskUC.UpdateTask(r.Context(), &input)
	if err != nil {
		response.SendJson(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.SendJson(w, task, http.StatusOK)
}

// DELETE /tasks/{id}
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		response.SendJson(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	taskID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || taskID <= 0 {
		response.SendJson(w, "invalid task id", http.StatusBadRequest)
		return
	}

	if err := h.taskUC.DeleteTask(r.Context(), taskID, userID); err != nil {
		response.SendJson(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.SendJson(w, map[string]string{"message": "task deleted successfully"}, http.StatusOK)
}
