package client

import (
	"github.com/gorilla/websocket"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

type Client struct {
	Socket *websocket.Conn
	Send   chan []byte
	Room   *Room
}

func (c *Client) read() {
	defer c.Socket.Close()

	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		c.Room.Forward <- msg
	}
}

func (c *Client) write() {
	defer c.Socket.Close()

	for msg := range c.Send {
		err := c.Socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
