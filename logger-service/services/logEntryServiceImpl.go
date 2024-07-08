package services

import (
	"logger/models"
	"logger/repositories"
)

type LogEntryServiceImpl struct {
	logEntryRepository repositories.LogEntryRepositoryInterface
}

func NewLogEntryService(logEntryRepository repositories.LogEntryRepositoryInterface) LogEntryServiceInterface {
	return LogEntryServiceImpl{
		logEntryRepository: logEntryRepository,
	}
}

func (ls LogEntryServiceImpl) InsertLogEntry(entry *models.LogEntry) *models.ResponseError {
	return ls.logEntryRepository.QueryInsertLogEntry(entry)
}

func (ls LogEntryServiceImpl) GetAllLogEntries() ([]*models.LogEntry, *models.ResponseError) {
	return ls.logEntryRepository.QueryGetAllLogEntries()
}

func (ls LogEntryServiceImpl) GetOneLogEntry(id string) (*models.LogEntry, *models.ResponseError) {
	return ls.logEntryRepository.QueryGetOneLogEntry(id)
}
