package database

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"loginstuff/models"
)

var DB *gorm.DB

func Connect() {
	// Establish connection to the database
	dsn := "host=localhost user=admin password=DBPASSWORD dbname=go_auth port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg("Failed to connect to database")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}
