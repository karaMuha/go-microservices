package server

import (
	"gateway/controllers"
	"net/http"

	"github.com/rs/cors"
)

func InitHttpServer(serverPort string) *http.Server {
	controller := controllers.NewController()
	router := http.NewServeMux()

	router.HandleFunc("GET /ping", controller.HandlePing)

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
		Addr:    serverPort,
		Handler: handler,
	}
}
