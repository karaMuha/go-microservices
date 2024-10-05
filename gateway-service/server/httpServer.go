package server

import (
	"gateway/controllers"
	"net/http"

	"github.com/rs/cors"
)

func InitHttpServer(serverPort string) *http.Server {
	controller := controllers.NewController()
	authController := controllers.NewAuthController("http://authentication-service:8080")
	logController := controllers.NewLogController("http://logger-service:8080")
	router := http.NewServeMux()

	router.HandleFunc("GET /ping", controller.HandlePing)

	router.HandleFunc("POST /signup", authController.HandleSignup)
	router.HandleFunc("POST /confirm/{email}/{token}", authController.HandleConfirmEmail)
	router.HandleFunc("GET /users/{email}", authController.HandleGetUserByEmail)
	router.HandleFunc("GET /users", authController.HandleGetAllUsers)

	router.HandleFunc("GET /log/{id}", logController.HandleGetLogById)
	router.HandleFunc("GET /log", logController.HandleGetAllLogs)

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
