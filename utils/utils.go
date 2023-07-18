package utils

import (
	"encoding/json"
	"net/http"
)

type WriteResponseParams struct {
	W       http.ResponseWriter `json:"w"`
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Success bool                `json:"success"`
	V       interface{}         `json:"v"`
}

type ApiResponse struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data" default:"nil"`
}

func WriteResponse(w http.ResponseWriter, status int, message string, success bool, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	response := &ApiResponse{
		Status:  status,
		Success: success,
		Message: message,
		Data:    v,
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func WriteSuccessResponse(w http.ResponseWriter, message string, v interface{}) error {
	return WriteResponse(w, http.StatusOK, message, true, v)
}

func WriteFailureResponse(w http.ResponseWriter, message string) error {
	return WriteResponse(w, http.StatusBadRequest, message, false, nil)
}

func WritePermissionDenied(w http.ResponseWriter) error {
	return WriteResponse(w, http.StatusForbidden, "Permission denied", false, nil)
}
