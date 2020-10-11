package models

import (
	"gorm.io/gorm"
)

// Rate struct
type Rate struct {
	gorm.Model
	From      string
	To        string
	Exchanges []Exchange `gorm:"foreignKey:RateID;references:ID"`
}
