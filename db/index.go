package db

import (
	"foreign-currency-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB var
var DB *gorm.DB

// Init func
func Init() {
	dsn := "root:password@tcp(127.0.0.1:3306)/transaction?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.Rate{})
	db.AutoMigrate(&models.Exchange{})

	DB = db
}
