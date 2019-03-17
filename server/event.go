package main

import "encoding/json"

// Event struct is used for storing the event as a Json object.
type Event struct {
	Username string      `json:"username"`
	Event    string      `json:"event"`
	Data     interface{} `json:"data"`
	Date     int         `json:"date"`
}

// UnMarshalEvent function is used to create an Event type from input
// Unmarshales input into Event type
func UnMarshalEvent(input []byte) (*Event, error) {
	event := new(Event)
	err := json.Unmarshal(input, event)
	return event, err
}

// Marshal method is used for marshaling event type
func (e *Event) marshal() []byte {
	output, _ := json.Marshal(e)
	return output
}
