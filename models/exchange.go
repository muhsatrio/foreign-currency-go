package models

import (
	"gorm.io/gorm"
)

// Exchange struct
type Exchange struct {
	gorm.Model
	Date string `json:"date"`
	From string `json:"from"`
	To   string `json:"to"`
	Rate uint   `json:"rate"`
}
