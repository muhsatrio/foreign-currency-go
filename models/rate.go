package models

import (
	"gorm.io/gorm"
)

// Rate struct
type Rate struct {
	gorm.Model
	From     string     `json:"from"`
	To       string     `json:"to"`
	Exchange []Exchange `json:"exchanges" gorm:"foreignKey:From"`
}
