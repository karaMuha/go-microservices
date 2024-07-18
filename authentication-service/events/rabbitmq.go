package events

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect() (*amqp.Connection, error) {
	count := 0

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")

		if err == nil {
			return c, nil
		}

		log.Println("RabbitMQ not yet ready...")
		count++

		if count > 5 {
			return nil, err
		}

		log.Println("backing off...")
		time.Sleep(10 * time.Second)
		continue
	}
}
