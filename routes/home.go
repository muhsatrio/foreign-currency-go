package routes

import "net/http"

func homeRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("API is Running"))
	default:
		http.Error(w, "Not Implemented", http.StatusBadRequest)
	}
}
