package rate

import (
	"encoding/json"
	"foreign-currency-go/db"
	"foreign-currency-go/models"
	"foreign-currency-go/utils"
	"net/http"
)

type reqBody struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// Create func
func Create(w http.ResponseWriter, r *http.Request) {
	var body reqBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		utils.BadRequestMessage(w, utils.ResponseError{
			Error: err.Error(),
		})
		return
	}

	rate := models.Rate{
		From: body.From,
		To:   body.To,
	}

	result := db.DB.Create(&rate)

	if result.Error != nil {
		utils.InternalServerErrorMessage(w, utils.ResponseError{
			Error: result.Error.Error(),
		})
		return
	}

	utils.CreatedMessage(w, utils.ResponseCreated{
		ID: rate.ID,
	})
}
