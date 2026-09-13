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
		response.SendJson(w, `{"status" : "OK"}`, http.StatusOK)
	}
}

func (h *Handler) AppVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		AppV := Handler{}
		response.SendJson(w, AppV, http.StatusOK)
	}
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		Version: cfg.Version,
	}
}
