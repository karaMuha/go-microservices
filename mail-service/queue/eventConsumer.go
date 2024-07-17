package queue

import (
	"encoding/json"
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
		"auth_topic",
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
			"auth_topic",
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

	forever := make(chan bool)
	go func() {
		for data := range messages {
			var eventPayload models.EventPayload
			_ = json.Unmarshal(data.Body, &eventPayload)

			go handlePayload(eventPayload)
		}
	}()
	<-forever

	return nil
}

func handlePayload(payload models.EventPayload) {
	switch payload.Name {
	case "signup":
		log.Println("Signup event received")
	}
}
