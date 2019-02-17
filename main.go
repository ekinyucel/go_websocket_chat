package main

import (
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "main package ", log.LstdFlags|log.Lshortfile)

func main() {
	http.Handle(`/`, http.FileServer(http.Dir("./assets")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := OpenWebSocketConnection(w, r)
		if err != nil {
			logger.Printf("Error creating web socket connection: %v", err)
			return
		}
		ws.On("message", func(e *Event) {
			logger.Printf("The message received: %s", e.Data.(string))
			ws.Out <- (&Event{
				Name: "response",
				Data: e.Data.(string),
				Date: e.Date,
			}).Marshal()
		})
	})

	logger.Println("The server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
