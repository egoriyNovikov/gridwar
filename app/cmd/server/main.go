package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/egoriyNovikov/gridwar/app/internal/config"
	httpHandler "github.com/egoriyNovikov/gridwar/app/internal/http"
	"github.com/egoriyNovikov/gridwar/app/internal/ws"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	hub := ws.NewHub()
	go hub.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		httpHandler.HandleWS(w, r, hub)
	})

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
