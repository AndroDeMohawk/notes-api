package handler

import (
	"net/http"
	"os"

	"github.com/AndroDeMohawk/notes-api/pkg/response"
)

type Handler struct {
}

func (h *Handler) HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.SendJson(w, `{"status" : "OK"}`, http.StatusOK)
	}
}

func (h *Handler) Version() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		AppV := os.Getenv("APP_VERSION")
		response.SendJson(w, AppV, http.StatusOK)
	}
}
