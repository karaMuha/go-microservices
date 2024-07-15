package main

import (
	"authentication/queue"
	"authentication/server"
	"log"
	"os"
)

func main() {
	log.Println("Starting authentication service")

	log.Println("Initializing database")
	db := server.ConnectToDb()
	log.Println("Connected to database")

	log.Println("Connecting to message queue")
	mqConnection, err := queue.Connect()

	if err != nil {
		log.Fatalf(err.Error())
	}

	defer mqConnection.Close()

	log.Println("Initializing http server")
	httpServer := server.InitHttpServer(db, mqConnection)

	log.Printf("Starting http server on port %s", os.Getenv("SERVER_PORT"))
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting http server: %v", err)
	}
}
