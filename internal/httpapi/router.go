package httpapi

import (
	"net/http"

	"github.com/AndroDeMohawk/notes-api/internal/httpapi/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter(info *handler.Handler) http.Handler {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", info.HealthCheck())
		r.Get("/version", info.AppVersion())
	})
	return r
}
