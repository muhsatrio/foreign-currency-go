package routes

import (
	"foreign-currency-go/utils"
	"net/http"
)

func homeRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("API is Running"))
	default:
		utils.ResponseError{Error: "Not Implemented"}.BadRequestMessage(w)
		return
	}
}
