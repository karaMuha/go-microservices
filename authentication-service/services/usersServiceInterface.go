package services

import "authentication/models"

type UsersServiceInterface interface {
	SignupUser(user *models.User) *models.ResponseError

	ConfirmUserRegistration(email string, verificationToken string) *models.ResponseError

	GetUserByEmail(email string) (*models.User, *models.ResponseError)

	GetAllUsers() ([]*models.User, *models.ResponseError)

	UpdateUser(user *models.User) *models.ResponseError

	DeleteUser(user *models.User) *models.ResponseError

	LoginUser(email string, password string) (string, *models.ResponseError)
}
