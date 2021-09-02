package client

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/l0s0s/WebSocketChat/trace"
	"github.com/stretchr/objx"
)

// Room is how all rooms look like.
type Room struct {
	Forward chan *Message
	Join    chan *Client
	Leave   chan *Client
	Clients map[*Client]bool
	Tracer  trace.Tracer
}

// NewRoom makes a new room.
func NewRoom() *Room {
	return &Room{
		Forward: make(chan *Message),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Clients: make(map[*Client]bool),
		Tracer:  trace.Off(),
	}
}

// Run is run some room.
func (r *Room) Run() {
	for {
		select {
		case client := <-r.Join:
			r.Clients[client] = true
			r.Tracer.Trace("New client joined")
		case client := <-r.Leave:
			delete(r.Clients, client)
			close(client.Send)
			r.Tracer.Trace("Client left")
		case msg := <-r.Forward:
			r.Tracer.Trace("Message received: ", string(msg.Message))

			for client := range r.Clients {
				client.Send <- msg
			}

			r.Tracer.Trace(" -- sent to client")
		}
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		r.Tracer.Trace(err)
		return
	}

	authCookie, err := req.Cookie("auth")
	if err != nil {
		r.Tracer.Trace(err)
		return
	}

	client := &Client{
		Socket: socket,
		Send:   make(chan *Message, messageBufferSize),
		Room:   r,
		userData: objx.MustFromBase64(authCookie.Value),

	}
	r.Join <- client

	defer func() { r.Leave <- client }()

	go client.write()
	client.read()
}
