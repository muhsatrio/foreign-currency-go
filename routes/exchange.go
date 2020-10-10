package routes

import "net/http"

func exchangeRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("GET /exchange"))
	case "POST":
		w.Write([]byte("POST /exchange"))
	default:
		http.Error(w, "Not Implemented", http.StatusBadRequest)
	}
}
