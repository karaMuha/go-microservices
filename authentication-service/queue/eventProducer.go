package queue

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type EventProducer struct {
	connection *amqp091.Connection
}

func NewEventProducer(connection *amqp091.Connection) (*EventProducer, error) {
	eventProducer := EventProducer{
		connection: connection,
	}

	err := eventProducer.setup()

	if err != nil {
		return nil, err
	}

	return &eventProducer, nil
}

func (ep *EventProducer) setup() error {
	channel, err := ep.connection.Channel()

	if err != nil {
		return err
	}

	return channel.ExchangeDeclare(
		"auth_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}

func (ep *EventProducer) PushEvent(event []byte, severity string) error {
	channel, err := ep.connection.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()

	log.Println("Publishing to Queue")

	err = channel.Publish(
		"auth_topic",
		severity,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        event,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
