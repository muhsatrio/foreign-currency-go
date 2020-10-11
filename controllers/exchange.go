package controllers

import (
	"foreign-currency-go/services/exchange"
	"net/http"
)

// ExchangeCreate func
func ExchangeCreate(w http.ResponseWriter, r *http.Request) {
	exchange.Create(w, r)
}
