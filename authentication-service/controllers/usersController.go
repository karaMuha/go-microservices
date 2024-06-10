package controllers

import (
	"authentication/models"
	"authentication/services"
	"authentication/utils"
	"encoding/json"
	"net/http"
)

type UsersController struct {
	usersService services.UsersServiceInterface
}

func NewUsersController(usersService services.UsersServiceInterface) *UsersController {
	return &UsersController{
		usersService: usersService,
	}
}

func (uc UsersController) HandleSignupUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	bodyDecoder := json.NewDecoder(r.Body)

	responseErr := parseUser(&user, bodyDecoder)

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	responseErr = uc.usersService.SignupUser(&user)

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc UsersController) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")

	user, responseErr := uc.usersService.GetUserByEmail(email)

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	responseJson, err := json.Marshal(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (uc UsersController) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, responseErr := uc.usersService.GetAllUsers()

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	responseJson, err := json.Marshal(&users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (uc UsersController) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {}

func (uc UsersController) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {}

func (uc UsersController) HandleLoginUser(w http.ResponseWriter, r *http.Request) {}

func (uc UsersController) HandleLogoutUser(w http.ResponseWriter, r *http.Request) {}

func parseUser(user *models.User, bodyDecoder *json.Decoder) *models.ResponseError {
	err := bodyDecoder.Decode(user)

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	err = utils.Validator.Struct(user)

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
