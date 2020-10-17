package rate

import (
	"errors"
	"foreign-currency-go/db"
	"foreign-currency-go/models"
	"foreign-currency-go/utils"
	"net/http"
)

// ReqBodyCreate struct
type ReqBodyCreate struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (body ReqBodyCreate) serviceCreate(w http.ResponseWriter) error {

	result := db.DB.Where(&models.Rate{From: body.From, To: body.To}).Find(&models.Rate{})

	if result.RowsAffected > 0 {
		utils.ResponseError{Error: "Rate inputed has been exist"}.BadRequestMessage(w)
		return errors.New("Rate inputed has been exist")
	}

	rate := models.Rate{
		From: body.From,
		To:   body.To,
	}

	result = db.DB.Create(&rate)

	if result.Error != nil {
		utils.ResponseError{Error: result.Error.Error()}.InternalServerErrorMessage(w)

		return result.Error
	}

	utils.ResponseCreated{ID: rate.ID}.CreatedMessage(w)

	return nil
}

// ReqBodyDelete struct
type ReqBodyDelete struct {
	ID uint `json:"id"`
}

func (body ReqBodyDelete) serviceDelete(w http.ResponseWriter) error {

	result := db.DB.Delete(&models.Rate{}, body.ID)

	if result.Error != nil {
		utils.ResponseError{Error: result.Error.Error()}.InternalServerErrorMessage(w)
		return result.Error
	}

	utils.DeletedMessage(w)

	return nil
}
