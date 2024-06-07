package server

import (
	"database/sql"
	"net/http"

	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func InitHttpServer(config *viper.Viper, db *sql.DB) *http.Server {
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
		Addr:    config.GetString("SERVER_PORT"),
		Handler: handler,
	}
}
