package main

import (
	"context"
	"log"
	"logger/events"
	"logger/server"
)

func main() {
	log.Println("Starting logger service")

	log.Println("Connecting to database")
	mongoClient, err := server.ConnectToDb()
	if err != nil {
		log.Fatalf("Could not connect to database")
	}

	// disconnect from mongo when server shuts down
	defer func() {
		err = mongoClient.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()

	log.Println("Connecting to message queue")
	mqConnection, err := events.Connect()

	if err != nil {
		log.Fatalf(err.Error())
	}

	defer mqConnection.Close()

	log.Println("Starting http server")
	httpServer, err := server.InitHttpServer(mongoClient, mqConnection)

	if err != nil {
		log.Fatal(err)
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting http server: %v", err)
	}
}
