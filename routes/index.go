package routes

import (
	"fmt"
	"net/http"
)

// RunRoutes func
func RunRoutes() {
	http.HandleFunc("/", homeRoutes)
	http.HandleFunc("/exchange", exchangeRoutes)
	http.HandleFunc("/rate", rateRoutes)

	fmt.Println("Server is running at port :8000")

	http.ListenAndServe(":8000", nil)
}
