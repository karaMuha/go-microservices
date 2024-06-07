package repositories

import "authentication/models"

type UsersRepositoryInterface interface {
	QueryCreateUser(user *models.User, hashedPassword string) *models.ResponseError

	QueryGetAllUsers() ([]*models.User, *models.ResponseError)

	QueryGetUserByEmail(email string) (*models.User, *models.ResponseError)

	QueryUpdateUser(user *models.User) *models.ResponseError

	QueryDeleteUser(id string) *models.ResponseError
}
