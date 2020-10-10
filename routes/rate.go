package routes

import "net/http"

func rateRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Write([]byte("POST /rate"))
	case "DELETE":
		w.Write([]byte("DELETE /rate"))
	default:
		http.Error(w, "Not Implemented", http.StatusBadRequest)
	}
}
