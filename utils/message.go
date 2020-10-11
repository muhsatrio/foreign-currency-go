package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseCreated struct
type ResponseCreated struct {
	ID uint `json:"id"`
}

// ResponseError struct
type ResponseError struct {
	Error string `json:"error"`
}

// CreatedMessage func
func CreatedMessage(w http.ResponseWriter, r ResponseCreated) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(r)
}

// BadRequestMessage func
func BadRequestMessage(w http.ResponseWriter, r ResponseError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(r)
}

// InternalServerErrorMessage func
func InternalServerErrorMessage(w http.ResponseWriter, r ResponseError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(r)
}
