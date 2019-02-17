package main

import (
	"encoding/json"
)

// EventHandler type is a function which accepts event as an argument
type EventHandler func(*Event)

// Event struct holds the event information. Name of an event, and the data type that event contains.
type Event struct {
	Name string      `json:"event"`
	Data interface{} `json:"data"`
	Date int         `json:"date"`
}

// GenerateEvent function is used to create an Event type from input
// Unmarshales input into Event type
func GenerateEvent(input []byte) (*Event, error) {
	event := new(Event)
	err := json.Unmarshal(input, event)
	return event, err
}

// Marshal method is used for marshaling event type
func (e *Event) Marshal() []byte {
	output, _ := json.Marshal(e)
	return output
}
