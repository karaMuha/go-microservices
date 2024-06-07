package controllers

import (
	"authentication/services"
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

func (uc UsersController) HandleSignupUser(w http.ResponseWriter, r *http.Request) {}

func (uc UsersController) HandleGetUsers(w http.ResponseWriter, r *http.Request) {}

func (uc UsersController) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {}

func (uc UsersController) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {}

func (uc UsersController) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {}

func (uc UsersController) HandleLoginUser(w http.ResponseWriter, r *http.Request) {}
