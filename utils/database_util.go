package utils

import (
	"fmt"
	"go-rest-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dbname string) {
	var dbError error
	dbPath := fmt.Sprintf("./database/%s.db", dbname)
	DB, dbError = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if dbError != nil {
		panic(dbError)
	}

	DB.AutoMigrate(&models.User{})
}
