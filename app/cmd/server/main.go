package main

import (
	"log"

	"github.com/egoriyNovikov/gridwar/app/internal/config"
	"github.com/egoriyNovikov/gridwar/app/internal/server"
	"github.com/egoriyNovikov/gridwar/app/internal/ws"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	hub := ws.NewHub()
	go hub.Run()

	srv := server.New(cfg, hub)
	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
