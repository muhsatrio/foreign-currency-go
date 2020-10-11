package routes

import (
	"foreign-currency-go/controllers"
	"net/http"
)

func exchangeRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("GET /exchange"))
	case "POST":
		controllers.ExchangeCreate(w, r)
	default:
		http.Error(w, "Not Implemented", http.StatusBadRequest)
	}
}
