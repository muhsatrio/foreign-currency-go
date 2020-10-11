package controllers

import (
	"encoding/json"
	"foreign-currency-go/db"
	"foreign-currency-go/models"
	"net/http"
)

type reqBody struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type responseRate struct {
	ID uint `json:"id"`
}

// RateCreate func
func RateCreate(w http.ResponseWriter, r *http.Request) {
	var body reqBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	exchange := models.Exchange{
		From: body.From,
		To:   body.To,
	}

	result := db.DB.Create(&exchange)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(responseRate{
		ID: exchange.ID,
	})
}
