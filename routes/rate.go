package routes

import (
	"foreign-currency-go/rate"
	"foreign-currency-go/utils"
	"net/http"
)

func rateRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		rate.ControllerCreate(w, r)
	case "DELETE":
		rate.ControllerDelete(w, r)
	default:
		utils.ResponseError{Error: "Not Implemented"}.BadRequestMessage(w)
		return
	}
}
