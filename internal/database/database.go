package database

import (
	"os"

	"github.com/rustzz/todos/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (db *gorm.DB) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Note{})
	return db
}
