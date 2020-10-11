package routes

import (
	"foreign-currency-go/controllers"
	"net/http"
)

func rateRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		controllers.RateCreate(w, r)
	case "DELETE":
		w.Write([]byte("DELETE /rate"))
	default:
		http.Error(w, "Not Implemented", http.StatusBadRequest)
		return
	}
}
