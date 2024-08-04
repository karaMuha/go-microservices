package main

import (
	"log"
	"mailer/events"
	"mailer/mailserver"
	"mailer/services"
)

func main() {
	log.Println("Starting mail service")

	log.Println("Connecting to message queue")
	mqConnection, err := events.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer mqConnection.Close()

	log.Println("Initializing mail server")
	mailServer := mailserver.NewMailServer()

	log.Println("Setting up event consumer")
	mailService := services.NewMailServiceImpl(mailServer)

	eventConsumer, err := events.NewEventConsumer(mqConnection, mailService)

	if err != nil {
		log.Fatalf("Error while initializing event consumer: %v", err)
	}

	log.Println("Start listening")
	err = eventConsumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})

	if err != nil {
		log.Fatalf("Cannot listen on queue: %v", err)
	}
}
