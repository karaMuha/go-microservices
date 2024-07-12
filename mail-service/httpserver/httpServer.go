package httpserver

import (
	"mailer/controllers"
	"mailer/mailserver"
	"mailer/services"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func InitHttpServer(mailServer *mailserver.MailServer) *http.Server {
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

	return &http.Server{
		Addr:    os.Getenv("SERVER_PORT"),
		Handler: handler,
	}
}
