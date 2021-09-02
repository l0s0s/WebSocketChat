package client

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// Client is how all clients look like.
type Client struct {
	Socket   *websocket.Conn
	Send     chan *Message
	Room     *Room
	userData map[string]interface{}
}

func (c *Client) read() {
	defer c.Socket.Close()

	for {
		var msg *Message

		err := c.Socket.ReadJSON(&msg)
		if err != nil {
			return
		}

		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		c.Room.Forward <- msg
	}
}

func (c *Client) write() {
	defer c.Socket.Close()

	for msg := range c.Send {
		err := c.Socket.WriteJSON(msg)
		if err != nil {
			break
		}
	}
}
