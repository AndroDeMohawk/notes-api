package app

import (
	"log"
	"net/http"

	"github.com/AndroDeMohawk/notes-api/internal/config"
	"github.com/AndroDeMohawk/notes-api/internal/httpapi"
	"github.com/go-chi/chi/v5"
)

func Run() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()
	httpapi.RegisterRoutes(router)
	httpServer := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}
	httpServer.ListenAndServe()
}
