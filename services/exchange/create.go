package exchange

import (
	"encoding/json"
	"foreign-currency-go/db"
	"foreign-currency-go/models"
	"foreign-currency-go/utils"
	"net/http"
)

type reqBody struct {
	Date string  `json:"date"`
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

// Create func
func Create(w http.ResponseWriter, r *http.Request) {
	var body reqBody
	var rate models.Rate

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		utils.BadRequestMessage(w, utils.ResponseError{
			Error: err.Error(),
		})
		return
	}

	result := db.DB.Where(&models.Rate{From: body.From, To: body.To}).First(&rate)

	if result.Error != nil {
		utils.BadRequestMessage(w, utils.ResponseError{
			Error: result.Error.Error(),
		})
		return
	}

	result = db.DB.Where(&models.Exchange{Date: body.Date, From: body.From, To: body.To}).Find(&models.Exchange{})

	if result.RowsAffected > 0 {
		utils.BadRequestMessage(w, utils.ResponseError{
			Error: "Rate on inputed Date has been exist",
		})
		return
	}

	exchange := models.Exchange{
		Date:   body.Date,
		From:   body.From,
		To:     body.To,
		Value:  body.Rate,
		RateID: rate.ID,
	}

	result = db.DB.Create(&exchange)

	if result.Error != nil {
		utils.InternalServerErrorMessage(w, utils.ResponseError{
			Error: result.Error.Error(),
		})
		return
	}

	utils.CreatedMessage(w, utils.ResponseCreated{
		ID: exchange.ID,
	})
}
