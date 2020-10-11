package controllers

import (
	"foreign-currency-go/services/rate"
	"net/http"
)

// RateCreate func
func RateCreate(w http.ResponseWriter, r *http.Request) {
	rate.Create(w, r)
}
