package app

import (
	"net/http"

	"github.com/AndroDeMohawk/notes-api/internal/info"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	UserHandler http.Handler //потом переименую как user.Handler
	InfoHandler info.Handler
}

func RegisterRoutes(h *Handler) http.Handler {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", h.InfoHandler.HealthCheck)
		r.Get("/version", h.InfoHandler.AppVersion)
	})
	return r
}
