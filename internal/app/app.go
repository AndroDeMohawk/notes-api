package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func Run() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := chi.NewRouter()
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	httpServer.ListenAndServe()
}
