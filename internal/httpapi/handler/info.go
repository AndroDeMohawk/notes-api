package handler

import (
	"net/http"

	"github.com/AndroDeMohawk/notes-api/internal/config"
	"github.com/AndroDeMohawk/notes-api/pkg/response"
)

type Handler struct {
	Version string
}

func (h *Handler) HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.SendJson(w, map[string]string{"status": "ok"}, http.StatusOK)
	}
}

func (h *Handler) AppVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		response.SendJson(w, map[string]string{"version": h.Version}, http.StatusOK)
	}
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		Version: cfg.Version,
	}
}
