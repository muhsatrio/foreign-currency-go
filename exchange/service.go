package exchange

import (
	"errors"
	"foreign-currency-go/db"
	"foreign-currency-go/models"
	"foreign-currency-go/utils"
	"net/http"
)

var w http.ResponseWriter

// ReqBodyCreate struct
type ReqBodyCreate struct {
	Date string  `json:"date"`
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

func (body ReqBodyCreate) serviceCreate(w http.ResponseWriter) error {
	var rate models.Rate

	result := db.DB.Where(&models.Rate{From: body.From, To: body.To}).First(&rate)

	if result.Error != nil {
		utils.BadRequestMessage(w, utils.ResponseError{
			Error: result.Error.Error(),
		})
		return result.Error
	}

	result = db.DB.Where(&models.Exchange{Date: body.Date, From: body.From, To: body.To}).Find(&models.Exchange{})

	if result.RowsAffected > 0 {
		utils.BadRequestMessage(w, utils.ResponseError{
			Error: "Rate on inputed Date has been exist",
		})
		return errors.New("Rate on inputed Date has been exist")
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
		return result.Error
	}

	utils.CreatedMessage(w, utils.ResponseCreated{
		ID: exchange.ID,
	})

	return nil
}
