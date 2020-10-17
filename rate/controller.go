package rate

import (
	"encoding/json"
	"foreign-currency-go/utils"
	"net/http"
)

// ControllerCreate func
func ControllerCreate(w http.ResponseWriter, r *http.Request) error {
	var body ReqBodyCreate

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		utils.ResponseError{Error: err.Error()}.BadRequestMessage(w)
		return err
	}

	body.serviceCreate(w)

	return nil
}

// ControllerDelete func
func ControllerDelete(w http.ResponseWriter, r *http.Request) error {
	var body ReqBodyDelete

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		utils.ResponseError{Error: err.Error()}.BadRequestMessage(w)
		return err
	}

	body.serviceDelete(w)

	return nil
}
