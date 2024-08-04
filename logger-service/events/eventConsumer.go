package events

import (
	"logger/models"
	"logger/services"

	"github.com/rabbitmq/amqp091-go"
)

type EventConsumer struct {
	connection      *amqp091.Connection
	logEntryService services.LogEntryServiceInterface
}

func NewEventConsumer(connection *amqp091.Connection, logEntryService services.LogEntryServiceInterface) (EventConsumer, error) {
	consumer := EventConsumer{
		connection:      connection,
		logEntryService: logEntryService,
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

	go func() {
		for data := range messages {
			go consumer.handleSignupEventPayload(data.Body)
		}
	}()

	return nil
}

func (consumer *EventConsumer) handleSignupEventPayload(payload []byte) {
	logEntry := &models.LogEntry{
		Category: "Registration",
		Data:     string(payload),
	}

	consumer.logEntryService.InsertLogEntry(logEntry)
}
