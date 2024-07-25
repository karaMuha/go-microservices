package models

type SignupEvent struct {
	Email             string `json:"email"`
	VerificationToken string `json:"verificationToken"`
}
