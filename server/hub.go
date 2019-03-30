package main

// Hub is used for maintaining the active clients and broadcasts messages to all clients
type Hub struct {
	clients    map[*Client]bool
	rooms      map[*Room]bool
	inbound    chan []byte
	register   chan *Client
	unregister chan *Client
}

func initHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		rooms:      make(map[*Room]bool),
		inbound:    make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// endlessly listens register / unregister / inbound channels
func (h *Hub) start() {
	for {
		select {
		case client := <-h.register:
			var room *Room
			if len(h.rooms) == 0 {
				room = initRoom(h, "general") // hardcoded for now
				h.rooms[room] = true
				go room.start()
			}
			client.room = room
			for r := range h.rooms {
				if r.name == "general" { // hardcoded for now
					r.clients[client] = true
					r.register <- client
				}
			}
		case client := <-h.unregister:
			for r := range h.rooms {
				if r.name == "general" { // hardcoded for now
					r.unregister <- client
				}
			}
		case message := <-h.inbound:
			for r := range h.rooms {
				if r.name == "general" { // hardcoded for now
					r.inbound <- message
				}
			}
		}
	}
}

func (h *Hub) send(message []byte) {
	for client := range h.clients {
		client.send <- message
	}
}
