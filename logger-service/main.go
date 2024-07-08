package main

import (
	"context"
	"log"
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

	httpServer := server.InitHttpServer(mongoClient)
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting http server: %v", err)
	}
}
