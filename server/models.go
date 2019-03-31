package main

// Response model represents the http response json object
type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}
