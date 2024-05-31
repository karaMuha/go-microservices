package main

import (
	"gateway/server"
	"log"
)

func main() {
	httpServer := server.InitHttpServer()
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting http server: %v", err)
	}
}
