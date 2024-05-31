package controllers

import (
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
