package db

import (
	"foreign-currency-go/models"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Migration var
var Migration *gormigrate.Gormigrate

// Migrate func
func Migrate(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "1",
			Migrate: func(tx *gorm.DB) error {
				type Rate struct {
					gorm.Model
					From     string
					To       string
					Exchange []models.Exchange `gorm:"foreignKey:RateID;OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}
				return db.AutoMigrate(&Rate{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("rate")
			},
		},
		{
			ID: "2",
			Migrate: func(tx *gorm.DB) error {
				type Exchange struct {
					gorm.Model
					Date   string
					From   string
					To     string
					Value  float64
					RateID uint
				}
				return db.AutoMigrate(&Exchange{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("exchange")
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalln("Failed to migrate: ", err)
	}
}
