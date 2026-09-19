package info

import (
	"net/http"

	"github.com/AndroDeMohawk/notes-api/config"
	"github.com/AndroDeMohawk/notes-api/pkg/response"
)

type Handler struct {
	version string
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response.SendJson(w, map[string]string{"status": "ok"}, http.StatusOK)
}

func (h *Handler) AppVersion(w http.ResponseWriter, r *http.Request) {
	response.SendJson(w, map[string]string{"version": h.version}, http.StatusOK)
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		version: cfg.Version,
	}
}
