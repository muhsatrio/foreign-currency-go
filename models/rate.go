package models

import (
	"gorm.io/gorm"
)

// Rate struct
type Rate struct {
	gorm.Model
	From     string
	To       string
	Exchange []Exchange `gorm:"foreignKey:RateID;OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
