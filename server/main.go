package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},	
		AllowedHeaders: []string{
			"*",	
		},	
	})

	server := &http.Server{
		Addr:         *addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      c.Handler(router),
		ErrorLog:     logger,
	}

	logger.Printf("The server is listening on port %v", *addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %q: %s\n", *addr, err)
	}
}

// SocketHandler handles web socket calls
func SocketHandler(w http.ResponseWriter, r *http.Request) {
	serveWebSocket(hub, w, r)
}

// LoginHandler handles authentication logic. TODO: ıt is not completed though. Proper logic must be implemented
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	logger.Println(r.Header)

	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	if err != nil {
		log.Fatalln("Error LoginHandler", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error LoginHandler", err)
	}

	if err := json.Unmarshal(body, &user); err != nil { // unmarshall body contents as a type Candidate
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error LoginHandler unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError) // unprocessable entity
			return
		}
	}

	w.Write([]byte(user.Username))
	return
}
