package database

import (
	"fmt"
	"github/jashandeep31/todo/database/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)

	_DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = _DB
	if err != nil {
		return err
	}

	DB.AutoMigrate(&models.Task{}, &models.User{})

	return nil
}
