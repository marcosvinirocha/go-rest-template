package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseFormat is a standard format for API responses
type ResponseFormat struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SendJSONResponse sends a JSON response with the given data and status code
func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// SendSuccessResponse sends a success JSON response
func SendSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	response := ResponseFormat{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}
	SendJSONResponse(w, http.StatusOK, response)
}

// SendErrorResponse sends an error JSON response
func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := ResponseFormat{
		Status:  statusCode,
		Message: message,
	}
	SendJSONResponse(w, statusCode, response)
}