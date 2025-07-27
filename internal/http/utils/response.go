package utils

import (
	"encoding/json"
	"net/http"
)

// make a struct that gives exact errors
type SetResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	statusSuccess = "success"
	statusFailure = "failed"
)

func SetResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func GeneralErrors(err error) SetResponseModel {
	return SetResponseModel{
		Status:  statusFailure,
		Message: err.Error(),
	}
}

