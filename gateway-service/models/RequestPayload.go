package models

type RequestPayload struct {
	Service     string
	Endpoint    string
	Action      string
	AuthPayload AuthPayload
}
