package server

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDb() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	clientOptions.SetAuth(options.Credential{
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error while connecting to mongo:", err)
		return nil, err
	}

	return client, nil
}
