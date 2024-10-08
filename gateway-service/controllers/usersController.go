package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/models"
	"io"
	"net/http"
	"strings"
	"time"
)

type UsersController struct {
	address string
}

func NewUsersController(address string) *UsersController {
	return &UsersController{
		address: address,
	}
}

func (ac UsersController) HandleSignup(w http.ResponseWriter, r *http.Request) {
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
	body, err := io.ReadAll(response.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(body)
}

func (ac UsersController) HandleConfirmEmail(w http.ResponseWriter, r *http.Request) {
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
	body, err := io.ReadAll(response.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(body)
}

func (ac UsersController) HandleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")

	url := fmt.Sprintf("%s/users/%s", ac.address, email)

	request, err := http.NewRequest("GET", url, nil)

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
	body, err := io.ReadAll(response.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if response.StatusCode != http.StatusOK {
		http.Error(w, string(body), response.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ac UsersController) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/users", ac.address)

	request, err := http.NewRequest("GET", url, nil)

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
	body, err := io.ReadAll(response.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if response.StatusCode != http.StatusOK {
		http.Error(w, string(body), response.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ac UsersController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/login", ac.address)
	request, err := http.NewRequest("POST", url, r.Body)

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
	body, err := io.ReadAll(response.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if response.StatusCode != http.StatusOK {
		http.Error(w, string(body), response.StatusCode)
		return
	}

	cookies := response.Cookies()
	var jwtCookie *http.Cookie

	for _, v := range cookies {
		if strings.EqualFold(v.Name, "jwt") {
			jwtCookie = v
		}
	}

	if jwtCookie == nil {
		http.Error(w, "No jwt cookie present", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, jwtCookie)
	w.WriteHeader(response.StatusCode)
}

func (ac UsersController) HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now(),
	})

	w.WriteHeader(http.StatusOK)
}
