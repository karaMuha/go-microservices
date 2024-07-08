package repositories

import "logger/models"

type LogEntryRepositoryInterface interface {
	QueryInsertLogEntry(entry *models.LogEntry) *models.ResponseError

	QueryGetAllLogEntries() ([]*models.LogEntry, *models.ResponseError)

	QueryGetOneLogEntry(id string) (*models.LogEntry, *models.ResponseError)
}
