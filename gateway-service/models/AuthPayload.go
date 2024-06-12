package models

import "time"

type AuthPayload struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
