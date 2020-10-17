package routes

import (
	"foreign-currency-go/exchange"
	"foreign-currency-go/utils"
	"net/http"
)

func exchangeRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("GET /exchange"))
	case "POST":
		exchange.ControllerCreate(w, r)
	default:
		utils.ResponseError{Error: "Not Implemented"}.BadRequestMessage(w)
		return
	}
}
