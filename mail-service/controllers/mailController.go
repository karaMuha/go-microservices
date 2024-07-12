package controllers

import (
	"encoding/json"
	"mailer/models"
	"mailer/services"
	"net/http"
)

type MailController struct {
	mailService services.MailServiceInterface
}

func NewMailController(mailService services.MailServiceInterface) *MailController {
	return &MailController{
		mailService: mailService,
	}
}

func (mc MailController) HandleSendMail(w http.ResponseWriter, r *http.Request) {
	var mail models.Mail
	err := json.NewDecoder(r.Body).Decode(&mail)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	responseErr := mc.mailService.SendMail(&mail)

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
	}

	w.WriteHeader(http.StatusOK)
}
