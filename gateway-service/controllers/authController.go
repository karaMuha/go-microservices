package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/models"
	"io"
	"net/http"
)

type AuthController struct {
	address string
}

func NewAuthController(address string) *AuthController {
	return &AuthController{
		address: address,
	}
}

func (ac AuthController) HandleSignup(w http.ResponseWriter, r *http.Request) {
	var signupUser models.SignupUser
	err := json.NewDecoder(r.Body).Decode(&signupUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(signupUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("%s/signup", ac.address)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	w.WriteHeader(response.StatusCode)
}

func (ac AuthController) HandleConfirmEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")
	verificationToken := r.PathValue("token")

	url := fmt.Sprintf("%s/confirm/%s/%s", ac.address, email, verificationToken)

	request, err := http.NewRequest("POST", url, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	w.WriteHeader(response.StatusCode)
}

func (ac AuthController) HandleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")

	url := fmt.Sprintf("%s/users/%s", ac.address, email)

	request, err := http.NewRequest("POST", url, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Error(w, "", response.StatusCode)
		return
	}

	responseBody, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func (ac AuthController) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/users", ac.address)

	request, err := http.NewRequest("POST", url, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Error(w, "", response.StatusCode)
		return
	}

	responseBody, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func (ac AuthController) HandleLogin(w http.ResponseWriter, r *http.Request) {}

func (ac AuthController) HandleLogout(w http.ResponseWriter, r *http.Request) {}
