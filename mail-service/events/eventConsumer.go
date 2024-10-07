package events

import (
	"encoding/json"
	"fmt"
	"mailer/models"
	"mailer/services"

	"github.com/rabbitmq/amqp091-go"
)

type EventConsumer struct {
	connection  *amqp091.Connection
	mailService services.MailServiceInterface
}

func NewEventConsumer(connection *amqp091.Connection, mailService services.MailServiceInterface) (*EventConsumer, error) {
	consumer := &EventConsumer{
		connection:  connection,
		mailService: mailService,
	}

	channel, err := consumer.connection.Channel()

	if err != nil {
		return nil, err
	}

	err = channel.ExchangeDeclare(
		"signup_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (consumer *EventConsumer) Listen(topics []string) error {
	channel, err := consumer.connection.Channel()

	if err != nil {
		return err
	}

	queue, err := channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for _, str := range topics {
		err = channel.QueueBind(
			queue.Name,
			str,
			"signup_topic",
			false,
			nil,
		)
	}

	if err != nil {
		return err
	}

	messages, err := channel.Consume(queue.Name, "", true, false, false, false, nil)

	if err != nil {
		return err
	}

	for data := range messages {
		go consumer.handleSignupEventPayload(data)
	}

	return nil
}

func (consumer *EventConsumer) handleSignupEventPayload(data amqp091.Delivery) {
	var eventPayload models.SignupEvent
	_ = json.Unmarshal(data.Body, &eventPayload)
	mailMessage := fmt.Sprintf("Please visit localhost:8080/users/confirm/%s/%s to complete your registration", eventPayload.Email, eventPayload.VerificationToken)
	mail := &models.Mail{
		To:      eventPayload.Email,
		Subject: "Confirm registration",
		Message: mailMessage,
	}

	_ = consumer.mailService.SendMail(mail)
}
