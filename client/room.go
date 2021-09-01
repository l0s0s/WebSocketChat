package client

type Room struct {
	Forward chan []byte
	Join    chan *Client
	Leave   chan *Client
	Clients map[*Client]bool
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.Join:
			r.Clients[client] = true
		case client := <-r.Leave:
			delete(r.Clients, client)
			close(client.Send)
		case msg := <-r.Forward:
			for client := range r.Clients {
				client.Send <- msg
			}
		}
	}
}
