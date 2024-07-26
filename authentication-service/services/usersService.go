package services

import (
	"authentication/events"
	"authentication/models"
	"authentication/repositories"
	"authentication/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	usersRepository repositories.UsersRepositoryInterface
	eventProducer   *events.EventProducer
}

func NewUsersService(usersRepository repositories.UsersRepositoryInterface, eventProducer *events.EventProducer) UsersServiceInterface {
	return &UsersService{
		usersRepository: usersRepository,
		eventProducer:   eventProducer,
	}
}

func (us UsersService) SignupUser(user *models.User) *models.ResponseError {
	userInDb, responseErr := us.GetUserByEmail(user.Email)

	if responseErr != nil {
		return responseErr
	}

	if userInDb != nil {
		return &models.ResponseError{
			Message: "Email already exists",
			Status:  http.StatusConflict,
		}
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	verificationToken := randstr.String(64)

	responseErr = us.usersRepository.QueryCreateUser(user, hashedPassword, verificationToken)

	if responseErr != nil {
		return responseErr
	}

	signupEvent := models.SignupEvent{
		Email:             user.Email,
		VerificationToken: verificationToken,
	}

	jsonSignupEvent, err := json.Marshal(&signupEvent)

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	err = us.eventProducer.PushEvent(jsonSignupEvent, "log.INFO")

	if err != nil {
		log.Println(err)
	}

	return nil
}

func (us UsersService) ConfirmUserRegistration(email string, verificationToken string) *models.ResponseError {
	user, respErr := us.GetUserByEmail(email)

	if respErr != nil {
		return respErr
	}

	if user == nil {
		return &models.ResponseError{
			Message: "User not found",
			Status:  http.StatusNotFound,
		}
	}

	if verificationToken != user.VerificationToken {
		return &models.ResponseError{
			Message: "Token does not match",
			Status:  http.StatusUnauthorized,
		}
	}

	user.Active = true

	return us.UpdateUser(user)
}

func (us UsersService) GetUserByEmail(email string) (*models.User, *models.ResponseError) {
	return us.usersRepository.QueryGetUserByEmail(email)
}

func (us UsersService) GetAllUsers() ([]*models.User, *models.ResponseError) {
	return us.usersRepository.QueryGetAllUsers()
}

func (us UsersService) UpdateUser(user *models.User) *models.ResponseError {
	responseErr := us.checkPermission(user)

	if responseErr != nil {
		return responseErr
	}

	return us.usersRepository.QueryUpdateUser(user)
}

func (us UsersService) DeleteUser(user *models.User) *models.ResponseError {
	responseErr := us.checkPermission(user)

	if responseErr != nil {
		return responseErr
	}

	return us.usersRepository.QueryDeleteUser(user.ID)
}

func (us UsersService) LoginUser(email string, password string) (string, *models.ResponseError) {
	userInDb, responseErr := us.usersRepository.QueryGetUserByEmail(email)

	if responseErr != nil {
		return "", responseErr
	}

	if userInDb == nil {
		return "", &models.ResponseError{
			Message: "User not found",
			Status:  http.StatusNotFound,
		}
	}

	if !userInDb.Active {
		return "", &models.ResponseError{
			Message: "User is inactive",
			Status:  http.StatusConflict,
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(userInDb.Password), []byte(password))

	if err != nil {
		return "", &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusUnauthorized,
		}
	}

	token, err := utils.GenerateJwt(userInDb.ID)

	if err != nil {
		return "", &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return token, nil
}

func (us UsersService) checkPermission(user *models.User) *models.ResponseError {
	userInDb, responseErr := us.usersRepository.QueryGetUserByEmail(user.Email)

	if responseErr != nil {
		return responseErr
	}

	if userInDb == nil {
		return &models.ResponseError{
			Message: "User not found",
			Status:  http.StatusNotFound,
		}
	}

	if user.ID != userInDb.ID {
		return &models.ResponseError{
			Message: "Access denied",
			Status:  http.StatusUnauthorized,
		}
	}

	return nil
}
