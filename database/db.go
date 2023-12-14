package database

import (
	"github.com/farhad-aman/cart-web-midterm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("mydb.db"), &gorm.Config{})
	if err != nil {
		log.Panicf("Failed to connect to database: %v\n", err)
	}

	log.Println("Database connection established")

	err = DB.AutoMigrate(&models.User{}, &models.Basket{})
	if err != nil {
		log.Panicf("Failed to migrate the database: %v\n", err)
	}
}
