package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":8080", "http server address")
var logger = log.New(os.Stdout, "main package ", log.LstdFlags|log.Lshortfile)
var hub *Hub

func main() {
	flag.Parse()

	hub = initHub()
	go hub.start()

	router := mux.NewRouter()
	router.HandleFunc("/ws", SocketHandler)
	router.HandleFunc("/login", LoginHandler).Methods("POST")

	server := NewServer(router, *addr)

	logger.Printf("The server is listening on port %v", *addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %q: %s\n", *addr, err)
	}
}
