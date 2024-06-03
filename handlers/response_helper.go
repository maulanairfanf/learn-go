package handlers

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	response := APIResponse{
		Status: http.StatusOK,
		Data:   data,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, status int, message string) {
	response := APIResponse{
		Status: status,
		Data: map[string]string{
			"error": message,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}