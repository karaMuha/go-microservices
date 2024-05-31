package main

import (
	"gateway/server"
	"log"
)

func main() {
	log.Println("Starting http server on port 8080")
	httpServer := server.InitHttpServer()
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting http server: %v", err)
	}
}
