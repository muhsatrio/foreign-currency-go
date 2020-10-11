package controllers

import (
	"encoding/json"
	"foreign-currency-go/db"
	"foreign-currency-go/models"
	"net/http"
)

type reqBodyExchange struct {
	Date string  `json:"date"`
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

type responseExchange struct {
	ID uint `json:"id"`
}

// ExchangeCreate func
func ExchangeCreate(w http.ResponseWriter, r *http.Request) {
	var body reqBodyExchange

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	exchange := models.Exchange{
		Date:  body.Date,
		From:  body.From,
		To:    body.To,
		Value: body.Rate,
	}

	result := db.DB.Create(&exchange)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(responseExchange{
		ID: exchange.ID,
	})
}
