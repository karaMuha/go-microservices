package server

import (
	"logger/controllers"
	"logger/events"
	"logger/repositories"
	"logger/services"
	"net/http"
	"os"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitHttpServer(mongoClient *mongo.Client, mqConnection *amqp091.Connection) (*http.Server, error) {
	logEntryRepository := repositories.NewLogEntryRepository(mongoClient)
	logEntryService := services.NewLogEntryService(logEntryRepository)
	logEntryController := controllers.NewLogEntryController(logEntryService)

	router := http.NewServeMux()

	router.HandleFunc("POST /log", logEntryController.HandleInsertLogEntry)
	router.HandleFunc("GET /log", logEntryController.HandleGetAllLogEntries)
	router.HandleFunc("GET /log/{id}", logEntryController.HandleGetOneLogEntry)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	handler := c.Handler(router)

	eventConsumer, err := events.NewEventConsumer(mqConnection, logEntryService)

	if err != nil {
		return nil, err
	}

	err = eventConsumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})

	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:    os.Getenv("SERVER_PORT"),
		Handler: handler,
	}, nil
}
