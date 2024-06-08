package server

import (
	"authentication/utils"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func InitHttpServer(db *sql.DB) *http.Server {
	err := utils.ReadPrivateKeyFromFile(os.Getenv("PRIVATE_KEY_PATH"))
	if err != nil {
		log.Fatalf("Error while reading private key: %v", err)
	}

	router := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	handler := c.Handler(router)

	return &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
}
