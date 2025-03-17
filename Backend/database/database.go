package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("social_media.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to datebase")
	}
	DB = database
}
