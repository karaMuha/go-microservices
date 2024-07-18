package models

type EventPayload struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}
