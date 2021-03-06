package exchange

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
