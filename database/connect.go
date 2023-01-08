package database

import (
	"log"
	"os"

	"github.com/XinceChan/blogbackend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	log.Println("Connect Successfully to database! :)")
	DB = db

	db.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
