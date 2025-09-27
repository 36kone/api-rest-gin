package database

import (
	"log"

	"github.com/36kone/api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DataBaseConnection() {
	connectinString := "host=localhost user=root password=root dbname=root"
	DB, err = gorm.Open(postgres.Open(connectinString))
	if err != nil {
		log.Panic("Failed to connect to database")
	}
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
