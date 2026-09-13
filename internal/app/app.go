package app

import (
	"log"
	"net/http"

	"github.com/AndroDeMohawk/notes-api/internal/config"
	"github.com/go-chi/chi/v5"
)

func Run() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()
	httpServer := &http.Server{
		Addr:    config.Port,
		Handler: router,
	}
	httpServer.ListenAndServe()
}
