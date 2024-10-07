package server

import (
	"gateway/controllers"
	"gateway/middlewares"
	"gateway/utils"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func InitHttpServer(serverPort string) *http.Server {
	err := utils.ReadPrivateKeyFromFile(os.Getenv("PRIVATE_KEY_PATH"))
	if err != nil {
		log.Fatalf("Error while reading private key: %v", err)
	}

	utils.SetProtectedRoutes()

	controller := controllers.NewController()
	authController := controllers.NewAuthController("http://authentication-service:8080")
	logController := controllers.NewLogController("http://logger-service:8080")
	router := http.NewServeMux()

	router.HandleFunc("GET /ping", controller.HandlePing)

	router.HandleFunc("POST /users/signup", authController.HandleSignup)
	router.HandleFunc("POST /users/confirm/{email}/{token}", authController.HandleConfirmEmail)
	router.HandleFunc("GET /users/{email}", authController.HandleGetUserByEmail)
	router.HandleFunc("GET /users", authController.HandleGetAllUsers)
	router.HandleFunc("POST /users/login", authController.HandleLogin)
	router.HandleFunc("POST /users/logout", authController.HandleLogout)

	router.HandleFunc("GET /logs/{id}", logController.HandleGetLogById)
	router.HandleFunc("GET /logs", logController.HandleGetAllLogs)

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
		Handler: middlewares.AuthMiddleware(handler),
	}
}
