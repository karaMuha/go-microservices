package server

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

type HttpServer struct {
	server *http.Server
}

func InitHttpServer() HttpServer {
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

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	return HttpServer{
		server: server,
	}
}

func (hs HttpServer) Start() {
	err := hs.server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}
