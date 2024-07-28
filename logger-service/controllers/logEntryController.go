package controllers

import (
	"encoding/json"
	"logger/services"
	"net/http"
)

type LogEntryController struct {
	logEntryService services.LogEntryServiceInterface
}

func NewLogEntryController(logEntryService services.LogEntryServiceInterface) *LogEntryController {
	return &LogEntryController{
		logEntryService: logEntryService,
	}
}

func (lc LogEntryController) HandleGetAllLogEntries(w http.ResponseWriter, r *http.Request) {
	logEntries, responseErr := lc.logEntryService.GetAllLogEntries()

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	responseJson, err := json.Marshal(&logEntries)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (lc LogEntryController) HandleGetOneLogEntry(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	logEntry, responseErr := lc.logEntryService.GetOneLogEntry(id)

	if responseErr != nil {
		http.Error(w, responseErr.Message, responseErr.Status)
		return
	}

	responseJson, err := json.Marshal(&logEntry)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
