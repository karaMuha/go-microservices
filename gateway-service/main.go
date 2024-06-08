package main

import (
	"gateway/server"
	"log"
	"os"
)

func main() {
	serverPort := os.Getenv("SERVER_PORT")
	log.Printf("Starting http server on port %s", serverPort)
	httpServer := server.InitHttpServer(serverPort)

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting http server: %v", err)
	}
}
