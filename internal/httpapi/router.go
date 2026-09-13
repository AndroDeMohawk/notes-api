package httpapi

import "github.com/go-chi/chi/v5"

func RegisterRoutes(router *chi.Mux) {
	router.Route("/api", func(r chi.Router) {

	})
}
