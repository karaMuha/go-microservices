package services

import "logger/models"

type LogEntryServiceInterface interface {
	InsertLogEntry(entry *models.LogEntry) *models.ResponseError

	GetAllLogEntries() ([]*models.LogEntry, *models.ResponseError)

	GetOneLogEntry(id string) (*models.LogEntry, *models.ResponseError)
}
