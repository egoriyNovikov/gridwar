package ws

import "github.com/gorilla/websocket"

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	send chan []byte
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{
		Hub:  hub,
		Conn: conn,
		send: make(chan []byte),
	}
}
