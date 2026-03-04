package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/egoriyNovikov/gridwar/app/internal/config"
	"github.com/egoriyNovikov/gridwar/app/internal/handler"
	"github.com/egoriyNovikov/gridwar/app/internal/ws"
)

type Server struct {
	addr string
	mux  *http.ServeMux
}

func New(cfg *config.Config, hub *ws.Hub) *Server {
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	mux := http.NewServeMux()

	staticDir := cfg.Server.StaticDir
	if staticDir == "" {
		staticDir = "web"
	}
	mux.Handle("/", http.FileServer(http.Dir(staticDir)))
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleWS(w, r, hub)
	})
	return &Server{addr: addr, mux: mux}
}

func (s *Server) Run() error {
	log.Printf("Starting server on %s", s.addr)
	return http.ListenAndServe(s.addr, s.mux)
}
