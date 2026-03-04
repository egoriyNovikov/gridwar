package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

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

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister(c)
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message: %v", err)
			return
		}
		if string(message) == "ping" {
			c.Conn.WriteMessage(websocket.TextMessage, []byte("pong"))
			continue
		}
		c.Hub.Broadcast(message)
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for message := range c.send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("WriteMessage: %v", err)
			return
		}
	}
}
