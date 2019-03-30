package main

// Room struct repsents a chat room inside a hub
type Room struct {
	hub        *Hub
	name       string
	clients    map[*Client]bool
	inbound    chan []byte
	register   chan *Client
	unregister chan *Client
}

func initRoom(hub *Hub, name string) *Room {
	return &Room{
		hub:        hub,
		name:       name,
		clients:    make(map[*Client]bool),
		inbound:    make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// endlessly listens register / unregister / inbound channels
func (r *Room) start() {
	for {
		select {
		case client := <-r.register:
			r.clients[client] = true
			m := &Message{Username: client.name, Data: len(r.clients), Room: r.name, Type: "connect"}
			r.send(m.marshal())
		case client := <-r.unregister:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				m := &Message{Data: len(r.clients), Room: r.name, Type: "disconnect"}
				r.send(m.marshal())
				close(client.send)
			}
		case message := <-r.inbound:
			for client := range r.clients {
				select {
				case client.send <- message: // sending incoming message to the send channel of each client
				default:
					close(client.send)
					delete(r.clients, client)
				}
			}
		}
	}
}

func (r *Room) send(message []byte) {
	for client := range r.clients {
		client.send <- message
	}
}
