package main

import (
	"authentication/config"
	"authentication/server"
	"log"
	"os"
)

func main() {
	log.Println("Starting eventom app")

	log.Println("Reading environment variables")
	appEnvironment := os.Getenv("APP_ENV")

	if appEnvironment == "" {
		log.Fatal("Could not get app environment")
	}

	config := config.ReadEnvFile(appEnvironment)

	log.Println("Initializing database")
	db := server.InitDatabase(config)

	log.Println("Initializinh http server")
	httpServer := server.InitHttpServer(config, db)

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting http server: %v", err)
	}
}
