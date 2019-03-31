package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// SocketHandler handles web socket calls
func SocketHandler(w http.ResponseWriter, r *http.Request) {
	serveWebSocket(hub, w, r)
}

// LoginHandler handles authentication logic. TODO: ıt is not completed though. Proper logic must be implemented
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	if err != nil {
		logger.Fatalln("Error LoginHandler", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		logger.Fatalln("Error LoginHandler", err)
	}

	if err := json.Unmarshal(body, &user); err != nil { // unmarshall body contents as a type Candidate
		logger.Println("username ", user.Username, " password ", user.Password)
		logger.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			logger.Fatalln("Error LoginHandler unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError) // unprocessable entity
			return
		}
	}

	response, err := json.Marshal(&Response{Message: user.Username, StatusCode: http.StatusOK})

	if err != nil {
		logger.Fatalln("response json marshall error")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}
