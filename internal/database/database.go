package database

import (
	"os"

	"github.com/rustzz/todos/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDatabase : ...
func ConnectDatabase() (db *gorm.DB) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

// Migrate : ...
func Migrate() {
	db := ConnectDatabase()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Note{})
	dbConn, _ := db.DB()
	dbConn.Close()
	return
}
