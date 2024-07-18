package httpserver

import (
	"mailer/controllers"
	"mailer/events"
	"mailer/mailserver"
	"mailer/services"
	"net/http"
	"os"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/cors"
)

func InitHttpServer(mailServer *mailserver.MailServer, mqConnection *amqp091.Connection) (*http.Server, error) {
	mailService := services.NewMailServiceImpl(mailServer)
	mailController := controllers.NewMailController(mailService)

	router := http.NewServeMux()
	router.HandleFunc("POST /send", mailController.HandleSendMail)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	handler := c.Handler(router)

	eventConsumer, err := events.NewEventConsumer(mqConnection, mailService)

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
