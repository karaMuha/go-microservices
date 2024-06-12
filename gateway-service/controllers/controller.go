package controllers

import (
	"bytes"
	"encoding/json"
	"gateway/models"
	"log"
	"net/http"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl Controller) HandlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("Server pinged!")
	w.WriteHeader(http.StatusOK)
}

func (ctrl Controller) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var requestPayload models.RequestPayload

	err := json.NewDecoder(r.Body).Decode(&requestPayload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch requestPayload.Service {
	case "auth":
		handleAuthRequest(w, requestPayload)
	default:
		http.Error(w, "No target", http.StatusBadRequest)
	}
}

func handleAuthRequest(w http.ResponseWriter, payload models.RequestPayload) {
	jsonData, err := json.Marshal(payload.AuthPayload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	request, err := http.NewRequest(payload.Action, payload.Endpoint, bytes.NewBuffer(jsonData))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Error(w, "placeholder", response.StatusCode)
		return
	}
}
