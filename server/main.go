package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":8080", "http server address")
var logger = log.New(os.Stdout, "main package ", log.LstdFlags|log.Lshortfile)

func main() {
	flag.Parse()

	hub := initHub()
	go hub.start()

	http.Handle(`/`, http.FileServer(http.Dir("./assets")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebSocket(hub, w, r)
		enableCors(&w)
	})

	logger.Printf("The server is listening on port %v", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		logger.Fatalf("ListenAndServe: %v", err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
