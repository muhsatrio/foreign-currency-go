package routes

import (
	"foreign-currency-go/rate"
	"net/http"
)

func rateRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		rate.ControllerCreate(w, r)
	case "DELETE":
		rate.ControllerDelete(w, r)
	default:
		http.Error(w, "Not Implemented", http.StatusBadRequest)
		return
	}
}
