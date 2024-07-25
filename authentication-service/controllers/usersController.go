package controllers

import (
	"authentication/models"
	"authentication/services"
	"authentication/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
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

func (uc UsersController) HandleConfirmRegistration(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")
	verificationToken := r.PathValue("token")

	responseErr := uc.usersService.ConfirmUserRegistration(email, verificationToken)

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

func (uc UsersController) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	bodyDecoder := json.NewDecoder(r.Body)

	responseErr := parseUser(&user, bodyDecoder)

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	responseErr = uc.usersService.UpdateUser(&user)

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc UsersController) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	bodyDecoder := json.NewDecoder(r.Body)

	responseErr := parseUser(&user, bodyDecoder)

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	responseErr = uc.usersService.DeleteUser(&user)

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc UsersController) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	var loginCredentials models.LoginCredentials
	bodyDecorder := json.NewDecoder(r.Body)

	responseErr := parseUser(&loginCredentials, bodyDecorder)

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	jwtToken, responseErr := uc.usersService.LoginUser(loginCredentials.Email, loginCredentials.Password)

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    jwtToken,
		Secure:   true,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour),
	})

	w.WriteHeader(http.StatusOK)
}

func (uc UsersController) HandleLogoutUser(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now(),
	})

	w.WriteHeader(http.StatusOK)
}

func parseUser(data any, bodyDecoder *json.Decoder) *models.ResponseError {
	err := bodyDecoder.Decode(data)

	if err != nil {
		log.Println(err)
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	err = utils.Validator.Struct(data)

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
