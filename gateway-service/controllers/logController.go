package controllers

import (
	"fmt"
	"io"
	"net/http"
)

type LogController struct {
	address string
}

func NewLogController(address string) *LogController {
	return &LogController{
		address: address,
	}
}

func (lc LogController) HandleGetLogById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	url := fmt.Sprintf("%s/log/%s", lc.address, id)

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

func (lc LogController) HandleGetAllLogs(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/log", lc.address)

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
