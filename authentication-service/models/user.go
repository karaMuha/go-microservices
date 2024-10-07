package models

import "time"

type User struct {
	ID                string    `json:"id" validate:"omitempty,uuid"`
	Email             string    `json:"email" validate:"required,email"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Password          string    `json:"password" validate:"required"`
	Active            bool      `json:"active"`
	VerificationToken string    `json:"verficationToken"`
	Role              string    `json:"role"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
