package http

import (
	"log"
	"net/http"

	"github.com/egoriyNovikov/gridwar/app/internal/ws"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWS(w http.ResponseWriter, r *http.Request, hub *ws.Hub) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	client := ws.NewClient(hub, conn)
	hub.Register(client)
	go client.ReadPump()
	go client.WritePump()
}
