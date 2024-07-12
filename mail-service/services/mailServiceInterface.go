package services

import "mailer/models"

type MailServiceInterface interface {
	SendMail(mail *models.Mail) *models.ResponseError
}
