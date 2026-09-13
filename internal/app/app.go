package app

import (
	"log"
	"net/http"

	"github.com/AndroDeMohawk/notes-api/internal/config"
	"github.com/AndroDeMohawk/notes-api/internal/httpapi"
	"github.com/AndroDeMohawk/notes-api/internal/httpapi/handler"
)

func Run() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	h := handler.NewHandler(cfg)
	router := httpapi.NewRouter(h)
	httpServer := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
