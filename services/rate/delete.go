package rate

import (
	"encoding/json"
	"foreign-currency-go/db"
	"foreign-currency-go/models"
	"foreign-currency-go/utils"
	"net/http"
)

type reqBodyDelete struct {
	ID uint `json:"id"`
}

// Delete func
func Delete(w http.ResponseWriter, r *http.Request) {
	var body reqBodyDelete

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		utils.BadRequestMessage(w, utils.ResponseError{
			Error: err.Error(),
		})
		return
	}

	result := db.DB.Delete(&models.Rate{}, body.ID)

	if result.Error != nil {
		utils.InternalServerErrorMessage(w, utils.ResponseError{
			Error: result.Error.Error(),
		})
		return
	}

	utils.DeletedMessage(w)
}
