package main

import (
	"log"
	"mailer/events"
	"mailer/httpserver"
	"mailer/mailserver"
)

func main() {
	log.Println("Starting mail service")

	mqConnection, err := events.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer mqConnection.Close()

	log.Println("Initialize mail server")
	mailServer := mailserver.NewMailServer()

	log.Println("Starting http server")
	httpServer, err := httpserver.InitHttpServer(mailServer, mqConnection)

	if err != nil {
		log.Fatal(err)
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting http server: %v", err)
	}
}
