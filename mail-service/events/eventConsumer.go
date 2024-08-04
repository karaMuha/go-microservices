package events

import (
	"encoding/json"
	"fmt"
	"log"
	"mailer/models"
	"mailer/services"

	"github.com/rabbitmq/amqp091-go"
)

type EventConsumer struct {
	connection  *amqp091.Connection
	mailService services.MailServiceInterface
}

func NewEventConsumer(connection *amqp091.Connection, mailService services.MailServiceInterface) (EventConsumer, error) {
	consumer := EventConsumer{
		connection:  connection,
		mailService: mailService,
	}

	channel, err := consumer.connection.Channel()

	if err != nil {
		return EventConsumer{}, err
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
		return EventConsumer{}, err
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
		var eventPayload models.SignupEvent
		_ = json.Unmarshal(data.Body, &eventPayload)

		go consumer.handleSignupEventPayload(eventPayload)
	}

	return nil
}

func (consumer *EventConsumer) handleSignupEventPayload(payload models.SignupEvent) {
	log.Printf("Signup event for user %s received with verification token %s", payload.Email, payload.VerificationToken)
	mailMessage := fmt.Sprintf("Please visit localhost:8081/confirm/%s/%s to complete your registration", payload.Email, payload.VerificationToken)
	mail := &models.Mail{
		To:      payload.Email,
		Subject: "Confirm registration",
		Message: mailMessage,
	}

	_ = consumer.mailService.SendMail(mail)
}
