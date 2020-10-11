package models

import (
	"gorm.io/gorm"
)

// Exchange struct
type Exchange struct {
	gorm.Model
	Date   string
	From   string
	To     string
	Value  float64
	RateID uint
}
