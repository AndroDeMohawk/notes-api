package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	taskHandler "github.com/AndroDeMohawk/notes-api/internal/task/handler"
	userHandler "github.com/AndroDeMohawk/notes-api/internal/user/handler"
)

// Handlers объединяет контроллеры и мидлвари приложения
type Handlers struct {
	User           *userHandler.UserHandler
	Task           *taskHandler.TaskHandler
	AuthMiddleware func(http.Handler) http.Handler
}

func RegisterRoutes(h *Handlers) http.Handler {
	r := chi.NewRouter()

	// Базовые системные мидлвари Chi
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	// Публичные маршруты
	r.Post("/auth/register", h.User.Register)
	r.Post("/auth/login", h.User.Login)

	// Защищенные маршруты (с мидлварью авторизации)
	r.Group(func(r chi.Router) {
		r.Use(h.AuthMiddleware)

		// Профиль пользователя
		r.Get("/users/me", h.User.GetProfile)
		r.Put("/users/me", h.User.UpdateProfile)

		// Задачи
		r.Route("/tasks", func(r chi.Router) {
			r.Post("/", h.Task.CreateTask)
			r.Get("/", h.Task.GetUserTasks)
			r.Get("/{id}", h.Task.GetTaskByID)
			r.Put("/{id}", h.Task.UpdateTask)
			r.Delete("/{id}", h.Task.DeleteTask)
		})
	})

	return r
}
