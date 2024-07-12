package models

type Mail struct {
	From     string `json:"from"`
	FromName string `json:"fromName"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Message  string `json:"message"`
	DataMap  map[string]any
}
