package events

import (
	"log"
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

	log.Println("Got channel")

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

	log.Println("Queues declared")

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

	log.Println("Queues bond")

	messages, err := channel.Consume(queue.Name, "", true, false, false, false, nil)

	if err != nil {
		return err
	}

	log.Println("Consuming")

	forever := make(chan bool)
	go func() {
		for data := range messages {
			go consumer.handleSignupEventPayload(data.Body)
		}
	}()
	<-forever

	log.Println("Forever thing")
	return nil
}

func (consumer *EventConsumer) handleSignupEventPayload(payload []byte) {
	logEntry := &models.LogEntry{
		Name: "Registration",
		Data: string(payload),
	}

	consumer.logEntryService.InsertLogEntry(logEntry)
}
