package database

import (
	"os"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/rustzz/todos/internal/models"
)

// ConnectDatabase : ...
func ConnectDatabase() (db *gorm.DB) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
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
