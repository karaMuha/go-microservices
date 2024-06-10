package models

type LoginCredentials struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
