package repositories

import (
	"context"
	"logger/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogEntryRepositoryImpl struct {
	mongoClient *mongo.Client
}

func NewLogEntryRepository(mongoClient *mongo.Client) LogEntryRepositoryInterface {
	return &LogEntryRepositoryImpl{
		mongoClient: mongoClient,
	}
}

func (lr *LogEntryRepositoryImpl) QueryInsertLogEntry(entry *models.LogEntry) *models.ResponseError {
	collection := lr.mongoClient.Database("logs").Collection("logs")

	_, err := collection.InsertOne(context.TODO(), models.LogEntry{
		Category:  entry.Category,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (lr *LogEntryRepositoryImpl) QueryGetAllLogEntries() ([]*models.LogEntry, *models.ResponseError) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := lr.mongoClient.Database("logs").Collection("logs")

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer cursor.Close(ctx)

	var logs []*models.LogEntry

	for cursor.Next(ctx) {
		var logItem models.LogEntry
		err := cursor.Decode(&logItem)

		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}

		logs = append(logs, &logItem)
	}

	return logs, nil
}

func (lr *LogEntryRepositoryImpl) QueryGetOneLogEntry(id string) (*models.LogEntry, *models.ResponseError) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := lr.mongoClient.Database("logs").Collection("logs")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	var entry models.LogEntry
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &entry, nil
}
