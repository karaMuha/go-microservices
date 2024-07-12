package main

import (
	"log"
	"mailer/httpserver"
	"mailer/mailserver"
)

func main() {
	log.Println("Starting mail service")

	log.Println("Initialize mail server")
	mailServer := mailserver.NewMailServer()

	log.Println("Starting http server")
	httpServer := httpserver.InitHttpServer(mailServer)
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting http server: %v", err)
	}
}
