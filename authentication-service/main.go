package main

import (
	"authentication/server"
	"log"
)

func main() {
	log.Println("Starting authentication service")

	log.Println("Initializing database")
	db := server.ConnectToDb()

	log.Println("Initializing http server")
	httpServer := server.InitHttpServer(db)

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting http server: %v", err)
	}
}
