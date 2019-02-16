package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle(`/`, http.FileServer(http.Dir("./assets")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := OpenWebSocketConnection(w, r)
		if err != nil {
			log.Printf("Error creating web socket connection: %v", err)
			return
		}
		ws.On("message", func(e *Event) {
			log.Printf("The message received: %s", e.Data.(string))
			ws.Out <- (&Event{
				Name: "response",
				Data: e.Data.(string),
			}).Marshal()
		})
	})

	log.Println("The server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
