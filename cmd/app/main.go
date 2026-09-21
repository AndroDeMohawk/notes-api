package main

import (
	"log"

	"github.com/AndroDeMohawk/notes-api/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Fatal app error: %v", err)
	}
}
