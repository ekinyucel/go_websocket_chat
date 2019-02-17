package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// it is called when upgrading the HTTP connection to a websocket connection.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true // disabling CORS
	},
}

// WebSocket struct stores the websocket connection details
type WebSocket struct {
	Conn   *websocket.Conn
	Out    chan []byte
	In     chan []byte
	Events map[string]EventHandler
}

// OpenWebSocketConnection is used for upgrading the connection and opening a web socket connection between the client and server
// Instantinating Websocket struct
func OpenWebSocketConnection(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("An error occurred while upgrading the connection")
	}

	ws := &WebSocket{
		Conn:   conn,
		Out:    make(chan []byte),
		In:     make(chan []byte),
		Events: make(map[string]EventHandler),
	}

	go ws.Reader() // reading real-time message from client
	go ws.Writer() // starting writer method of websocket as a goroutine

	return ws, nil
}

// Reader method runs and collects the messages from client
func (ws *WebSocket) Reader() {
	defer func() {
		ws.Conn.Close()
	}()
	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			logger.Printf("WebSocket error message: %v", err)
			break
		}
		event, err := GenerateEvent(message)
		if err != nil {
			logger.Printf("Unable to parse the message: %v", err)
		} else {
			logger.Printf("Message: %v", event)
		}
		if action, ok := ws.Events[event.Name]; ok {
			action(event)
		}
	}
}

// Writer method runs and collects the messages from client
func (ws *WebSocket) Writer() {
	defer func() {
		ws.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-ws.Out:
			if !ok {
				ws.Conn.WriteMessage(websocket.CloseMessage, make([]byte, 0))
			}
			w, err := ws.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			w.Close()
		}
	}
}

// On method is used for assigning actions to events
func (ws *WebSocket) On(eventName string, action EventHandler) *WebSocket {
	ws.Events[eventName] = action
	return ws
}
