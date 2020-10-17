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
func (r ResponseCreated) CreatedMessage(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(r)
}

// DeletedMessage func
func DeletedMessage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(""))
}

// BadRequestMessage func
func (r ResponseError) BadRequestMessage(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(r)
}

// InternalServerErrorMessage func
func (r ResponseError) InternalServerErrorMessage(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(r)
}

// NotFoundMessage func
func (r ResponseError) NotFoundMessage(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(r)
}
