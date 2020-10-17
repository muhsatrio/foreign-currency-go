package routes

import (
	"foreign-currency-go/exchange"
	"net/http"
)

func exchangeRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("GET /exchange"))
	case "POST":
		exchange.ControllerCreate(w, r)
	default:
		http.Error(w, "Not Implemented", http.StatusBadRequest)
	}
}
