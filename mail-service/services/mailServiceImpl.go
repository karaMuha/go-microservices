package services

import (
	"mailer/mailserver"
	"mailer/models"
	"net/http"
)

type MailServiceImpl struct {
	mailServer *mailserver.MailServer
}

func NewMailServiceImpl(mailServer *mailserver.MailServer) MailServiceInterface {
	return &MailServiceImpl{
		mailServer: mailServer,
	}
}

func (ms *MailServiceImpl) SendMail(mail *models.Mail) *models.ResponseError {
	err := ms.mailServer.SendSMTPMail(mail)

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
