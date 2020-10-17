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

	if result.RowsAffected <= 0 {
		utils.ResponseError{Error: "Rate not found"}.NotFoundMessage(w)
		return errors.New("Rate not found")
	}

	if result.Error != nil {
		utils.ResponseError{Error: result.Error.Error()}.BadRequestMessage(w)
		return result.Error
	}

	result = db.DB.Where(&models.Exchange{Date: body.Date, From: body.From, To: body.To}).Find(&models.Exchange{})

	if result.RowsAffected > 0 {
		utils.ResponseError{Error: "Rate on inputed Date has been exist"}.BadRequestMessage(w)
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
		utils.ResponseError{Error: result.Error.Error()}.InternalServerErrorMessage(w)

		return result.Error
	}

	utils.ResponseCreated{ID: exchange.ID}.CreatedMessage(w)

	return nil
}
