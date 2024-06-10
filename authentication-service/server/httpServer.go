package server

import (
	"authentication/controllers"
	"authentication/repositories"
	"authentication/services"
	"authentication/utils"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
)

func InitHttpServer(db *sql.DB) *http.Server {
	// setup utils
	err := utils.ReadPrivateKeyFromFile(os.Getenv("PRIVATE_KEY_PATH"))
	if err != nil {
		log.Fatalf("Error while reading private key: %v", err)
	}
	utils.Validator = validator.New()

	// initialize layers
	usersRepository := repositories.NewUsersRepository(db)
	usersService := services.NewUsersService(usersRepository)
	usersController := controllers.NewUsersController(usersService)

	router := http.NewServeMux()

	router.HandleFunc("POST /signup", usersController.HandleSignupUser)
	router.HandleFunc("GET /users/{email}", usersController.HandleGetUser)
	router.HandleFunc("GET /users", usersController.HandleGetAllUsers)
	router.HandleFunc("POST /login", usersController.HandleLoginUser)
	router.HandleFunc("POST /logout", usersController.HandleLogoutUser)

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
		Addr:    os.Getenv("SERVER_PORT"),
		Handler: handler,
	}
}
